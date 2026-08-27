// gcc 05_pthreads.c -o pthreads05 -lpthread

#include <pthread.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>

pthread_mutex_t mutexFuel;
pthread_cond_t condFuel;
int fuel = 0;

void* fuelFill(void* arg) {
  (void)arg;
  for(int i = 0; i < 5; i++) {
    pthread_mutex_lock(&mutexFuel);

    fuel += 15;
    printf("fuelFill: Fuel filled. Fuel: %d\n", fuel);

    pthread_mutex_unlock(&mutexFuel);
    pthread_cond_signal(&condFuel);

    sleep(1);
  }
  return NULL;
}

void* fuelUse(void* arg) {
  (void)arg;
  pthread_mutex_lock(&mutexFuel);

  // NOTE: always wait in a loop on the predicate, never a bare if. Two reasons:
  // the signal may arrive before we start waiting, and pthread_cond_wait is
  // allowed to wake spuriously.
  while(fuel < 40) {
    printf("fuelUse: Not enough fuel. Waiting...\n");

    // NOTE: pthread_cond_wait unlocks mutexFuel and
    // waits for signal from condFuel AND for mutexFuel to be unlocked
    pthread_cond_wait(&condFuel, &mutexFuel);
    printf("fuelUse: condFuel triggered\n");
  }

  fuel -= 40;
  printf("fuelUse: Used fuel. Fuel: %d\n", fuel);

  pthread_mutex_unlock(&mutexFuel);
  return NULL;
}

int main(void) {
  pthread_t threads[2];
  int rc;

  if((rc = pthread_mutex_init(&mutexFuel, NULL))) {
    fprintf(stderr, "pthread_mutex_init: %s\n", strerror(rc));
    return 1;
  }
  if((rc = pthread_cond_init(&condFuel, NULL))) {
    fprintf(stderr, "pthread_cond_init: %s\n", strerror(rc));
    return 1;
  }

  if((rc = pthread_create(&threads[0], NULL, fuelFill, NULL))) {
    fprintf(stderr, "pthread_create: %s\n", strerror(rc));
    return 1;
  }
  if((rc = pthread_create(&threads[1], NULL, fuelUse, NULL))) {
    fprintf(stderr, "pthread_create: %s\n", strerror(rc));
    return 1;
  }

  for(int i = 0; i < 2; i++) {
    if((rc = pthread_join(threads[i], NULL))) {
      fprintf(stderr, "pthread_join: %s\n", strerror(rc));
      return 1;
    }
  }

  pthread_mutex_destroy(&mutexFuel);
  pthread_cond_destroy(&condFuel);

  return 0;
}
