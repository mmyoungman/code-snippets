#include "lib-mmy.h"

typedef struct {
   u64 capacity;
   u64 size;
   u64 *data;
} arr;

arr* create_arr() {
   arr *result = malloc(sizeof(arr));
   result->capacity = 128;
   result->size = 0;
   result->data = malloc(sizeof(u64) * 128);

   return result;
}

void add(arr *array, u64 element) {
   if(array->size >= array->capacity) {
      array->capacity *= 2;
      array->data = realloc(array->data, sizeof(u64) * array->capacity);
   }
   array->data[array->size] = (u64)element;
   array->size++;
}

void remove_unordered(arr *array, u64 index) {
   if(index > array->size) {
      log_err("Trying to remove at index %d, which is greater than array->size %d, so ignoring remove request", index, array->size);
      return;
   }
   else {
      if(array->size == 1) {
         array->size--;
         return;
      }
      array->data[index] = array->data[array->size];
      array->size--;
   }
}

void remove_ordered(arr *array, u64 index) {
   if(index > array->size) {
      return;
   }
   else {
      if(array->size == 1) {
         array->size--;
         return;
      }
      u64 pos = 0;
      array->size--;
      while(pos < index) { pos++; }
      while(pos < array->size) { 
         array->data[pos] = array->data[pos+1]; 
         pos++;
      }
   }
}

int main(int argc, char** argv) {
   arr *array = create_arr();

   for(int i = 0; i < 1000; i++) {
      //dbg("%d", i);
      add(array, i);
   }

   for(int i = 0; i < 995; i++) {
      dbg("Removing %d...", array->data[0]);
      remove_ordered(array, 0);
   }

   if(array->size == 0) {
      printf("Array is empty\n");
   }
   for(int i = 0; i < array->size; i++) {
      //dbg("%ld", array->data[i]);
      if(i == array->size-1) {
         printf("%ld\n", array->data[i]);
         break;
      }
      else {
         printf("%ld, ", array->data[i]);
      }
   }
   return 0;
}
