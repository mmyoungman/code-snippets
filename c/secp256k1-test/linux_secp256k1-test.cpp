#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <string.h>

#include "build/include/secp256k1.h"
#include "build/include/secp256k1_extrakeys.h"
#include "build/include/secp256k1_schnorrsig.h"

#if defined(_WIN32)
/*
 * The defined WIN32_NO_STATUS macro disables return code definitions in
 * windows.h, which avoids "macro redefinition" MSVC warnings in ntstatus.h.
 */
#define WIN32_NO_STATUS
#include <windows.h>
#undef WIN32_NO_STATUS
#include <ntstatus.h>
#include <bcrypt.h>
#elif defined(__linux__) || defined(__APPLE__) || defined(__FreeBSD__)
#include <sys/random.h>
#elif defined(__OpenBSD__)
#include <unistd.h>
#else
#error "Couldn't identify the OS"
#endif

#include <stddef.h>
#include <limits.h>
#include <stdio.h>

struct program_state {
  bool has_context = false;
  secp256k1_context *ctx;
  bool has_keypair = false;
  secp256k1_keypair keypair;
  unsigned char sec_key[32];
  unsigned char serialized_pubkey[32];
};

/* Returns 1 on success, and 0 on failure. */
static int fill_random(unsigned char* data, size_t size) {
#if defined(_WIN32)
    NTSTATUS res = BCryptGenRandom(NULL, data, size, BCRYPT_USE_SYSTEM_PREFERRED_RNG);
    if (res != STATUS_SUCCESS || size > ULONG_MAX) {
        return 0;
    } else {
        return 1;
    }
#elif defined(__linux__) || defined(__FreeBSD__)
    /* If `getrandom(2)` is not available you should fallback to /dev/urandom */
    ssize_t res = getrandom(data, size, 0);
    if (res < 0 || (size_t)res != size ) {
        return 0;
    } else {
        return 1;
    }
#elif defined(__APPLE__) || defined(__OpenBSD__)
    /* If `getentropy(2)` is not available you should fallback to either
     * `SecRandomCopyBytes` or /dev/urandom */
    int res = getentropy(data, size);
    if (res == 0) {
        return 1;
    } else {
        return 0;
    }
#endif
    return 0;
}

#if defined(_MSC_VER)
// For SecureZeroMemory
#include <Windows.h>
#endif
/* Cleanses memory to prevent leaking sensitive info. Won't be optimized out. */
static void secure_erase(void *ptr, size_t len) {
#if defined(_MSC_VER)
    /* SecureZeroMemory is guaranteed not to be optimized out by MSVC. */
    SecureZeroMemory(ptr, len);
#elif defined(__GNUC__)
    /* We use a memory barrier that scares the compiler away from optimizing out the memset.
     *
     * Quoting Adam Langley <agl@google.com> in commit ad1907fe73334d6c696c8539646c21b11178f20f
     * in BoringSSL (ISC License):
     *    As best as we can tell, this is sufficient to break any optimisations that
     *    might try to eliminate "superfluous" memsets.
     * This method used in memzero_explicit() the Linux kernel, too. Its advantage is that it is
     * pretty efficient, because the compiler can still implement the memset() efficently,
     * just not remove it entirely. See "Dead Store Elimination (Still) Considered Harmful" by
     * Yang et al. (USENIX Security 2017) for more background.
     */
    memset(ptr, 0, len);
    __asm__ __volatile__("" : : "r"(ptr) : "memory");
#else
    void *(*volatile const volatile_memset)(void *, int, size_t) = memset;
    volatile_memset(ptr, 0, len);
#endif
}

int generateContext(program_state *state) {
  assert(state->has_context == false);
  secp256k1_context *context = secp256k1_context_create(SECP256K1_CONTEXT_NONE);

  unsigned char randomize[32];
  if (!fill_random(randomize, sizeof(randomize))) {
    printf("Failed to generate randomness\n");
    return 1;
  }
  int return_val = secp256k1_context_randomize(context, randomize);
  assert(return_val);

  state->ctx = context;
  state->has_context = true;

  return 0;
}

