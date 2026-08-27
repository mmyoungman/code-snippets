// gcc 07_pthreads.c -o pthreads07 -lpthread

#include <errno.h>
#include <pthread.h>
#include <semaphore.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

#define NUM_THREADS 8

sem_t semaphore;

void* work(void* arg) {
  // NOTE: the sem_* functions do NOT follow the pthread_* convention. They
  // return -1 and set errno, so they want perror()/errno, not strerror(rc).
  //
  // sem_wait is also interruptible: a signal delivered to this thread makes it
  // fail with EINTR without having decremented the semaphore, so retry.
  while(sem_wait(&semaphore) == -1) {
    if(errno != EINTR) {
      perror("sem_wait");
      return NULL;
    }
  }

  sleep(1);
  int index = *(int*)arg;
  printf("Starting thread %d\n", index);

  if(sem_post(&semaphore) == -1) {
    perror("sem_post");
  }
  return NULL;
}

int main(void) {
  pthread_t threads[NUM_THREADS];
  int rc;

  int semCount = 2;
  if(sem_init(&semaphore, 0, semCount) == -1) {
    perror("sem_init");
    return 1;
  }

  int index[NUM_THREADS];
  for(int i = 0; i < NUM_THREADS; i++) {
    index[i] = i;
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if((rc = pthread_create(&threads[i], NULL, work, &index[i]))) {
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

  sem_destroy(&semaphore);

  return 0;
}
