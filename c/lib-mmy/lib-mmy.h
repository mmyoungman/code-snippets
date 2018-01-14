/*
   lib-mmy.h
   Last change: 14 Jan 2018

   000. (a) Type defines
        (b) Assert macro

   001. (a) unsigned long stb_srand(unsigned long seed)
        (b) unsigned long stb_rand()
        (c) double stb_frand()
   Copied from https://github.com/nothings/stb/ (public domain). 
   Uses Meresenne Twister and LCG to seed. Changed so automatically 
   seeded with time(NULL) if srand() hasn't been called.
   (a) seeds the random number generator. 
   (b) returns a random number between 0 and ULONG_MAX. 
   (c) returns a random number between 0 and 1. 

   002. (a) float mathSqrt(float input)
        (b) double mathSqrt(double input)
        (c) int mathMin(int a, int b)
        (d) int mathMax(int a, int b)
        (e) int mathAbs(int a)
        (f) int mathPower(int num, int pow)
   Math operations. a,b use intrinsics. c,d,e: 
   https://graphics.stanford.edu/%7Eseander/bithacks.html

   003. (a) int str_len(char* str)
        (b) void str_copy(char *s, char *copy)
        (c) char* str_copy(char *s)
        (d) int str_equal(char *a, char *b)
        (e) int str_beginswith(char *start, char *str)
        (f) int str_endswith(char *end, char *str) // Untested
        (g) char* str_concat(char *str, char *addition)
        (h) void str_lower(char* str)
        (i) void str_upper(char* str)
        (j) void str_sort(char* str)
        (k) char** str_split(char* str, char c, int* size)
        (l) int str_toint(char* str)
   ANSI string operations.

*/

// 000. START 
#if 0

#include <stdint.h>

typedef uint8_t u8;
typedef uint16_t u16;
typedef uint32_t u32;
typedef uint64_t u64;

typedef int8_t s8; 
typedef int16_t s16; 
typedef int32_t s32; 
typedef int64_t s64; 

typedef float f32;
typedef double f64;

#define Kilobytes(Value) ((Value)*1024)
#define Megabytes(Value) (Kilobytes(Value)*1024)
#define Gigabytes(Value) (Megabytes(Value)*1024)

#if DEBUG
#define assert(expression) if(!(expression)) { *(int*)0 = 0; }
#else
#define assert(expression)
#endif

#endif
// 000. END

// 001. START
#if 0

#include <time.h>
#include <stdlib.h>
#include <string.h>

typedef struct { char d[4]; } stb__4;
typedef struct { char d[8]; } stb__8;

void stb_swap(void *p, void *q, size_t sz) {
    char buffer[256];
    if (p == q) return;
    if (sz == 4) {
        stb__4 temp    = * ( stb__4 *) p;
        * (stb__4 *) p = * ( stb__4 *) q;
        * (stb__4 *) q = temp;
        return;
    } else if (sz == 8) {
        stb__8 temp    = * ( stb__8 *) p;
        * (stb__8 *) p = * ( stb__8 *) q;
        * (stb__8 *) q = temp;
        return;
    }

    while (sz > sizeof(buffer)) {
        stb_swap(p, q, sizeof(buffer));
        p = (char *) p + sizeof(buffer);
        q = (char *) q + sizeof(buffer);
        sz -= sizeof(buffer);
    }

    memcpy(buffer, p     , sz);
    memcpy(p     , q     , sz);
    memcpy(q     , buffer, sz);
}

static unsigned long stb__rand_seed=0;

unsigned long stb_srandLCG(unsigned long seed) {
    unsigned long previous = stb__rand_seed;
    stb__rand_seed = seed;
    return previous;
}

unsigned long stb_randLCG(void) {
    stb__rand_seed = stb__rand_seed * 2147001325 + 715136305; // NOTE(stb): BCPL generator
    // shuffle non-random bits to the middle, and xor to decorrelate with seed
    return 0x31415926 ^ ((stb__rand_seed >> 16) + (stb__rand_seed << 16));
}

double stb_frandLCG(void) {
    return stb_randLCG() / ((double) (1 << 16) * (1 << 16));
}

