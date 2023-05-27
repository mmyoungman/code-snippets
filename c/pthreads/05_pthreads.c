// gcc 05_pthreads.c -o pthreads05 -lpthread

#include <pthread.h>
#include <stdio.h>
#include <unistd.h>

pthread_mutex_t mutexFuel;
pthread_cond_t condFuel;
int fuel = 0;

void fuelFill() {
  for(int i = 0; i < 5; i++) {
    pthread_mutex_lock(&mutexFuel);

    fuel += 15;
    printf("fuelFill: Fuel filled. Fuel: %d\n", fuel);

    pthread_mutex_unlock(&mutexFuel);
    pthread_cond_signal(&condFuel);

    sleep (1);
  }
}

void fuelUse() {
  pthread_mutex_lock(&mutexFuel);

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
}

int main(int argc, char* argv[]) {
  pthread_t threads[2];

  pthread_mutex_init(&mutexFuel, NULL);
  pthread_cond_init(&condFuel, NULL);

  if(pthread_create(&threads[0], NULL, (void*)&fuelFill, NULL)) {
    return 1;
  }
  if(pthread_create(&threads[1], NULL, (void*)&fuelUse, NULL)) {
    return 1;
  }

  for(int i = 0; i < 2; i++) {
    if(pthread_join(threads[i], NULL)) {
      return 1;
    }
  }

  pthread_mutex_destroy(&mutexFuel);
  pthread_cond_destroy(&condFuel);

  return 0;
}