int generatePublicPrivateKeys(program_state *state) {
  assert(state->has_context);
  assert(state->has_keypair == false);
  while (1) {
    if (!fill_random(state->sec_key, sizeof(state->sec_key))) {
      printf("Failed to generate randomness\n");
      return 1;
    }
    if (secp256k1_keypair_create(state->ctx, &state->keypair, state->sec_key)) {
      break;
    }
  }

  secp256k1_xonly_pubkey pubkey;
  int return_val = secp256k1_keypair_xonly_pub(state->ctx, &pubkey, NULL, &state->keypair);
  assert(return_val);

  return_val = secp256k1_xonly_pubkey_serialize(state->ctx, state->serialized_pubkey, &pubkey);
  assert(return_val);

  state->has_keypair = true;
  return 0;
}

int generateSignature(program_state *state,
                      unsigned char eventId[32],
                      unsigned char signature[64]) {
  assert(state->has_context);
  assert(state->has_keypair);

  unsigned char auxiliary_rand[32];
  if (!fill_random(auxiliary_rand, sizeof(auxiliary_rand))) {
    printf("Failed to generate randomness\n");
    return 1;
  }

  int return_val = secp256k1_schnorrsig_sign32(
    state->ctx,
    signature,
    eventId,
    &state->keypair,
    auxiliary_rand);
  assert(return_val);

  return 0;
}

int validateSignature(unsigned char eventId[32],
                      unsigned char signature[64],
                      unsigned char serialized_pubkey[32]) {
  secp256k1_xonly_pubkey pubkey;
  if (!secp256k1_xonly_pubkey_parse(secp256k1_context_static, &pubkey, serialized_pubkey)) {
    printf("Failed to parse the public key\n");
    return 1;
  }

  int is_signature_valid =
      secp256k1_schnorrsig_verify(secp256k1_context_static, signature, eventId, 32, &pubkey);

  if(!is_signature_valid) {
    return 1;
  }

  return 0;
}

void exitProgram(program_state *state) {
  assert(state->has_context);
  secp256k1_context_destroy(state->ctx);
  secure_erase(state->sec_key, sizeof(state->sec_key));
}

static void print_hex(unsigned char* data, size_t size) {
    size_t i;
    printf("0x");
    for (i = 0; i < size; i++) {
        printf("%02x", data[i]);
    }
    printf("\n");
}

int main(int argc, char *argv[]) {
  program_state *state = (program_state*) calloc(1, sizeof(program_state));
  state->has_context = false;
  state->has_keypair = false;
  if( generateContext(state) ) printf("Failed to generate context\n");

  if( generatePublicPrivateKeys(state) ) printf("Failed to generate keys\n");

  unsigned char eventId[32];
  // generate msg_hash
  unsigned char msg[13] = "Hello World!";
  unsigned char tag[18] = "my_fancy_protocol";
  int return_val;
  return_val = secp256k1_tagged_sha256(state->ctx, eventId, tag, sizeof(tag), msg,
                                       sizeof(msg));
  assert(return_val);

  unsigned char signature[64];
  if ( generateSignature(state, eventId, signature) )
    printf("Failed to generate signature\n");

  /*** Verification ***/
  if ( validateSignature(eventId, signature, state->serialized_pubkey) ) {
    printf("Failed to validate signature\n");
  }

  printf("Validated signature\n");

  /* Verify a signature. This will return 1 if it's valid and 0 if it's not. */

  printf("Secret Key: ");
  print_hex(state->sec_key, sizeof(state->sec_key));
  printf("Public Key: ");
  print_hex(state->serialized_pubkey, sizeof(state->serialized_pubkey));
  printf("Signature: ");
  print_hex(signature, sizeof(signature));

  secure_erase(state->sec_key, sizeof(state->sec_key));
  exitProgram(state);
  return 0;
}
