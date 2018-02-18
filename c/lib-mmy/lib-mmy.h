/*
   lib-mmy.h
   Last change: 18 Feb 2018

   000. 
      (a) Typedefs
      (b) Useful macros
      (c) Assert/Debug macros
      (d) Logging macros

   001.
      (a) void memset(unsigned char *ptr, unsigned char value, u64 size)

   002. 
   Copied from https://github.com/nothings/stb/ (public domain). 
   Uses Meresenne Twister and LCG to seed. Changed so automatically 
   seeded with time(NULL) if srand() hasn't been called.
   (a) seeds the random number generator. 
   (b) returns a random number between 0 and ULONG_MAX. 
   (c) returns a random number between 0 and 1. 
      (a) unsigned long stb_srand(unsigned long seed)
      (b) unsigned long stb_rand()
      (c) double stb_frand()

   003.
   Math operations. a,b use intrinsics. c,d,e: 
   https://graphics.stanford.edu/%7Eseander/bithacks.html
      (a) float mth_sqrt(float input)
      (b) double mth_sqrt(double input)
      (c) int mth_min(int a, int b)
      (d) int mth_max(int a, int b)
      (e) int mth_abs(int a)
      (f) int mth_pow(int num, int pow)

   004. 
   ANSI string operations.
      (a) int str_len(char* str)
      (b) int str_equal(char *a, char *b)
      (c) void str_copy(char *s, char *copy)
      (d) char* str_copy(char *s)
      (e) int str_beginswith(char *str, char *start)
      (f) int str_endswith(char *str, char *end)
      (g) char* str_concat(char *str, char *addition)
      (h) void str_lower(char* str)
      (i) void str_upper(char* str)
      (j) int str_isalpha(char* str)
      (k) int str_isint(char* str)
      (l) char* str_lstrip(char* str, char tostrip)
      (m) char* str_rstrip(char* str, char tostrip)
      (n) void str_sort(char* str)
      (o) char** str_split(char* str, char c, int* size)
      (p) int str_toint(char* str)
      (q) char* str_inttostr(int num)

*/

// 000. START 
#if 1

#include <stdio.h>
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

#define kilobytes(value) ((value)*(u64)1024)
#define megabytes(value) (kilobytes(value)*1024)
#define gigabytes(value) (megabytes(value)*1024)
#define terabytes(value) (gigabytes(value)*1024)

#define alloc(size) calloc(1, size) // calloc so memory is zeroed

#ifdef DEBUG
#define dbg(msg, ...) fprintf(stderr, "[DEBUG] (%s:%d) " msg "\n", \
                              __FILE__, __LINE__, ##__VA_ARGS__)
#define assert(expr) if(!(expr)) { dbg("Assert failed: " #expr); *(int*)0 = 0; }
#else
#define dbg(msg, ...)
#define assert(expr)
#endif


#define log_err(msg, ...) fprintf(stderr, "[ERROR] (%s:%d) " msg "\n", \
                                  __FILE__, __LINE__, ##__VA_ARGS__) 
#define log_warn(msg, ...) fprintf(stderr, "[WARN] (%s:%d) " msg "\n", \
                                   __FILE__, __LINE__, ##__VA_ARGS__)
#define log_info(msg, ...) fprintf(stderr, "[INFO] (%s:%d) " msg "\n", \
                                   __FILE__, __LINE__, ##__VA_ARGS__)

#endif
// 000. END

// 001. START
#if 1
// NOTE: Untested
void memset(unsigned char *ptr, unsigned char value, u64 size) {
   for(u64 i = 0; i < size; i++) {
      *ptr = value;
      ptr++;
   }
}
#endif
// 001. END

// 002. START
#if 1

#include <stdlib.h>
#include <time.h>
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
// 002. END

// 003. START
#if 1

#include <emmintrin.h>

//float mth_sqrt(float input) {
//    float result = 0.0f;
//    __m128 one = _mm_set_ss(input);
//    __m128 two = _mm_sqrt_ss(one); // SSE
//    result = _mm_cvtss_f32(two);
//
//    return result;
//}

double mth_sqrt(double input) {
    double result = 0.0f;
    __m128d one = _mm_set_sd(input);
    __m128d two = _mm_sqrt_pd(one); // SSE2
    result = _mm_cvtsd_f64(two);

     return result;
}

int mth_min(int a, int b) {
    int result = b ^ ((a ^ b) & -(a < b));

    return result;
}

int mth_max(int a, int b) {
    int result = a ^ ((a ^ b) & -(a < b));

    return result;
}

int mth_abs(int a) {
     unsigned int result;
     int const mask = a >> sizeof(int) * 8 - 1;
     result = (a + mask) ^ mask;

    return result;
}

