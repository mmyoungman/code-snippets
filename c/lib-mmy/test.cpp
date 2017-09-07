#include "lib-mmy.h"

#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <limits.h> // for INT_MIN etc.

int main()
{
  int size;

  char* testStr = stringCopy("This:Is:A:Test:To:Use:With:Split");
  printf("testStr: \"%s\"\n", testStr);
  char** split = stringSplit(testStr, ':', &size);
  for(int i = 0; i < size; i++) {
    printf("split[%d]: \"%s\"\n", i, split[i]);
  }

  printf("\n");

  testStr = stringCopy("???Another?Test?For????Split??");
  printf("testStr: \"%s\"\n", testStr);

  split = stringSplit(testStr, '?', &size);
  for(int i = 0; i < size; i++) {
    printf("split[%d]: \"%s\"\n", i, split[i]);
  }

  printf("\n");

  testStr = stringCopy("YetAnotherTestForSplit");
  printf("testStr: \"%s\"\n", testStr);
  split = stringSplit(testStr, 'e', &size);
  for(int i = 0; i < size; i++) {
    printf("split[%d]: \"%s\"\n", i, split[i]);
  }

    /*
    int a = 12;
    int b = 1;

    printf("min: %d, max: %d\n", mmy_min(a, b), mmy_max(a, b));

    a = -112;
    b = -100;

    printf("abs(a): %d, abs(b): %d\n", mmy_abs(a), mmy_abs(b));
    */
    //u16 one = 12;
    //printf("one: %d\n", one);

    //f64 d = 15.0f;
    //printf("sqrt(15): %f\n", mmy_sqrt(d));

    //f32 f = 15.0f;
    //printf("sqrt(15): %f\n", mmy_sqrt(f));
    //
    //d = 121.0f;
    //printf("sqrt(121): %f\n", mmy_sqrt(d));

    //printf("RAND_MAX: %d\n", RAND_MAX);

    //unsigned long x = 0;
    //printf("ULONG --0: %lu\n", --x);
    //printf("ULONG_MAX: %lu\n", ULONG_MAX);

    //x = stb_rand();
    //printf("x:         %lu\n", x);

    //x = stb_rand();
    //printf("x:         %lu\n", x);

    //x = stb_rand();
    //printf("x:         %lu\n", x);

    //double d = 0;

    ////stb_srand(time(NULL));

    //d = stb_frand();
    //printf("d: %f\n", d);

    //d = stb_frand();
    //printf("d: %f\n", d);

    //d = stb_frand();
    //printf("d: %f\n", d);
}
