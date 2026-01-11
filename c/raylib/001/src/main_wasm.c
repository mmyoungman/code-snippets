#include "../include/raylib.h"
#include <emscripten.h>

#include "main.c"

int main(void) {
	InitWindow(800, 600, "Raylib test window");

	emscripten_set_main_loop(main_render, 0, 1);

	CloseWindow();

	return 0;
}

