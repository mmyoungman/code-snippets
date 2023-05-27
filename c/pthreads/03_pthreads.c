// gcc 03_pthreads.c -o pthreads03 -lpthread

#include <pthread.h>
#include <stdio.h>
#include <unistd.h>

#define NUM_THREADS 8

int square(void* arg) {
  int index = *(int*)arg;
  printf("Starting thread %d\n", index);
  sleep(2);
  int result = index * index;
  printf("Thread %d, result: %d\n", index, result);
  return result;
}

int main(int argc, char* argv[]) {
  pthread_t threads[NUM_THREADS];
  int index[NUM_THREADS];
  int squareResults[NUM_THREADS];

  for(int i = 0; i < NUM_THREADS; i++) {
    index[i] = i;
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if(pthread_create(&threads[i], NULL, (void*)&square, &index[i])) {
      return 1;
    }
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if(pthread_join(threads[i], (void*)&squareResults[i])) {
      return 1;
    }
  }

  printf("TIME FOR THE RESULTS!\n");
  for(int i = 0; i < NUM_THREADS; i++) {
    printf("Square of %d: %d\n", i, squareResults[i]);
  }
  return 0;
}