int mth_pow(int num, int pow) {
    int result = 1;
    for(int i = 0; i < pow; i++) {
        result *= num;
    }
    return result;
}

#endif
// 003. END

// 004. START
#if 1

//#include <stdlib.h>

int str_len(char *str) {
     char* ptr = str;
     while(*ptr != 0)
          ptr++;
     return ptr - str;
}

int str_equal(char *a, char *b) {
  while((*a != '\0') && (*a == *b)) {
    a++, b++;
  }
  return ((*a == '\0') && (*b == '\0'));
}

//void str_copy(char *s, char *copy) {
//     while(*s != '\0') {
//          *copy = *s;
//          s++, copy++;
//     }
//     *copy = '\0';
//}

char* str_copy(char *s) {
  char* copy = (char*)alloc(sizeof(char)*(str_len(s)+1));
  char* copyPtr = copy;
  while(*s != '\0') {
    *copyPtr = *s;
    s++, copyPtr++;
  }
  *copyPtr = '\0';
  return copy;
}

int str_beginswith(char* str, char* start) {
  while((*start != '\0') && (*start == *str)) {
    start++, str++;
  }
  return *start == '\0';
}

int str_endswith(char* str, char* end) {
    int endLength = str_len(end);
    while(*end != '\0') { end++; }
    while(*str != '\0') { str++; }

    while(*str == *end && endLength > 0) {
        str--, end--;
        endLength--;
    }
    return *str == *end;
}

char* str_concat(char* str, char* addition) {
    int newLength = str_len(str) + str_len(addition) + 1;
    str = (char*)realloc(str, sizeof(char) * newLength);
    char* strPtr = str;
    while(*strPtr != '\0') { strPtr++; }
    while(*addition != '\0') {
        *strPtr = *addition;
        strPtr++, addition++;
    }
    *strPtr = '\0';
    return str;
}

void str_lower(char* str) {
  while(*str != '\0') {
    if(*str >= 'A' && *str <= 'Z') {
      *str += 'a' - 'A';
    }
    str++;
  }
}

void str_upper(char* str) {
  while(*str != '\0') {
    if(*str >= 'a' && *str <= 'z') {
      *str -= 'a' - 'A';
    }
    str++;
  }
}

int str_isalpha(char* str) {
   while(*str != '\0') {
      if((*str < 'a' || *str > 'z') && (*str < 'A' || *str > 'Z')) {
         return 0;
      }
      str++;
   }
   return 1;
}

int str_isint(char* str) {
   if(*str == '-') {
      str++;
   }
   while(*str != '\0') {
      if(*str < '0' || *str > '9') {
         return 0;
      }
      str++;
   }
   return 1;
}

char* str_lstrip(char* str, char tostrip) {
   while(*str == tostrip) {
      str++;
   }
   return str;
}

char* str_rstrip(char* str, char tostrip) {
   char* strPtr = str;
   while(strPtr[1] != '\0') {
      strPtr++;
   }
   while(*strPtr == tostrip) {
      *strPtr = '\0';
      strPtr--;
   }
   return str;
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
  char** result = (char**)alloc(sizeof(char*)*numStrs);
  strPtr = str;
  int i = 0;
  while(numStrs > 0) {
    if(*strPtr == '\0') {
      numStrs--;
      result[i] = str;
      str = strPtr + 1;
      i++;
    }
    strPtr++;
  }
  *size = i;
  return result;
}

int str_toint(char *str) {
   int result = 0;
   char *strPtr = str;

   int length = str_len(str);

   while (length > 0) {
      length--;
      if (*strPtr < '0' || *strPtr > '9') {
         strPtr++;
         continue;
      } else {
         // Calculate value based on position (i.e. value * 10^position)
         int exponent = 1;
         for (int i = 0; i < length; i++) {
            exponent *= 10;
         }
         result += (*strPtr - 48) * exponent;
         strPtr++;
      }
   }
   if (str[0] == '-')
      result = -result;

   return result;
}

char* str_inttostr(int num) {
   int negative = 0;
   if (num < 0) {
      negative = 1;
      num = -num;
   }

   int len = 1;
   int temp = num;
   while (temp >= 10) {
      len++;
      temp /= 10;
   }

   char *res = (char *)alloc(sizeof(char) *
                              (len + negative + 1)); // extra for '-' and '\0'

   char *resPtr = res;
   if (negative) {
      *resPtr = '-';
      resPtr++;
   }
   while (len > 0) {
      int mod = 1;
      int templen = len;
      while (templen > 0) {
         mod *= 10;
         templen--;
      }
      *resPtr = '0' + ((num % mod) / (mod / 10)); // isolate wanted digit
      resPtr++;
      len--;
   }

   *resPtr = '\0';
   return res;
}

#endif
// 004. END
