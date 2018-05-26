#include "lib-mmy.h"

int main()
{

dbg("Starting tests...\n\n");

   // Test 000.
   assert(1024 == kilobytes(1));
   assert((1024*1024) == megabytes(1));
   assert(((u64)10*1024*1024*1024) == gigabytes(10));

   // log_err("Log error: %d, %s", 42, "string literal");
   // log_warn("Log warning: %d, %s", 12, "another string literal");
   // log_info("Log info: %d, %s", 13, "yet another string literal");
   // dbg("Test msg: %d, %s", 14, "not another bloody string literal");


   // Test 001.


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
   free(copy);

   assert(str_beginswith(str, "a stri"));
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
   assert(str_isalpha(str));
   str[10] = '#';
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
   
   assert(str_toint("1234") == 1234);
   assert(str_toint("-12345") == -12345);

   // shouldAssert(str_toint("12-12"));

   dbg("End of tests\n\n");
}