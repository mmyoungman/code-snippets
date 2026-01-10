mkdir -p build
cd build
emcc -o main.html ../src/main_wasm.c -Os -Wall ../lib/libraylib_wasm.a -s USE_GLFW=3 --shell-file ../src/minshell.html
emrun main.html
