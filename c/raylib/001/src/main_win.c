#include "../include/raylib.h"

#include "main.c"

int main(void) {
	InitWindow(800, 600, "Raylib test window");

	while(!WindowShouldClose()) {
		main_render();
	}

	CloseWindow();

	return 0;
}
