// compiled with g++ -shared -o hello.so -fPIC dynamicLibraryHelloWorld.cpp
#include <stdio.h>

extern "C" {

void hello() {
    printf("Hello world!\n");
}

}
