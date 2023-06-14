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

# zlib (for IXWebSocket)
ZLIB_VERSION="1.2.13"
if [ ! -f build/contrib/zlib-${ZLIB_VERSION}/libz.a ]; then
  if [ ! -f zips/zlib-${ZLIB_VERSION}.tar.gz ]; then
    echo "zlib-${ZLIB_VERSION}.tar.gz not found in zips directory"
    exit 1
  else
    tar -xf zips/zlib-${ZLIB_VERSION}.tar.gz -C build/contrib
    cmake -S build/contrib/zlib-${ZLIB_VERSION} -B build/contrib/zlib-${ZLIB_VERSION}
    make -C build/contrib/zlib-${ZLIB_VERSION}
    cp build/contrib/zlib-${ZLIB_VERSION}/libz.a build/lib/
    cp build/contrib/zlib-${ZLIB_VERSION}/*.h build/include
  fi
fi

# mbedtls (for IXWebSocket)
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

# IXWebSocket
IXWEBSOCKET_VERSION="11.4.4"
if [ ! -f build/lib/libixwebsocket.a ]; then
  if [ ! -f zips/IXWebSocket-${IXWEBSOCKET_VERSION}.zip ]; then
    echo "IXWebSocket-${IXWEBSOCKET_VERSION}.zip not found in zips directory"
    exit 1
  else
    unzip -d build/contrib zips/IXWebSocket-${IXWEBSOCKET_VERSION}.zip

    cmake -DUSE_TLS=1 -DUSE_MBED_TLS=1 -DMBEDTLS_VERSION_GREATER_THAN_3="build/include" -DMBEDTLS_INCLUDE_DIRS="build/include" -DMBEDTLS_LIBRARY="build/lib/libmbedtls.a" -DMBEDCRYPTO_LIBRARY="build/lib/libmbedcrypto.a" -DMBEDX509_LIBRARY="build/lib/libmbedx509.a" -DZLIB_INCLUDE_DIR="build/include" -DZLIB_LIBRARY_RELEASE="build/lib/zlib.a" -S build/contrib/IXWebSocket-${IXWEBSOCKET_VERSION} -B build/contrib/IXWebSocket-${IXWEBSOCKET_VERSION}
    make -C build/contrib/IXWebSocket-${IXWEBSOCKET_VERSION}

    cp build/contrib/IXWebSocket-${IXWEBSOCKET_VERSION}/libixwebsocket.a build/lib/
    mkdir -p build/include/ixwebsocket
    cp build/contrib/IXWebSocket-${IXWEBSOCKET_VERSION}/ixwebsocket/*.h build/include/ixwebsocket
  fi
fi

DEBUG_FLAGS="-g"
OPTS="-Wall -Wunused-variable -fdiagnostics-absolute-paths -fno-exceptions"
DEPS="-lixwebsocket -lmbedtls -lmbedx509 -lmbedcrypto -lz"
INC="-Lbuild/lib -Ibuild/include"

if [ "$BUILD" = "DEBUG" ]; then
	clang++ -o build/secure-websockets-client linux_secure-websockets-client.cpp -DDEBUG $DEBUG_FLAGS $INC $DEPS $OPTS
else
	clang++ -o build/secure-websockets-client linux_secure-websockets-client.cpp $INC $DEPS $OPTS
fi
