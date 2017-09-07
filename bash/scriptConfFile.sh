#!/bin/bash

if [ -n "$1" ]
then
  echo "Reading config..."
  source $1
  echo "Username: $username"
  echo "Password: $password"

else
  echo "Provide config file as an argument"
fi