void stb_shuffle(void *p, size_t n, size_t sz, unsigned long seed) {
    char *a;
    unsigned long old_seed;
    int i;
    if (seed) {
        old_seed = stb_srandLCG(seed);
    }
    a = (char *) p + (n-1) * sz;

    for (i=n; i > 1; --i) {
        int j = stb_randLCG() % i;
        stb_swap(a, (char *) p + j * sz, sz);
        a -= sz;
    }
    if (seed) {
        stb_srandLCG(old_seed);
    }
}

void stb_reverse(void *p, size_t n, size_t sz) {
    int i,j = n-1;
    for (i=0; i < j; ++i,--j) {
        stb_swap((char *) p + i * sz, (char *) p + j * sz, sz);
    }
}

// NOTE(stb): public domain Mersenne Twister by Michael Brundage
#define STB__MT_LEN       624

int stb__mt_index = STB__MT_LEN*sizeof(unsigned long)+1;
unsigned long stb__mt_buffer[STB__MT_LEN];

int srandcalled = 0;

void stb_srand(unsigned long seed) {
    srandcalled = 1;

    int i;
    unsigned long old = stb_srandLCG(seed);
    for (i = 0; i < STB__MT_LEN; i++)
        stb__mt_buffer[i] = stb_randLCG();
    stb_srandLCG(old);
    stb__mt_index = STB__MT_LEN*sizeof(unsigned long);
}

#define STB__MT_IA           397
#define STB__MT_IB           (STB__MT_LEN - STB__MT_IA)
#define STB__UPPER_MASK      0x80000000
#define STB__LOWER_MASK      0x7FFFFFFF
#define STB__MATRIX_A        0x9908B0DF
#define STB__TWIST(b,i,j)    ((b)[i] & STB__UPPER_MASK) | ((b)[j] & STB__LOWER_MASK)
#define STB__MAGIC(s)        (((s)&1)*STB__MATRIX_A)

unsigned long stb_rand() {
    if(!srandcalled) {
        stb_srand(time(NULL));
    }

    unsigned long * b = stb__mt_buffer;
    int idx = stb__mt_index;
    unsigned long s,r;
    int i;

    if (idx >= STB__MT_LEN*sizeof(unsigned long)) {
        if (idx > STB__MT_LEN*sizeof(unsigned long)) {
            stb_srand(0);
        }
        idx = 0;
        i = 0;
        for (; i < STB__MT_IB; i++) {
            s = STB__TWIST(b, i, i+1);
            b[i] = b[i + STB__MT_IA] ^ (s >> 1) ^ STB__MAGIC(s);
        }
        for (; i < STB__MT_LEN-1; i++) {
            s = STB__TWIST(b, i, i+1);
            b[i] = b[i - STB__MT_IB] ^ (s >> 1) ^ STB__MAGIC(s);
        }

        s = STB__TWIST(b, STB__MT_LEN-1, 0);
        b[STB__MT_LEN-1] = b[STB__MT_IA-1] ^ (s >> 1) ^ STB__MAGIC(s);
    }
    stb__mt_index = idx + sizeof(unsigned long);

    r = *(unsigned long *)((unsigned char *)b + idx);

    r ^= (r >> 11);
    r ^= (r << 7) & 0x9D2C5680;
    r ^= (r << 15) & 0xEFC60000;
    r ^= (r >> 18);

    return r;
}

double stb_frand(void) {
    return stb_rand() / ((double) (1 << 16) * (1 << 16));
}

#endif
// 001. END

// 002. START
#if 0

#include <emmintrin.h>

float mathSqrt(float input) {
    float result = 0.0f;
    __m128 one = _mm_set_ss(input);
    __m128 two = _mm_sqrt_ss(one); // SSE
    result = _mm_cvtss_f32(two);

    return result;
}

double mathSqrt(double input) {
    double result = 0.0f;
    __m128d one = _mm_set_sd(input);
    __m128d two = _mm_sqrt_pd(one); // SSE2
    result = _mm_cvtsd_f64(two);

     return result;
}

int mathMin(int a, int b) {
    int result = b ^ ((a ^ b) & -(a < b));

    return result;
}

