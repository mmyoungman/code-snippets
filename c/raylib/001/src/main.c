void UpdateDrawFrame(void) {
	BeginDrawing();

	ClearBackground(RAYWHITE);
	DrawText("Some text inside the window!", 190, 200, 20, LIGHTGRAY);

	EndDrawing();
}
