// gcc 10_pthreads.c -o pthreads10 -lpthread

#include <assert.h>
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <unistd.h>

#define NUM_THREADS 4
#define WORK_QUEUE_MAX 1024

pthread_mutex_t mutexQueue;
pthread_cond_t condNotEmpty;
pthread_cond_t condNotFull;

typedef enum TaskType {
  PRINT_ID,
  SUM,
} TaskType;

// Every TaskType gets its own args struct, even a single-field one. Reaching
// into Task.id for a payload would conflate queue bookkeeping with task data,
// and leaves nowhere to put a second field later.
typedef struct PrintIdArgs {
  int value;
} PrintIdArgs;

typedef struct SumArgs {
  int a, b;
} SumArgs;

// NOTE: a union costs every task the size of its largest member, and the queue
// is WORK_QUEUE_MAX of these. Fine while the variants are small (16 bytes each
// here); if one ever needs a big inline buffer, give that variant a pointer to
// its own storage rather than inflating every slot.
typedef struct Task {
  int id;        // queue bookkeeping, never task payload
  TaskType type;
  union {
    PrintIdArgs printIdArgs;
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

// Ring buffer: same queue as 09, but dequeue is O(1) instead of shifting the
// whole array down by one.
typedef struct Queue {
  int currentTask;
  int numTasksToDo;
  Task tasks[WORK_QUEUE_MAX];
} Queue;

Queue workQueue = {
  .currentTask = 0,
  .numTasksToDo = 0,
};

int shuttingDown = 0;

// Each case returns, so falling out of the switch means the discriminant
// matched nothing - that is what the assert catches. The missing `default:`
// label is deliberate: without one, -Wswitch fails the build naming any
// TaskType added to the enum but not handled here. Adding `default:` would
// silence that check and demote it to a runtime assert.
void dispatchTask(Task* task) {
  switch(task->type) {
    case PRINT_ID:
      printId(task->printIdArgs.value);
      return;
    case SUM:
      sum(task->sumArgs.a, task->sumArgs.b);
      return;
  }
  assert(0 && "unknown TaskType");
}

void executeTask(Task* task) {
  printf("Starting taskId: %d\n", task->id);
  dispatchTask(task);
  printf("Completed taskId: %d\n", task->id);
}

void submitTask(Task task) {
  pthread_mutex_lock(&mutexQueue);

  // Back-pressure: without this the ring buffer wraps onto tasks that have not
  // been consumed yet. No crash, no warning - the work just silently vanishes.
  while(workQueue.numTasksToDo == WORK_QUEUE_MAX) {
    pthread_cond_wait(&condNotFull, &mutexQueue);
  }
  assert(workQueue.numTasksToDo < WORK_QUEUE_MAX);

  int index = (workQueue.currentTask + workQueue.numTasksToDo) % WORK_QUEUE_MAX;
  workQueue.tasks[index] = task;
  workQueue.numTasksToDo++;
  pthread_mutex_unlock(&mutexQueue);
  pthread_cond_signal(&condNotEmpty);
}

void requestShutdown(void) {
  pthread_mutex_lock(&mutexQueue);
  shuttingDown = 1;
  pthread_mutex_unlock(&mutexQueue);
  // NOTE: broadcast, not signal - every idle worker must wake and see the flag
  pthread_cond_broadcast(&condNotEmpty);
}

void* startThread(void* args) {
  int index = *(int*)args;
  printf("Thread %d: Starting\n", index);
  while(1) {
    pthread_mutex_lock(&mutexQueue);
    while(workQueue.numTasksToDo == 0 && !shuttingDown) {
      pthread_cond_wait(&condNotEmpty, &mutexQueue);
    }
    // Drain first, then exit: only stop once the queue is actually empty.
    if(workQueue.numTasksToDo == 0) {
      pthread_mutex_unlock(&mutexQueue);
      break;
    }

    Task task = workQueue.tasks[workQueue.currentTask];
    workQueue.numTasksToDo--;
    workQueue.currentTask++;
    assert(workQueue.numTasksToDo >= 0);
    assert(workQueue.currentTask <= WORK_QUEUE_MAX);
    if(workQueue.currentTask == WORK_QUEUE_MAX) {
      workQueue.currentTask = 0;
    }
    pthread_mutex_unlock(&mutexQueue);
    pthread_cond_signal(&condNotFull);

    executeTask(&task);
  }
  printf("Thread %d: Exiting\n", index);
  return NULL;
}

int main(void) {
  srand(time(NULL));

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

  // Generate tasks. Now that submitTask applies back-pressure this count can
  // exceed WORK_QUEUE_MAX safely - try 100000 to watch the producer block.
  for(int i = 0; i < 1000; i++) {
    Task newTask;
    newTask.id = i;
    if(i % 2 == 0) {
      newTask.type = PRINT_ID;
      // Same number as newTask.id here, but it is PRINT_ID's own payload -
      // id stays purely queue bookkeeping.
      newTask.printIdArgs.value = i;
    }
    else {
      newTask.type = SUM;
      newTask.sumArgs.a = rand() % 100;
      newTask.sumArgs.b = rand() % 100;
    }
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
