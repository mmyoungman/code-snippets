// gcc 02_pthreads.c -o pthreads02 -lpthread

#include <pthread.h>
#include <stdio.h>
#include <unistd.h>

#define NUM_THREADS 8

int counter = 0;
pthread_mutex_t mutex;

void work() {
  for(int i = 0; i < 1000000; i++) {
    pthread_mutex_lock(&mutex);
    counter++;
    pthread_mutex_unlock(&mutex);
  }
}

int main(int argc, char* argv[]) {
  pthread_t threads[NUM_THREADS];

  pthread_mutex_init(&mutex, NULL);

  for(int i = 0; i < NUM_THREADS; i++) {
    if(pthread_create(&threads[i], NULL, (void*)&work, NULL)) {
      return 1;
    }
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if(pthread_join(threads[i], NULL)) {
      return 1;
    }
  }

  pthread_mutex_destroy(&mutex);

  printf("Counter: %d\n", counter);
  return 0;
}
