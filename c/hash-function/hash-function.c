#include "lib-mmy.h"

int num_buckets = 8;

u64 hash_chars(char* msg) {
   u64 value = 0;
   u64 expo = 1;
   for(int i = 0; i < str_len(msg); i++) {
      value += msg[i] * expo;
      expo *= 10;
   }

   return (value % 2147483647) % num_buckets;
}

u64 hash_multrandfloat(char* msg) {
   u64 value = 0;
   u64 expo = 1;
   for(int i = 0; i < str_len(msg); i++) {
      value += msg[i] * expo;
      expo *= 10;
   }

   f64 test = value * 10759857.185972389572938572;
   u64 testint = (u64)test;

   return (testint % 2147483647) % num_buckets;
}

u64 hash_multrand(char* msg) {
   u64 value = 0;
   u64 expo = 1;
   for(int i = 0; i < str_len(msg); i++) {
      value += msg[i] * expo;
      expo *= 10;
   }

   value *= 2358729837;
   //value *= 1358395298358729837;

   return (value % 2147483647) % num_buckets;
}

u64 hash_crcvariant(char* msg) {
   u64 value = 0;
   u64 expo = 1;
   for(int i = 0; i < str_len(msg); i++) {
      value += msg[i] * expo;
      expo *= 10;
   }

   u64 h = value;
   u64 highorder = h & 0xf8000000;
   h = h << 5;
   h = h ^ (highorder >> 27);

   return ((h ^ value) % 2147483647) % num_buckets;
}

//http://www.burtleburtle.net/bob/hash/hashfaq.html
u64 hash_bjenkins(char* msg) {
   u16 h = 0;
   for(int i = 0; i < str_len(msg); i++) {
      h += msg[i];
      h += (h << 10);
      h ^= (h >> 6);
   }
   h += (h << 3);
   h ^= (h >> 11);
   h += (h << 15);

   return h % num_buckets;
}

int main(int argc, char** argv) {
   char* msg = str_copy("aaaaaaaa");
   int total = 0;
   //int sum_chars[num_buckets];
   //for(int i = 0; i < num_buckets; i++) {
   //   sum_chars[i] = 0;
   //}
   //for(int i = 0; i < str_len(msg); i++) {
   //   while(msg[i] < 'z') {
   //      sum_chars[hash_chars(msg)] += 1;
   //      msg[i] += 1; 
   //   }
   //}
   //for(int i = 0; i < num_buckets; i++) {
   //   dbg("hash_chars -- bucket: %d, count: %d", i, sum_chars[i]);
   //}

   //free(msg);
   //msg = str_copy("aaaaaaaa");
   //int sum_multrandfloat[num_buckets];
   //for(int i = 0; i < num_buckets; i++) {
   //   sum_multrandfloat[i] = 0;
   //}
   //for(int i = 0; i < str_len(msg); i++) {
   //   while(msg[i] < 'z') {
   //      sum_multrandfloat[hash_multrandfloat(msg)] += 1;
   //      msg[i] += 1; 
   //      total += 1;
   //   }
   //}
   //for(int i = 0; i < num_buckets; i++) {
   //   dbg("hash_multrandfloat -- bucket: %d, count: %d", i, sum_multrandfloat[i]);
   //}
   //dbg("hash_multrandfloat total: %d", total);

   //free(msg);
   //msg = str_copy("aaaaaaaa");
   //total = 0;
   //int sum_multrand[num_buckets];
   //for(int i = 0; i < num_buckets; i++) {
   //   sum_multrand[i] = 0;
   //}
   //for(int i = 0; i < str_len(msg); i++) {
   //   while(msg[i] < 'z') {
   //      sum_multrand[hash_multrand(msg)] += 1;
   //      msg[i] += 1; 
   //      total += 1;
   //   }
   //}
   //for(int i = 0; i < num_buckets; i++) {
   //   dbg("hash_multrand-- bucket: %d, count: %d", i, sum_multrand[i]);
   //}
   //dbg("hash_multrand total: %d", total);
   //total = 0;
   //for(int i = 0; i < num_buckets; i++) {
   //   total += sum_multrand[i];
   //}
   //dbg("hash_multrand calculated total: %d", total);

   //free(msg);
   //msg = str_copy("aaaaaaaa");
   //int sum_crcvariant[num_buckets];
   //for(int i = 0; i < num_buckets; i++) {
   //   sum_crcvariant[i] = 0;
   //}
   //for(int i = 0; i < str_len(msg); i++) {
   //   while(msg[i] < 'z') {
   //      sum_crcvariant[hash_crcvariant(msg)] += 1;
   //      msg[i] += 1; 
   //   }
   //}
   //for(int i = 0; i < num_buckets; i++) {
   //   dbg("hash_crcvariant -- bucket: %d, count: %d", i, sum_crcvariant[i]);
   //}

   free(msg);
   msg = str_copy("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");
   total = 0;
   int sum_bjenkins[num_buckets];
   for(int i = 0; i < num_buckets; i++) {
      sum_bjenkins[i] = 0;
   }
   for(int i = 0; i < str_len(msg); i++) {
      while(msg[i] < 'z') {
         sum_bjenkins[hash_bjenkins(msg)] += 1;
         msg[i] += 1; 
         total += 1;
      }
   }
   for(int i = 0; i < num_buckets; i++) {
      dbg("hash_bjenkins -- bucket: %d, count: %d", i, sum_bjenkins[i]);
   }
   dbg("hash_bjenkins total: %d", total);
   total = 0;
   for(int i = 0; i < num_buckets; i++) {
      total += sum_bjenkins[i];
   }
   dbg("hash_bjenkins calculated total: %d", total);
}
