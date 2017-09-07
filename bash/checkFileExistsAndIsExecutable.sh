#!/bin/bash

file=$1

if [ -e $file ]
then
  echo -e "File $file exists!"
else
  echo -e "File $file doesn't exist!"
fi

if [ -x $file ]
then
  echo -e "File $file exists and is executable!"
else
  echo -e "File $file either doesn't exist or isn't executable!"
fi