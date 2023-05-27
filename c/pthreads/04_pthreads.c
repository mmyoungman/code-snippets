// gcc 04_pthreads.c -o pthreads04 -lpthread

#include <pthread.h>
#include <stdio.h>
#include <unistd.h>

#define NUM_THREADS 8

pthread_mutex_t mutex;

void work(void* arg) {
  int thread_index = *(int*)arg;
  printf("Starting thread %d\n", thread_index);
  sleep(1);
  while(pthread_mutex_trylock(&mutex)) {
    //printf("Thread %d waiting for lock\n", thread_index);
    sleep(1);
  }

  printf("Thread %d got the lock!\n", thread_index);
  sleep(1);
  pthread_mutex_unlock(&mutex);
}

int main(int argc, char* argv[]) {
  pthread_t threads[NUM_THREADS];
  int thread_index[NUM_THREADS];

  pthread_mutex_init(&mutex, NULL);

  for(int i = 0; i < NUM_THREADS; i++) {
    thread_index[i] = i;
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if(pthread_create(&threads[i], NULL, (void*)&work, &thread_index[i])) {
      return 1;
    }
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if(pthread_join(threads[i], NULL)) {
      return 1;
    }
  }

  pthread_mutex_destroy(&mutex);

  return 0;
}