int mathMax(int a, int b) {
    int result = a ^ ((a ^ b) & -(a < b));

    return result;
}

int mathAbs(int a) {
     unsigned int result;
     int const mask = a >> sizeof(int) * 8 - 1;
     result = (a + mask) ^ mask;

    return result;
}

int mathPower(int num, int pow) {
    int result = 1;
    for(int i = 0; i < pow; i++) {
        result *= num;
    }
    return result;
}

#endif
// 002. END

// 003. START
#if 1

#include <stdlib.h>

int str_len(char *str) {
     char* ptr = str;
     while(*ptr != 0)
          ptr++;
     return ptr - str;
}

void str_copy(char *s, char *copy) {
     while(*s != '\0') {
          *copy = *s;
          s++, copy++;
     }
     *copy = '\0';
}

char* str_copy(char *s) {
  char* copy = (char*)malloc(sizeof(char)*(str_len(s)+1));
  char* sPtr = s;
  char* copyPtr = copy;
  while(*sPtr != '\0') {
    *copyPtr = *sPtr;
    sPtr++, copyPtr++;
  }
  *copyPtr = '\0';
  return copy;
}

int str_equal(char *a, char *b) {
  while((*a != '\0') && (*a == *b)) {
    a++, b++;
  }
  return ((*a == '\0') && (*b == '\0'));
}

int str_beginswith(char* a, char *str) {
  while((*a != '\0') && (*a == *str))
    a++, str++;
  return *a == '\0';
}

int str_endswith(char* str, char* end) {
    char* strPtr = str;
    int endLength = str_len(end);
    while(*end != '\0') { end++; }
    while(*strPtr != '\0') { strPtr++; }

    while(*strPtr == *end && endLength > 0) {
        strPtr--, end--;
        endLength--;
    }
    return *strPtr == *end;
}

void str_concat(char* str, char* addition) {
    int newLength = str_len(str) + str_len(addition) + 1;
    str = (char*)realloc(str, sizeof(char) * newLength);
    char* strPtr = str;
    while(*strPtr != '\0') { strPtr++; }
    char* addPtr = addition;
    while(*addPtr != '\0') {
        *strPtr = *addPtr;
        strPtr++, addPtr++;
    }
    *strPtr = '\0';
}

void str_lower(char* str) {
  char* strPtr = str;
  while(*strPtr != '\0') {
    if(*strPtr >= 'A' && *strPtr <= 'Z') {
      *strPtr += 'a' - 'A';
    }
    strPtr++;
  }
}

void str_upper(char* str) {
  char* strPtr = str;
  while(*strPtr != '\0') {
    if(*strPtr >= 'a' && *strPtr <= 'z') {
      *strPtr -= 'a' - 'A';
    }
    strPtr++;
  }
}

void str_sort(char* str) {
   int len = str_len(str);
   for(int i = 0; i < len; i++) {
      for(int j = i+1; j < len; j++) {
         if(str[i] > str[j]) {
            char temp = str[i];
            str[i] = str[j];
            str[j] = temp;
         }
      }
   }
}

char** str_split(char* str, char c, int* size) {
  int numStrs = 1;
  char* strPtr = str;
  while(*strPtr != '\0') {
    if(*strPtr == c) {
      *strPtr = '\0';
      numStrs++;
    }
    strPtr++;
  }
  char** result = (char**)malloc(sizeof(char*)*numStrs);
  char* strStart = str;
  strPtr = strStart;
  int i = 0;
  while(numStrs > 0) {
    if(*strPtr == '\0') {
      numStrs--;
      result[i] = strStart;
      strStart = strPtr + 1;
      i++;
    }
    strPtr++;
  }
  *size = i;
  return result;
}

int str_toint(char* str) {
  int result = 0;
  char* strPtr = str;

  int length = str_len(str);
  while(length > 0) {
    length--;
    // Calculate value based on position (i.e. value * 10^position)
    int exponent = 1;
    for(int i = 0; i < length; i++) {
        exponent *= 10;
    }
    result += (*strPtr - 48) * exponent;
    strPtr++;
  }
  return result;
}

#endif
// 003. END
