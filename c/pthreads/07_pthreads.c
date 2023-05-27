// gcc 07_pthreads.c -o pthreads07 -lpthread

#include <pthread.h>
#include <semaphore.h>
#include <stdio.h>
#include <unistd.h>

#define NUM_THREADS 8

sem_t semaphore;

void work(void* arg) {
  sem_wait(&semaphore);
  sleep(1);
  int index = *(int*)arg;
  printf("Starting thread %d\n", index);
  sem_post(&semaphore);
}

int main(int argc, char* argv[]) {
  pthread_t threads[NUM_THREADS];

  int semCount = 2;
  sem_init(&semaphore, 0, semCount);

  int index[NUM_THREADS];
  for(int i = 0; i < NUM_THREADS; i++) {
    index[i] = i;
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if(pthread_create(&threads[i], NULL, (void*)&work, &index[i])) {
      return 1;
    }
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if(pthread_join(threads[i], NULL)) {
      return 1;
    }
  }

  sem_destroy(&semaphore);

  return 0;
}
