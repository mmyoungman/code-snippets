// gcc 08_pthreads.c -o pthreads08 -lpthread

#include <assert.h>
#include <pthread.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

#define NUM_THREADS 8
#define QUEUE_MAX 256

pthread_mutex_t mutexQueue;
pthread_cond_t condNotEmpty;
pthread_cond_t condNotFull;

typedef struct Task {
  int id;
} Task;

Task taskQueue[QUEUE_MAX];
int taskCount = 0;
int shuttingDown = 0;

void executeTask(Task* task) {
  sleep(1);
  printf("Completed taskId: %d\n", task->id);
}

void submitTask(Task task) {
  pthread_mutex_lock(&mutexQueue);

  // Back-pressure: a bounded queue needs the producer to block when full.
  // Without this, submitTask would run off the end of taskQueue - silent
  // memory corruption rather than a crash, which is the worst kind.
  while(taskCount == QUEUE_MAX) {
    pthread_cond_wait(&condNotFull, &mutexQueue);
  }
  assert(taskCount < QUEUE_MAX);

  taskQueue[taskCount] = task;
  taskCount++;
  pthread_mutex_unlock(&mutexQueue);
  pthread_cond_signal(&condNotEmpty);
}

void requestShutdown(void) {
  pthread_mutex_lock(&mutexQueue);
  shuttingDown = 1;
  pthread_mutex_unlock(&mutexQueue);
  // NOTE: broadcast, not signal. Every idle worker is blocked on condNotEmpty
  // and each one has to wake up and notice the flag; a signal would release
  // just one and leave the rest parked forever.
  pthread_cond_broadcast(&condNotEmpty);
}

void* startThread(void* args) {
  int index = *(int*)args;
  printf("Thread %d: Starting\n", index);
  while(1) {
    pthread_mutex_lock(&mutexQueue);
    while(taskCount == 0 && !shuttingDown) {
      pthread_cond_wait(&condNotEmpty, &mutexQueue);
    }
    // Drain first, then exit: only stop once the queue is actually empty.
    if(taskCount == 0) {
      pthread_mutex_unlock(&mutexQueue);
      break;
    }

    Task task = taskQueue[0];
    for(int i = 0; i < taskCount - 1; i++) {
      taskQueue[i] = taskQueue[i+1];
    }
    taskCount--;
    assert(taskCount >= 0);
    pthread_mutex_unlock(&mutexQueue);
    pthread_cond_signal(&condNotFull);

    executeTask(&task);
  }
  printf("Thread %d: Exiting\n", index);
  return NULL;
}

int main(void) {
  pthread_t threads[NUM_THREADS];
  int rc;

  if((rc = pthread_mutex_init(&mutexQueue, NULL))) {
    fprintf(stderr, "pthread_mutex_init: %s\n", strerror(rc));
    return 1;
  }
  if((rc = pthread_cond_init(&condNotEmpty, NULL))) {
    fprintf(stderr, "pthread_cond_init: %s\n", strerror(rc));
    return 1;
  }
  if((rc = pthread_cond_init(&condNotFull, NULL))) {
    fprintf(stderr, "pthread_cond_init: %s\n", strerror(rc));
    return 1;
  }

  int indexes[NUM_THREADS];
  for(int i = 0; i < NUM_THREADS; i++) {
    indexes[i] = i;
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if((rc = pthread_create(&threads[i], NULL, startThread, &indexes[i]))) {
      fprintf(stderr, "pthread_create: %s\n", strerror(rc));
      return 1;
    }
  }

  // Generate tasks
  for(int i = 0; i < 100; i++) {
    Task newTask = { .id = i, };
    submitTask(newTask);
  }

  requestShutdown();

  for(int i = 0; i < NUM_THREADS; i++) {
    if((rc = pthread_join(threads[i], NULL))) {
      fprintf(stderr, "pthread_join: %s\n", strerror(rc));
      return 1;
    }
  }

  pthread_mutex_destroy(&mutexQueue);
  pthread_cond_destroy(&condNotEmpty);
  pthread_cond_destroy(&condNotFull);

  return 0;
}
