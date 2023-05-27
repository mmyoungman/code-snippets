// gcc 08_pthreads.c -o pthreads08 -lpthread

#include <pthread.h>
#include <stdio.h>
#include <unistd.h>

#define NUM_THREADS 8

pthread_mutex_t mutexQueue;
pthread_cond_t condQueue;

typedef struct Task {
  int id;
} Task;

Task taskQueue[256];
int taskCount = 0;

void executeTask(Task* task) {
  sleep(1);
  printf("Completed taskId: %d\n", task->id);
}

void submitTask(Task task) {
  pthread_mutex_lock(&mutexQueue);
  taskQueue[taskCount] = task;
  taskCount++;
  pthread_mutex_unlock(&mutexQueue);
  pthread_cond_signal(&condQueue);
}

void startThread(void* args) {
  int index = *(int*)args;
  printf("Thread %d: Starting\n", index);
  while(1) {
    pthread_mutex_lock(&mutexQueue);
    while (taskCount == 0) {
      pthread_cond_wait(&condQueue, &mutexQueue);
    }

    Task task = taskQueue[0];
    for(int i = 0; i < taskCount - 1; i++) {
      taskQueue[i] = taskQueue[i+1];
    }
    taskCount--;
    pthread_mutex_unlock(&mutexQueue);
    executeTask(&task);
  }
}

int main(int argc, char* argv[]) {
  pthread_t threads[NUM_THREADS];
  pthread_mutex_init(&mutexQueue, NULL);
  pthread_cond_init(&condQueue, NULL);

  int indexes[NUM_THREADS];
  for(int i = 0; i < NUM_THREADS; i++) {
    indexes[i] = i;
  }

  for(int i = 0; i < NUM_THREADS; i++) {
    if(pthread_create(&threads[i], NULL, (void*)&startThread, &indexes[i])) {
      return 1;
    }
  }

  // Generate tasks
  for(int i = 0; i < 100; i++) {
    Task newTask = { .id = i, };
    submitTask(newTask);
  }

  // threads never finish so never join
  for(int i = 0; i < NUM_THREADS; i++) {
    if(pthread_join(threads[i], NULL)) {
      return 1;
    }
  }

  pthread_mutex_destroy(&mutexQueue);
  pthread_cond_destroy(&condQueue);

  return 0;
}
