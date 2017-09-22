// g++ -rdynamic -Wl,-soname=hello.so -Wall -o test dynamicLibraryLoad.cpp -ldl

#include <stdio.h>
#include <dlfcn.h>

void (*hello)();

int main() {
    void *handle;

    handle = dlopen("./hello.so", RTLD_LAZY);
    if(handle == NULL) { printf("dlopen returned NULL!\n"); }

    dlerror();
    
    hello = (void (*)()) dlsym(handle, "hello");
    if(hello == NULL) { printf("dlsym returned NULL!\n"); }

    hello();

    dlclose(handle);
}
