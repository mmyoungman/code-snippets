/*
   mmy.h
   Last change: 9 Sep 2016

   001. (a) unsigned long stb_srand(unsigned long seed)
        (b) unsigned long stb_rand()
        (c) double stb_frand()
   Copied from https://github.com/nothings/stb/ (public domain). Uses Meresenne 
   Twister and LCG. (a) seeds the random number generator. (b) returns a random 
   number between 0 and ULONG_MAX. (c) returns a random number between 0 and 1. 
   Changed so automatically seeded with time(NULL) if srand() hasn't been called.

*/

// 001. START
#include <time.h>
#include <stdlib.h>
#include <string.h>

typedef struct { char d[4]; } stb__4;
typedef struct { char d[8]; } stb__8;

void stb_swap(void *p, void *q, size_t sz)
{
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

unsigned long stb_srandLCG(unsigned long seed)
{
    unsigned long previous = stb__rand_seed;
    stb__rand_seed = seed;
    return previous;
}

unsigned long stb_randLCG(void)
{
    stb__rand_seed = stb__rand_seed * 2147001325 + 715136305; // NOTE(stb): BCPL generator
    // shuffle non-random bits to the middle, and xor to decorrelate with seed
    return 0x31415926 ^ ((stb__rand_seed >> 16) + (stb__rand_seed << 16));
}

double stb_frandLCG(void)
{
    return stb_randLCG() / ((double) (1 << 16) * (1 << 16));
}

void stb_shuffle(void *p, size_t n, size_t sz, unsigned long seed)
{
    char *a;
    unsigned long old_seed;
    int i;
    if (seed)
        old_seed = stb_srandLCG(seed);
    a = (char *) p + (n-1) * sz;

    for (i=n; i > 1; --i) {
        int j = stb_randLCG() % i;
        stb_swap(a, (char *) p + j * sz, sz);
        a -= sz;
    }
    if (seed)
        stb_srandLCG(old_seed);
}

void stb_reverse(void *p, size_t n, size_t sz)
{
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

void stb_srand(unsigned long seed)
{
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

unsigned long stb_rand()
{
    if(!srandcalled)
    {
        stb_srand(time(NULL));
    }

    unsigned long * b = stb__mt_buffer;
    int idx = stb__mt_index;
    unsigned long s,r;
    int i;

    if (idx >= STB__MT_LEN*sizeof(unsigned long)) {
        if (idx > STB__MT_LEN*sizeof(unsigned long))
            stb_srand(0);
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

double stb_frand(void)
{
    return stb_rand() / ((double) (1 << 16) * (1 << 16));
}

// 001. END
