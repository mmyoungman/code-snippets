// gcc 10_pthreads.c -o pthreads10 -lpthread

#include <pthread.h>
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>

#define NUM_THREADS 4
#define WORK_QUEUE_MAX 1024

pthread_mutex_t mutexQueue;
pthread_cond_t condQueue;

typedef enum TaskType {
  PRINT_ID,
  SUM,
} TaskType;

typedef struct SumArgs {
  int a, b;
} SumArgs;

typedef struct Task {
  int id;
  TaskType type;
  union {
    SumArgs sumArgs;
    // other TaskType args structs go here
  };
} Task;

void printId(int id) {
  printf("PrintId: %d\n", id);
}

void sum(int a, int b) {
  printf("Sum: %d + %d = %d\n", a, b, a + b);
}

typedef struct Queue {
  int currentTask;
  int numTasksToDo;
  Task tasks[WORK_QUEUE_MAX];
} Queue;

Queue workQueue = {
  .currentTask = 0,
  .numTasksToDo = 0,
};

void executeTask(Task* task) {
  printf("Starting taskId: %d\n", task->id);
  switch(task->type) {
    case PRINT_ID:
      printId(task->id);
      break;
    case SUM:
      sum(task->sumArgs.a, task->sumArgs.b);
      break;
    default:
      // assert
      break;
  }
  printf("Completed taskId: %d\n", task->id);
}

void submitTask(Task task) {
  pthread_mutex_lock(&mutexQueue);
  // assert numTasksToDo < WORK_QUEUE_MAX
  int index = (workQueue.currentTask + workQueue.numTasksToDo) % WORK_QUEUE_MAX;
  workQueue.tasks[index] = task;
  workQueue.numTasksToDo++;
  pthread_mutex_unlock(&mutexQueue);
  pthread_cond_signal(&condQueue);
}

void startThread(void* args) {
  int index = *(int*)args;
  printf("Thread %d: Starting\n", index);
  while(1) {
    pthread_mutex_lock(&mutexQueue);
    while (workQueue.numTasksToDo == 0) {
      pthread_cond_wait(&condQueue, &mutexQueue);
    }

    Task task = workQueue.tasks[workQueue.currentTask];
    workQueue.numTasksToDo--;
    workQueue.currentTask++;
    // assert numTasksToDo >= 0
    // assert currentTask <= WORK_QUEUE_MAX
    if(workQueue.currentTask == WORK_QUEUE_MAX) {
      workQueue.currentTask = 0;
    }
    pthread_mutex_unlock(&mutexQueue);
    executeTask(&task);
  }
}

int main(int argc, char* argv[]) {
  srand(time(NULL));

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
  for(int i = 0; i < 1000; i++) {
    Task newTask;
    newTask.id = i;
    if(i % 2 == 0) {
      newTask.type = PRINT_ID;
    }
    else {
      newTask.type = SUM;
      newTask.sumArgs.a = rand() % 100;
      newTask.sumArgs.b = rand() % 100;
    }
    submitTask(newTask);
  }

  // threads never finish, so never join
  for(int i = 0; i < NUM_THREADS; i++) {
    if(pthread_join(threads[i], NULL)) {
      return 1;
    }
  }

  pthread_mutex_destroy(&mutexQueue);
  pthread_cond_destroy(&condQueue);

  return 0;
}
