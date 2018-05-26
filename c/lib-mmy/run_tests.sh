#!/bin/bash
flags="-ggdb -msse2 -Wno-write-strings"

g++ $flags -D DEBUG tests.cpp -o tests
./tests

#rm tests
