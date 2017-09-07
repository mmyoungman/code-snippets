#!/bin/bash

var1=0
var2=0
fileToRun=defaultfile.txt

while [[ $# -gt 0 ]]; do
  key="$1"
  case "$key" in
    -v|--var1)
    var1=1
    ;;
    --var2)
    var2=1
    ;;
    -f|--file)
    shift
    fileToRun="$1"
    ;;
    *)
    echo "Unknown option '$key'"
    ;;
  esac
  shift
done

if [ "$var1" = 0 ]; then

someCommand --var2:$var2 $fileToRun

else

echo "Do something else!"

fi