#include "lib-mmy.h"

u64 hash_chars(char* key) {
   int num_buckets = 16;
   u64 value = 0;
   u64 expo = 1;
   for(int i = 0; i < str_len(key); i++) {
      value += key[i] * expo;
      expo *= 10;
   }

   return (value % 2147483647) % num_buckets;
}

// This seems to be the best so far...
u64 hash_multrandfloat(char* key) {
   int num_buckets = 16;
   u64 value = 0;
   u64 expo = 1;
   for(int i = 0; i < str_len(key); i++) {
      value += key[i] * expo;
      expo *= 10;
   }

   f64 test = value * 10759857.185972389572938572;
   u64 testint = (u64)test;

   return (testint % 2147483647) % num_buckets;
}

u64 hash_multrand(char* key) {
   int num_buckets = 16;
   u64 value = 0;
   u64 expo = 1;
   for(int i = 0; i < str_len(key); i++) {
      value += key[i] * expo;
      expo *= 10;
   }

   value *= 2358729837;
   //value *= 1358395298358729837;

   return (value % 2147483647) % num_buckets;
}

u64 hash_crcvariant(char* key) {
   int num_buckets = 16;
   u64 value = 0;
   u64 expo = 1;
   for(int i = 0; i < str_len(key); i++) {
      value += key[i] * expo;
      expo *= 10;
   }

   u64 h = value;
   u64 highorder = h & 0xf8000000;
   h = h << 5;
   h = h ^ (highorder >> 27);

   return ((h ^ value) % 2147483647) % num_buckets;
}

int main(int argc, char** argv) {
   char* key = str_copy("aaaaaaaa");
   int sum_chars[16];
   for(int i = 0; i < 16; i++) {
      sum_chars[i] = 0;
   }
   for(int i = 0; i < 8; i++) {
      while(key[i] < 'z') {
         sum_chars[hash_chars(key)] += 1;
         key[i] += 1; 
      }
   }
   for(int i = 0; i < 16; i++) {
      dbg("hash_chars -- bucket: %d, count: %d", i, sum_chars[i]);
   }

   free(key);
   key = str_copy("aaaaaaaa");
   int sum_multrandfloat[16];
   for(int i = 0; i < 16; i++) {
      sum_multrandfloat[i] = 0;
   }
   for(int i = 0; i < 8; i++) {
      while(key[i] < 'z') {
         sum_multrandfloat[hash_multrandfloat(key)] += 1;
         key[i] += 1; 
      }
   }
   for(int i = 0; i < 16; i++) {
      dbg("hash_multrandfloat -- bucket: %d, count: %d", i, sum_multrandfloat[i]);
   }

   free(key);
   key = str_copy("aaaaaaaa");
   int sum_multrand[16];
   for(int i = 0; i < 16; i++) {
      sum_multrand[i] = 0;
   }
   for(int i = 0; i < 8; i++) {
      while(key[i] < 'z') {
         sum_multrand[hash_multrand(key)] += 1;
         key[i] += 1; 
      }
   }
   for(int i = 0; i < 16; i++) {
      dbg("hash_multrand-- bucket: %d, count: %d", i, sum_multrand[i]);
   }

   //free(key);
   //key = str_copy("aaaaaaaa");
   //int sum_crcvariant[16];
   //for(int i = 0; i < 16; i++) {
   //   sum_crcvariant[i] = 0;
   //}
   //for(int i = 0; i < 8; i++) {
   //   while(key[i] < 'z') {
   //      sum_crcvariant[hash_crcvariant(key)] += 1;
   //      key[i] += 1; 
   //   }
   //}
   //for(int i = 0; i < 16; i++) {
   //   dbg("hash_crcvariant -- bucket: %d, count: %d", i, sum_crcvariant[i]);
   //}
}
