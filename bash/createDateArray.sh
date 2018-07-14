#!/bin/bash

for (( i=0; i<30; i++ ))
do
    dates[i]=$(date -d-"$i"day +%d-%m-%y)
    echo "${dates[i]}"
done

#for i in "${dates[@]}"
#do
#    echo $i
#done
#
#echo "${dates[29]}"
