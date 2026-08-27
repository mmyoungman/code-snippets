// gcc 06_pthreads.c -o pthreads06 -lpthread

// NOTE: pthread_barrier_* and rand_r are POSIX, not ISO C. gcc's default
// (-std=gnu*) exposes them anyway, but under -std=c11 they vanish unless you
// ask for POSIX explicitly. Must come before any #include.
#define _POSIX_C_SOURCE 200809L

#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <unistd.h>

#define NUM_THREADS 16
// NOTE: must divide NUM_THREADS exactly. With a count that doesn't (e.g. 3
// into 16) there is always a remainder of threads parked at the barrier
// waiting for a group that never fills.
#define BARRIER_COUNT 4

pthread_barrier_t barrier;

typedef struct Worker {
  int index;
  unsigned int seed;
} Worker;

void* work(void* args) {
  Worker* worker = args;
  printf("Thread %d: Starting\n", worker->index);
  while(1) {
    // NOTE: rand() shares one global state across threads. glibc happens to
    // lock it, but POSIX doesn't require rand() to be thread-safe - use
    // rand_r() with per-thread seed state instead.
    sleep((rand_r(&worker->seed) % 20) + 1);
    printf("Thread %d: Waiting at barrier\n", worker->index);
    pthread_barrier_wait(&barrier);
    printf("Thread %d: Passed barrier\n", worker->index);
    sleep(20);
  }
  return NULL;
}

int main(void) {
  pthread_t threads[NUM_THREADS];
  int rc;
  Worker workers[NUM_THREADS];

  unsigned int baseSeed = (unsigned int)time(NULL);
  for(int i = 0; i < NUM_THREADS; i++) {
    workers[i].index = i;
    workers[i].seed = baseSeed + (unsigned int)i;
  }

  printf("barrierCount: %d\n", BARRIER_COUNT);

  if((rc = pthread_barrier_init(&barrier, NULL, BARRIER_COUNT))) {
    fprintf(stderr, "pthread_barrier_init: %s\n", strerror(rc));
    return 1;
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if((rc = pthread_create(&threads[i], NULL, work, &workers[i]))) {
      fprintf(stderr, "pthread_create: %s\n", strerror(rc));
      return 1;
    }
  }

  // NOTE: work() loops forever, so these joins never return - Ctrl-C to stop.
  // See 08/09/10 for the shutdown-flag pattern that lets workers exit cleanly.
  for(int i = 0; i < NUM_THREADS; i++) {
    if((rc = pthread_join(threads[i], NULL))) {
      fprintf(stderr, "pthread_join: %s\n", strerror(rc));
      return 1;
    }
  }

  pthread_barrier_destroy(&barrier);

  return 0;
}
