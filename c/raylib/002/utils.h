#ifndef UTILS_H
#define UTILS_H

#include <stdint.h>
#include <stdio.h>

typedef uint8_t u8;
typedef uint16_t u16;
typedef uint32_t u32;
typedef uint64_t u64;

typedef int8_t s8;
typedef int16_t s16;
typedef int32_t s32;
typedef int64_t s64;

typedef float f32;
typedef double f64;

#define kilobytes(value) ((value) * (u64)1024)
#define megabytes(value) (kilobytes(value) * 1024)
#define gigabytes(value) (megabytes(value) * 1024)
#define terabytes(value) (gigabytes(value) * 1024)

#ifdef TEST
#define dbg(msg, ...)                                                          \
  fprintf(stderr, "[DEBUG] (%s:%d) " msg "\n", __FILE__, __LINE__,             \
          ##__VA_ARGS__)
int shouldAssert = 0;
int assertFired = 0;

#define assert(expr)                                                           \
  if (!(expr)) {                                                               \
    assertFired = 1;                                                           \
    if (!shouldAssert) {                                                       \
      dbg("Assert failed: " #expr);                                            \
    }                                                                          \
  }

#define shouldAssert(expr)                                                     \
  shouldAssert = 1;                                                            \
  assertFired = 0;                                                             \
  assert(expr) shouldAssert = 0;                                               \
  if (!assertFired) {                                                          \
    log_err("Assert didn't fail: " #expr);                                     \
  }

#elif DEBUG
#define dbg(msg, ...)                                                          \
  fprintf(stderr, "[DEBUG] (%s:%d) " msg "\n", __FILE__, __LINE__,             \
          ##__VA_ARGS__)
#define assert(expr)                                                           \
  if (!(expr)) {                                                               \
    dbg("Assert failed: " #expr);                                              \
    exit(1);                                                                   \
  }
#define shouldAssert(expr)                                                     \
  log_warn("shouldAssert should only be used for testing");

#else
#define dbg(msg, ...)
#define assert(expr)
#define shouldAssert(expr)
#endif

#define log_err(msg, ...)                                                      \
  fprintf(stderr, "[ERROR] (%s:%d) " msg "\n", __FILE__, __LINE__,             \
          ##__VA_ARGS__)
#define log_warn(msg, ...)                                                     \
  fprintf(stderr, "[WARN] (%s:%d) " msg "\n", __FILE__, __LINE__, ##__VA_ARGS__)
#define log_info(msg, ...)                                                     \
  fprintf(stderr, "[INFO] (%s:%d) " msg "\n", __FILE__, __LINE__, ##__VA_ARGS__)

// Memory utils
#include <stdlib.h> // for malloc, realloc, calloc
void xmemset(unsigned char *ptr, unsigned char value, u64 size) {
  for (u64 i = 0; i < size; i++) {
    *ptr = value;
    ptr++;
  }
}

void xmemcpy(unsigned char *dst, unsigned char *src, u64 size) {
  while (size > 0) {
    *dst = *src;
    src++, dst++;
    size--;
  }
}

void *xmalloc(size_t num_bytes) {
  void *ptr = malloc(num_bytes);
  if (!ptr) {
    perror("xmalloc failed");
    exit(1);
  }
  return ptr;
}

void *xcalloc(size_t nitems, size_t num_bytes) {
  void *ptr = calloc(nitems, num_bytes);
  if (!ptr) {
    perror("xcalloc failed");
    exit(1);
  }
  return ptr;
}

void *xrealloc(void *ptr, size_t num_bytes) {
  ptr = realloc(ptr, num_bytes);
  if (!ptr) {
    perror("xrealloc failed");
    exit(1);
  }
  return ptr;
}

#endif // UTILS_H
