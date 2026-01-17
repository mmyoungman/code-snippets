#include "raylib.h"
#include "utils.h"
#include <stdio.h>

#define MIN(a, b) ((a)<(b) ? (a) : (b))

int main(void) {
  const u32 screenWidth = 1280;
  const u32 screenHeight = 720;
  SetConfigFlags(FLAG_WINDOW_RESIZABLE | FLAG_VSYNC_HINT);
  InitWindow(screenWidth, screenHeight, "Raylib test 002");

  u32 currentMonitor = GetCurrentMonitor();
  printf("Current monitor: %d\n", currentMonitor);

  u32 gameScreenWidth = 640;
  u32 gameScreenHeight = 480;

  log_info("Screen width %d\n", screenWidth);
  log_info("Screen height %d\n", screenHeight);

  log_info("Game screen width %d\n", gameScreenWidth);
  log_info("Game screen height %d\n", gameScreenHeight);

  #if DEBUG
  Vector2 monitor1Position = GetMonitorPosition(1);
  SetWindowPosition(monitor1Position.x + 100, monitor1Position.y + 100);
  #endif

  SetTargetFPS(60);

  RenderTexture2D target = LoadRenderTexture(gameScreenWidth, gameScreenHeight);
  SetTextureFilter(target.texture, TEXTURE_FILTER_BILINEAR);

  while (!WindowShouldClose()) {
    // Need to virtualise mouse/touch x/y too: https://github.com/raysan5/raylib/blob/master/examples/core/core_window_letterbox.c

    BeginTextureMode(target);
    {
      ClearBackground(WHITE);
      // Draw game here
      DrawRectangle(0, 0, gameScreenWidth, gameScreenHeight, GREEN);
    }
    EndTextureMode();

    BeginDrawing();
    {
      ClearBackground(BLACK);
      f64 scale = MIN((f64)GetScreenWidth()/gameScreenWidth, (f64)GetScreenHeight()/gameScreenHeight);
      DrawTexturePro(target.texture,
                     (Rectangle){0.0f, 0.0f, (f64)target.texture.width, (f64)-target.texture.height },
                     (Rectangle){(GetScreenWidth() - ((f64)gameScreenWidth*scale))*0.5f, (GetScreenHeight() - ((f64)gameScreenHeight*scale))*0.5f, (f64)gameScreenWidth*scale, (f64)gameScreenHeight*scale },
                     (Vector2) { 0, 0 },
                     0.0f,
                     (Color){255, 255, 255, 255});
    }
    EndDrawing();
  }

  CloseWindow();

  return 0;
}
