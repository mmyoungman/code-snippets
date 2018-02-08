#include "lib-mmy.h"

int main()
{

#ifdef DEBUG
   printf("\nStarting debug tests...\n\n");
#else
   printf("\nStarting non-debug tests..\n\n");
#endif

   // Test 000.
   assert(1024 == kilobytes(1));
   assert((1024*1024) == megabytes(1));
   assert(((uint64)10*1024*1024*1024) == gigabytes(10));
   dbg("gigabytes(10): %ll", gigabytes(10));

   // Test 001.
   log_err("Log error: %d, %s", 42, "string literal");
   log_warn("Log warning: %d, %s", 12, "another string literal");
   log_info("Log info: %d, %s", 13, "yet another string literal");
   
   dbg("Test msg: %d, %s", 14, "not another bloody string literal");


   // Test 002.


   // Test 003.
   assert(mth_min(1, 2) == 1);
   assert(mth_max(1, 2) == 2);


   // Test 004.
   assert(str_len("123") == 3);
   assert(str_len("") == 0);
   assert(str_len("0123456789") == 10);

   assert(str_equal("123", "123"));
   assert(!str_equal("123", "1234"));
   assert(!str_equal("12345", "1234"));

   char *str = "a string";
   char *copy = str_copy(str);
   copy[0] = 'a';
   copy[1] = ' ';
   assert(str_equal(str, copy));
   dbg("str: %s, copy: %s", str, copy);
   free(copy);

   //str = "different string, much longer than previous string";
   //str_copy(str, copy);
   //dbg("str: %s, copy: %s", str, copy);
   //assert(str_equal(str, copy));

   assert(str_beginswith("a stri", str));
   assert(str_endswith(str, "tring"));

   str = str_copy("Test ");
   str = str_concat(str, "string");
   assert(str_equal(str, "Test string"));
   free(str);

   str = str_copy("SHoUT");
   str_lower(str);
   assert(str_equal(str, "shout"));
   free(str);

   str = str_copy("quIet");
   str_upper(str);
   assert(str_equal(str, "QUIET"));
   free(str);

   str = str_copy("aaalksfjnhekwjbegjabwegij");
   dbg("isalpha(str): %d, str: %s", str_isalpha(str), str);
   assert(str_isalpha(str));
   str[10] = '#';
   dbg("isalpha(str): %d, str: %s", str_isalpha(str), str);
   assert(!str_isalpha(str));
   free(str);

   str = str_copy("-190572985672019359876298");
   assert(str_isint(str));
   free(str);

   str = str_copy("eeeeeeeeeeeklhgfgh         ");
   assert(str_equal(str_rstrip(str, ' '), "eeeeeeeeeeeklhgfgh"));
   assert(str_equal(str_lstrip(str, 'e'), "klhgfgh"));


   str = str_copy("agfecbd");
   str_sort(str);
   assert(str_equal(str, "abcdefg"));
   free(str);

   str = str_copy("This:Is::A:Test:To:Use:With:Split::");
   int size;
   char** split = str_split(str, ':', &size);
   assert(size == 11);
   assert(str_equal(split[0], "This"));
   assert(str_equal(split[5], "To"));
   assert(str_equal(split[8], "Split"));
   assert(str_equal(split[9], ""));
   assert(str_equal(split[10], ""));
   
   dbg("str_toint(\"1234\"): %d", str_toint("1234"));
   dbg("str_toint(\"-12345\"): %d", str_toint("-12345"));
   assert(str_toint("1234") == 1234);
   assert(str_toint("-12345") == -12345);

   dbg("str_inttostr(1234): \"%s\"", str_inttostr(1234));
   dbg("str_inttostr(-12345): \"%s\"", str_inttostr(-12345));

   printf("\nEnd of tests.\n\n");

   /*
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
