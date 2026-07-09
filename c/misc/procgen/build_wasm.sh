mkdir -p wasm
cd wasm

emcc ../procgen.cpp -o index.html -s USE_SDL=2 -s USE_SDL_MIXER=2 -s ALLOW_MEMORY_GROWTH=1 -D__EMSCRIPTEN__
emrun --port 8080 .
