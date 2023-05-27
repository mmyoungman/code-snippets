// gcc 06_pthreads.c -o pthreads06 -lpthread

#include <pthread.h>
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>

#define NUM_THREADS 16

pthread_barrier_t barrier;

void work(void* args) {
  int index = *(int *)args;
  printf("Thread %d: Starting\n", index);
  while(1) {
    sleep((rand() % 20) + 1);
    printf("Thread %d: Waiting at barrier\n", index);
    pthread_barrier_wait(&barrier);
    printf("Thread %d: Passed barrier\n", index);
    sleep(20);
  }
}

int main(int argc, char* argv[]) {
  pthread_t threads[NUM_THREADS];
  int indexes[NUM_THREADS];

  for(int i = 0; i < NUM_THREADS; i++) {
    indexes[i] = i;
  }

  int barrierCount = 3;
  printf("barrierCount: %d\n", barrierCount);

  pthread_barrier_init(&barrier, NULL, barrierCount);

  for(int i = 0; i < NUM_THREADS; i++) {
    if(pthread_create(&threads[i], NULL, (void*)&work, &indexes[i])) {
      return 1;
    }
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if(pthread_join(threads[i], NULL)) {
      return 1;
    }
  }

  pthread_barrier_destroy(&barrier);

  return 0;
}
