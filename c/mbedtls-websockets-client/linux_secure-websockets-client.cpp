#include <stdlib.h>
#include <stdio.h>

#define MBEDTLS_DEPRECATED_REMOVED
#include <mbedtls/entropy.h>
#include<mbedtls/net_sockets.h>

#define SERVER_NAME "localhost"
#define SERVER_PORT "9393"

int main(int argc, char* argv[]) {
  // https://mbed-tls.readthedocs.io/en/latest/kb/how-to/mbedtls-tutorial/#setup
  mbedtls_net_context server_fd;
  mbedtls_net_init(&server_fd);

  //mbedtls_ssl_context ssl;
  //mbedtls_ssl_init(&ssl);

  mbedtls_ssl_config conf;
  mbedtls_ssl_config_init(&conf);

  //mbedtls_entropy_context entropy;
  //mbedtls_entropy_init(&entropy);

  if(mbedtls_net_connect(&server_fd, SERVER_NAME, SERVER_PORT, MBEDTLS_NET_PROTO_TCP)) {
    printf("Something went wrong 1!\n");
    exit(1);
  }

  if(mbedtls_ssl_config_defaults(&conf, MBEDTLS_SSL_IS_CLIENT, MBEDTLS_SSL_TRANSPORT_STREAM, MBEDTLS_SSL_PRESET_DEFAULT)) {
    printf("Something went wrong 2!\n");
    exit(1);
  }

  // @MarkFix should switch to VERIFY_REQUIRED at some point
  mbedtls_ssl_conf_authmode(&conf, MBEDTLS_SSL_VERIFY_NONE);

  //mbedtls_net_accept

  //mbedtls_net_send
  //mbedtls_net_recv
}
