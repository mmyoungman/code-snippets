#!/bin/bash

g++ -c dynamicLibraryHello.cpp
g++ -shared -fPIC -o hello.so dynamicLibraryHello.o
# or g++ -shared -fPIC -o hello.so dynamicLibraryHello.cpp
g++ -o test dynamicLibraryLoad.cpp -ldl
