#!/bin/bash

DIR_TO_CHECK=$1
#DIR_TO_CHECK='/home/mark/Desktop/*'

echo $(sha1sum $DIR_TO_CHECK) > newSha1sum.txt

old=$(sha1sum oldSha1sum.txt)
new=$(sha1sum newSha1sum.txt)

if [ "${old:0:40}" != "${new:0:40}" ]
then
  echo "Directory has changed!"
  echo $(cat newSha1sum.txt) > oldSha1sum.txt
else
  echo "Directory hasn't changed!"
fi

rm newSha1sum.txt
