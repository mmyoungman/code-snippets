#!/bin/bash

check_cmd() {
  command -v "$1" >/dev/null 2>&1
}

commands=("unzip" "cmake" "unzipp" "clang++")

for cmd in "${commands[@]}"; do
  if check_cmd "$cmd"; then
    echo "$cmd is installed"
  else
    echo "$cmd is not installed"
  fi
done
