#!/bin/bash
flags="-ggdb -msse2 -Wno-write-strings"

g++ $flags test.cpp -o test
