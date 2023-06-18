#!/bin/bash

# default build script options
ROOT_DIR="secure-websockets-client"
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

#ZLIB_VERSION="1.2.13"
#if [ ! -f build/contrib/zlib-${ZLIB_VERSION}/libz.a ]; then
#  if [ ! -f zips/zlib-${ZLIB_VERSION}.tar.gz ]; then
#    echo "zlib-${ZLIB_VERSION}.tar.gz not found in zips directory"
#    exit 1
#  else
#    tar -xf zips/zlib-${ZLIB_VERSION}.tar.gz -C build/contrib
#    cmake -S build/contrib/zlib-${ZLIB_VERSION} -B build/contrib/zlib-${ZLIB_VERSION}
#    make -C build/contrib/zlib-${ZLIB_VERSION}
#    cp build/contrib/zlib-${ZLIB_VERSION}/libz.a build/lib/
#    cp build/contrib/zlib-${ZLIB_VERSION}/*.h build/include
#  fi
#fi

MBEDTLS_VERSION="3.4.0"
if [ ! -f build/contrib/mbedtls-${MBEDTLS_VERSION}/library/libmbedtls.a ]; then
  if [ ! -f zips/mbedtls-${MBEDTLS_VERSION}.zip ]; then
    echo "mbedtls-${MBEDTLS_VERSION}.zip not found in zips directory"
    exit 1
  else
    unzip -d build/contrib/ zips/mbedtls-${MBEDTLS_VERSION}.zip
    cmake -S build/contrib/mbedtls-${MBEDTLS_VERSION} -B build/contrib/mbedtls-${MBEDTLS_VERSION}
    make -C build/contrib/mbedtls-${MBEDTLS_VERSION}
    cp build/contrib/mbedtls-${MBEDTLS_VERSION}/library/*.a build/lib/
    cp -r build/contrib/mbedtls-${MBEDTLS_VERSION}/include/mbedtls build/include
    cp -r build/contrib/mbedtls-${MBEDTLS_VERSION}/include/psa build/include
  fi
fi

DEBUG_FLAGS="-g"
OPTS="-Wall -Wunused-variable -fdiagnostics-absolute-paths -fno-exceptions"
DEPS="-lmbedtls -lmbedx509 -lmbedcrypto"
INC="-Lbuild/lib -Ibuild/include"

if [ "$BUILD" = "DEBUG" ]; then
	clang++ -o build/secure-websockets-client linux_secure-websockets-client.cpp -DDEBUG $DEBUG_FLAGS $INC $DEPS $OPTS
else
	clang++ -o build/secure-websockets-client linux_secure-websockets-client.cpp $INC $DEPS $OPTS
fi
