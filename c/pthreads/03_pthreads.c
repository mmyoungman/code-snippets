// gcc 03_pthreads.c -o pthreads03 -lpthread

#include <pthread.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

#define NUM_THREADS 8

// Getting a result out of a thread: give each thread its own job struct with
// an input field and an output field, and join with NULL.
//
// NOTE: the tempting alternative - `pthread_join(t, (void*)&myInt)` - is a
// buffer overflow. pthread_join takes a void** and writes a full sizeof(void*)
// (8 bytes on x86-64) through it, so pointing it at a 4-byte int clobbers the
// 4 bytes that follow. If you really want to pass a value back through the
// return slot, it has to go via void*: `return (void*)(intptr_t)result;`.
typedef struct SquareJob {
  int index;   // in
  int result;  // out
} SquareJob;

void* square(void* arg) {
  SquareJob* job = arg;
  printf("Starting thread %d\n", job->index);
  sleep(2);
  job->result = job->index * job->index;
  printf("Thread %d, result: %d\n", job->index, job->result);
  return NULL;
}

int main(void) {
  pthread_t threads[NUM_THREADS];
  int rc;
  SquareJob jobs[NUM_THREADS];

  for(int i = 0; i < NUM_THREADS; i++) {
    jobs[i].index = i;
    jobs[i].result = 0;
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if((rc = pthread_create(&threads[i], NULL, square, &jobs[i]))) {
      fprintf(stderr, "pthread_create: %s\n", strerror(rc));
      return 1;
    }
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if((rc = pthread_join(threads[i], NULL))) {
      fprintf(stderr, "pthread_join: %s\n", strerror(rc));
      return 1;
    }
  }

  printf("TIME FOR THE RESULTS!\n");
  for(int i = 0; i < NUM_THREADS; i++) {
    printf("Square of %d: %d\n", i, jobs[i].result);
  }
  return 0;
}
