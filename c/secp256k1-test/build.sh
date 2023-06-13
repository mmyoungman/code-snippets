#!/bin/bash

# default build script options
ROOT_DIR="secp256k1-test"
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
	echo "Run this script from project root - should be called '${ROOT_DIR}'"
	exit 1
fi

mkdir -p build/lib
mkdir -p build/contrib
mkdir -p build/include

# secp256k1
SECP256K1_VERSION="0.3.2"
if [ ! -f build/lib/libsecp256k1.a ]; then
  if [ ! -f zips/secp256k1-${SECP256K1_VERSION}.zip ]; then
    echo "secp256k1-${SECP256K1_VERSION}.zip not found in zips directory"
    exit 1
  else
    unzip -d build/contrib zips/secp256k1-${SECP256K1_VERSION}.zip

    cmake -S build/contrib/secp256k1-${SECP256K1_VERSION} -B build/contrib/secp256k1-${SECP256K1_VERSION} -DSECP256K1_DISABLE_SHARED=ON
    make -C build/contrib/secp256k1-${SECP256K1_VERSION}

    cp -r build/contrib/secp256k1-${SECP256K1_VERSION}/include/. build/include
    cp build/contrib/secp256k1-${SECP256K1_VERSION}/src/libsecp256k1.a build/lib
  fi
fi

DEBUG_FLAGS="-g"
OPTS="-Wall -Wunused-variable -fdiagnostics-absolute-paths -fno-exceptions"
DEPS="-lsecp256k1"
INC="-Lbuild/lib -Ibuild/include -Wl,-Rbuild/lib"

if [ "$BUILD" = "DEBUG" ]; then
	clang++ -o build/secp256k1-test linux_secp256k1-test.cpp -DDEBUG $DEBUG_FLAGS $INC $DEPS $OPTS
else
	clang++ -o build/secp256k1-test linux_secp256k1-test.cpp $INC $DEPS $OPTS
fi
