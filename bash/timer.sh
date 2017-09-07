#!/bin/bash

START=`date +%s.%N`
time echo "Test"
END=`date+%s.%N`

RUNTIME=$(echo "$END - $START" | bc -l)

echo "Run time: $RUNTIME seconds"