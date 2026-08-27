// gcc 02_pthreads.c -o pthreads02 -lpthread

#include <pthread.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

#define NUM_THREADS 8

int counter = 0;
pthread_mutex_t mutex;

void* work(void* arg) {
  (void)arg;
  for(int i = 0; i < 1000000; i++) {
    pthread_mutex_lock(&mutex);
    counter++;
    pthread_mutex_unlock(&mutex);
  }
  return NULL;
}

int main(void) {
  pthread_t threads[NUM_THREADS];
  int rc;

  if((rc = pthread_mutex_init(&mutex, NULL))) {
    fprintf(stderr, "pthread_mutex_init: %s\n", strerror(rc));
    return 1;
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if((rc = pthread_create(&threads[i], NULL, work, NULL))) {
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

  printf("Counter: %d\n", counter);
  return 0;
}
