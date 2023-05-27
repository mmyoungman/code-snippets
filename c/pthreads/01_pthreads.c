// gcc 01_pthreads.c -o pthreads01 -lpthread

#include <pthread.h>
#include <stdio.h>
#include <unistd.h>

void work() {
  printf("Starting thread\n");
  sleep(2);
  printf("Ending thread\n");
}

int main(int argc, char* argv[]) {
  pthread_t t1, t2;
  if(pthread_create(&t1, NULL, (void*)&work, NULL)) {
    return 1;
  }
  if(pthread_create(&t2, NULL, (void*)&work, NULL)) {
    return 1;
  }
  if(pthread_join(t1, NULL)) {
    return 1;
  }
  if(pthread_join(t2, NULL)) {
    return 1;
  }
  return 0;
}
