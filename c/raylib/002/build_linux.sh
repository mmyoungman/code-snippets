#!/bin/bash

# default build options
ROOT_DIR="002"
BUILD="DEBUG"

# parse script arguments
while [[ $# -gt 0 ]]; do
    key="$1"
    case "$key" in
    	-r|--release)
    		BUILD="RELEASE"
    		;;
    esac
    shift
done

# only allow to be run from project root
dir=`basename $PWD`
if [ "$dir" != "${ROOT_DIR}" ]; then
  echo "Run $(basename "$0") from project root directory '${ROOT_DIR}' - you are in $(pwd)"
  exit 1
fi

RAYLIB_VERSION=5.5
RAYLIB_DIR=raylib-${RAYLIB_VERSION}_linux_amd64
if [ ! -d contrib/${RAYLIB_DIR} ]; then
  echo "You need ${RAYLIB_DIR}.zip to be extracted to contrib directory for this to run"
  exit 1
fi

mkdir -p build

DEBUG_FLAGS="-g"
INC="-I./contrib/$RAYLIB_DIR/include"
DEPS="-L./contrib/$RAYLIB_DIR/lib/ -l:libraylib.a -lGL -lm -lpthread -ldl -lrt -lX11"
OPTS="-Wall -Wunused-variable -fdiagnostics-absolute-paths -fno-exceptions"
PREFIX=""

# use bear to generate compile_commands.json for clangd
check_cmd() {
  command -v "$1" >/dev/null 2>&1
}
if check_cmd "bear"; then
  PREFIX="bear --"
else
  echo "NOTE: bear not installed, so not using it to generate compile_commands.json"
fi

if [ "$BUILD" = "DEBUG" ]; then
  $PREFIX clang -o build/main main_linux.c -DDEBUG $DEBUG_FLAGS $INC $DEPS $OPTS
else
  $PREFIX clang -o build/main main_linux.c $INC $DEPS $OPTS
fi
