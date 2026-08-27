// gcc 01_pthreads.c -o pthreads01 -lpthread

#include <pthread.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

// NOTE: a thread start routine must be `void* fn(void*)`. Declaring it with
// that exact signature means pthread_create needs no cast, which lets the
// compiler check the call. Casting the function pointer instead (e.g.
// `(void*)&work`) silences every warning and is undefined behaviour.
//
// Write the parameter out even when unused. `void* work()` is not the same
// thing: pre-C23 that means "unspecified parameters" and happens to be
// compatible, but in C23 `()` means `(void)` and pthread_create rejects it.
// gcc 15+ defaults to C23, so this is a live error, not a pedantic one.
void* work(void* arg) {
  (void)arg;
  printf("Starting thread\n");
  sleep(2);
  printf("Ending thread\n");
  return NULL;
}

int main(void) {
  pthread_t t1, t2;
  int rc;

  // NOTE: pthread functions return the error code directly. They do NOT set
  // errno, so use strerror(rc) rather than perror().
  if((rc = pthread_create(&t1, NULL, work, NULL))) {
    fprintf(stderr, "pthread_create: %s\n", strerror(rc));
    return 1;
  }
  if((rc = pthread_create(&t2, NULL, work, NULL))) {
    fprintf(stderr, "pthread_create: %s\n", strerror(rc));
    return 1;
  }

  if((rc = pthread_join(t1, NULL))) {
    fprintf(stderr, "pthread_join: %s\n", strerror(rc));
    return 1;
  }
  if((rc = pthread_join(t2, NULL))) {
    fprintf(stderr, "pthread_join: %s\n", strerror(rc));
    return 1;
  }

  return 0;
}
