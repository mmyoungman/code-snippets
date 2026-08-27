// gcc 04_pthreads.c -o pthreads04 -lpthread

#include <errno.h>
#include <pthread.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

#define NUM_THREADS 8

pthread_mutex_t mutex;

void* work(void* arg) {
  int thread_index = *(int*)arg;
  printf("Starting thread %d\n", thread_index);
  sleep(1);

  // NOTE: test against EBUSY specifically. `while(pthread_mutex_trylock(...))`
  // spins forever on any other error (EINVAL, EOWNERDEAD, ...).
  int rc;
  while((rc = pthread_mutex_trylock(&mutex)) == EBUSY) {
    //printf("Thread %d waiting for lock\n", thread_index);
    // The 1s backoff is deliberately crude, to show what you give up by
    // polling: a thread can sit idle for up to a second after the lock frees.
    // pthread_mutex_lock would hand it over immediately.
    sleep(1);
  }
  if(rc) {
    fprintf(stderr, "pthread_mutex_trylock: %s\n", strerror(rc));
    return NULL;
  }

  printf("Thread %d got the lock!\n", thread_index);
  sleep(1);
  pthread_mutex_unlock(&mutex);
  return NULL;
}

int main(void) {
  pthread_t threads[NUM_THREADS];
  int rc;
  int thread_index[NUM_THREADS];

  if((rc = pthread_mutex_init(&mutex, NULL))) {
    fprintf(stderr, "pthread_mutex_init: %s\n", strerror(rc));
    return 1;
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    thread_index[i] = i;
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if((rc = pthread_create(&threads[i], NULL, work, &thread_index[i]))) {
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

  pthread_mutex_destroy(&mutex);

  return 0;
}
