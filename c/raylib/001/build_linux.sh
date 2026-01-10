mkdir -p build
cd build
gcc -o main ../src/main_linux.c ../lib/libraylib_linux.a -lGL -lm -lpthread -ldl -lrt -lX11
./main
