#include <stdlib.h>
#include <string.h>

#include "mbedtls/net_sockets.h"
#include "mbedtls/ssl.h"
#include "mbedtls/entropy.h"
#include "mbedtls/ctr_drbg.h"
#include "mbedtls/base64.h"
#include "mbedtls/sha1.h"
#include "mbedtls/debug.h"

#define DEBUG_LEVEL 0

//#define SERVER_URL "wss://ws.ifelse.io"
#define SERVER_NAME "ws.ifelse.io"
//#define SERVER_NAME "mark.youngman.info"
#define SERVER_PORT "443"

const size_t maxConcatLen = 4096;

void generateSecStr(unsigned char *out, size_t outlen) {
  srand(time(NULL));
  const char ALLOWED[] = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";

  unsigned char randStr[16];

  for(int i = 0; i < 16; i++) {
    randStr[i] = ALLOWED[rand() % sizeof(ALLOWED)];
  }

  size_t encodedLen;
  mbedtls_base64_encode(out, outlen, &encodedLen, randStr, 16);

  //assert outlen >= encodedLen // maybe just >?
}

int verifySecWebSocketKey(unsigned char *in, size_t inlen, unsigned char *expected) {
  int concatLen = inlen + 36;
  if (concatLen <= maxConcatLen)
  {
    printf("concatLen > maxConcatLen");
    return 1;
  }

  unsigned char concat[maxConcatLen];
  memcpy(concat, in, inlen);
  memcpy(&concat[inlen], "258EAFA5-E914-47DA-95CA-C5AB0DC85B11", 36);

  unsigned char sha1result[20];
  if( mbedtls_sha1(concat, concatLen, sha1result)) {
    printf("sha1 failed\n");
  }

  unsigned char encodedResult[128];
  int encodedLen = 128;
  size_t writtenLen;
  mbedtls_base64_encode(encodedResult, encodedLen, &writtenLen, sha1result, 20);

  printf("Sec-WebSocket-Accept: %s\n\n", encodedResult);

  //assert encodedLen >= writtenLen

  unsigned char *a = encodedResult;
  unsigned char *b = expected;
  while((*a != '\0') && (*b != '\0') && (*a == *b))
      a++, b++;
  int result = ((*a == '\0') && (*b == '\0'));

  printf("verifySecWebSocketKey result: %d\n", result);
  return result;
}


static void my_debug_func(void *ctx, int level, const char *file, int line, const char *str) {
  ((void) level);
  fprintf((FILE *)ctx, "debug_func:\n%s:%04d: %s", file, line, str);
  fflush((FILE *)ctx);
}

