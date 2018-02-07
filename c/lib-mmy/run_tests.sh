#!/bin/bash
flags="-ggdb -msse2 -Wno-write-strings"

g++ $flags tests.cpp -o tests
g++ $flags -D DEBUG tests.cpp -o tests-debug
./tests
./tests-debug
rm tests tests-debug