int main(int argc, char *argv[]) {
  mbedtls_debug_set_threshold(DEBUG_LEVEL);

  mbedtls_entropy_context entropy;
  mbedtls_entropy_init(&entropy);

  mbedtls_ctr_drbg_context ctr_drbg;
  mbedtls_ctr_drbg_init(&ctr_drbg);
  int ret;
  const char *pers = "ssl_client1";
  if(( ret = mbedtls_ctr_drbg_seed(&ctr_drbg, mbedtls_entropy_func, &entropy,
                                   (const unsigned char*)pers,
                                   strlen(pers))) != 0) {
    printf("mbedtls_ctr_drbg_seed returned %d\n", ret);
    return 1;
  }

  mbedtls_net_context server_fd;
  mbedtls_net_init(&server_fd);
  if(( ret = mbedtls_net_connect(&server_fd, SERVER_NAME,
                                  SERVER_PORT, MBEDTLS_NET_PROTO_TCP))) {
    printf("mbedtls_net_connect returned %d\n\n", ret);
    return 1;
  }

  mbedtls_ssl_config conf;
  mbedtls_ssl_config_init(&conf);
  if((ret = mbedtls_ssl_config_defaults(&conf,
                                        MBEDTLS_SSL_IS_CLIENT,
                                        MBEDTLS_SSL_TRANSPORT_STREAM,
                                        MBEDTLS_SSL_PRESET_DEFAULT)) != 0) {
    printf("mbedtls_ssl_config_defaults returned %d\n\n", ret);
    return 1;
  }

  mbedtls_ssl_conf_authmode(&conf, MBEDTLS_SSL_VERIFY_NONE);
  /*
  mbedtls_x509_crt cacert;
  mbedtls_x509_crt_init(&cacert);

  const char *cafile = "path/to/trusted-ca-list-pem";
  if((ret = mbedtls_x509_crt_parse_file(&cacert, cafile)) != 0) {
    printf("mbedtls_x509_crt_parse returned -0x%x\n\n", -ret);
    return 1;
  }

  mbedtls_ssl_conf_ca_chain(&conf, &cacert, NULL);
  */

  mbedtls_ssl_conf_rng(&conf, mbedtls_ctr_drbg_random, &ctr_drbg);
  mbedtls_ssl_conf_dbg(&conf, my_debug_func, stdout);

  mbedtls_ssl_context ssl;
  mbedtls_ssl_init(&ssl);

  if((ret = mbedtls_ssl_setup(&ssl, &conf)) != 0) {
    printf("mbedtls_ssl_setup returned %d\n\n", ret);
    return 1;
  }

  if((ret = mbedtls_ssl_set_hostname(&ssl, SERVER_NAME)) != 0) {
    printf("mbedtls_ssl_set_hostname return %d\n\n", ret);
    return 1;
  }

  mbedtls_ssl_set_bio(&ssl, &server_fd, mbedtls_net_send, mbedtls_net_recv, NULL);

  unsigned char webSocketKey[128];
  generateSecStr(webSocketKey, 128);

  size_t webSocketKeyLen = strlen((char*)webSocketKey);
  unsigned char expected[128] = "blah";
  verifySecWebSocketKey(webSocketKey, webSocketKeyLen, expected);

  unsigned char *buf = (unsigned char*)calloc(1, 1024);
  int buflen = sprintf((char*) buf,
                       "GET / HTTP/1.1\r\n"
                       "Host: " SERVER_NAME "\r\n"
                       "Upgrade: websocket\r\n"
                       "Connection: Upgrade\r\n"
                       "Sec-WebSocket-Key: %s\r\n"
                       "Sec-WebSocket-Protocol: chat, superchat\r\n"
                       "Sec-WebSocket-Version: 13\r\n\r\n", webSocketKey);
  while((ret = mbedtls_ssl_write(&ssl, buf, buflen)) <= 0) {
    if(ret != 0) {
      printf("mbedtls_ssl_write returned -0x%x\n\n", -ret);
      return 1;
    }
  }

  printf("%d bytes written\n\n"
         "%s\n\n",
         ret,
         (char*)buf);

  do {
    unsigned char *readBuf = (unsigned char*)calloc(1, 4096);
    ret = mbedtls_ssl_read(&ssl, readBuf, 4096);
    if(ret <= 0) {
      printf("\n\nEOF\n\n");
      break;
    }

    buf[ret] = '\0';
    printf("%s", (char*)readBuf);
    fflush(stdout);

    // @MarkFix validate 101 response

    // @MarkFix below doesn't work because websockets - read rfc6455
    //int newLen = sprintf((char*) readBuf, "test\r\n");
    //while ((ret = mbedtls_ssl_write(&ssl, readBuf, newLen)) <= 0) {
    //  if (ret != 0) {
    //    printf("failed to write \"test\" message");
    //    return 1;
    //  }
    //}
  } while(1);

  // close
  //mbedtls_x509_crt_free(&cacert);
  mbedtls_net_free(&server_fd);
  mbedtls_ssl_free(&ssl);
  mbedtls_ssl_config_free(&conf);
  mbedtls_ctr_drbg_free(&ctr_drbg);
  mbedtls_entropy_free(&entropy);

  return 0;
}
