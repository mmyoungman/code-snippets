/*
  To run, you need:
    - gcc/g++
    - SDL2 (Ubuntu 16.04, "sudo apt install libsdl2-2.0-0 libsdl2-dev")

  To build: g++ -ggdb procgen.cpp -lSDL2 -o procgen
*/

#include <SDL2/SDL.h>
#include <SDL2/SDL_events.h>

#include <stdio.h>
#include <stdlib.h> // malloc, free

#include "procgen-lib.h"

#ifdef __EMSCRIPTEN__
#include <emscripten.h>
#endif

#define SCREEN_WIDTH 960
#define SCREEN_HEIGHT 540

const int GRID_BORDER = 10;
const int GRID_WIDTH = 94;
const int GRID_HEIGHT = 52;
const int CELL_WIDTH = (SCREEN_WIDTH-(GRID_BORDER*2))/GRID_WIDTH;
const int CELL_HEIGHT = (SCREEN_HEIGHT-(GRID_BORDER*2))/GRID_HEIGHT;

struct framebuffer
{
    uint8_t *data;
    int w, h;
    int bytesperpixel;
};

struct gridstatus
{
    int cells[GRID_WIDTH][GRID_HEIGHT];
    int w, h;
};

SDL_Window *window;
SDL_Renderer *renderer;
framebuffer buffer;
SDL_Texture *texture;

// Init gridstatus
gridstatus grid;
int currentX;
int currentY;
int hubBlobSize;
int status;

// Pre loop stuff
int running;
int fullscreen;
int newTime;
int prevTime;
int deltaTime;
SDL_Event e;

enum cellstatus
{
    CELL_OFF,
    CELL_ON,
    CELL_WALL,
    CELL_SHORTCUT,
    CELL_EXIT,
    CELL_HUB_START,
    CELL_HUB,
    CELL_AREA1,
    CELL_AREA2,
    CELL_AREA3,
    CELL_AREA4,
    CELL_AREA5,
    CELL_END
};

void clear(gridstatus *gs)
{
    for(int i = 0; i < GRID_WIDTH; i++)
    {
        for(int j = 0; j < GRID_HEIGHT; j++)
        {
            gs->cells[i][j] = CELL_OFF;
        }
    }
}

void fillGrid(framebuffer fb, gridstatus gs)
{
    int pitch = fb.w * fb.bytesperpixel;
    uint8_t *row = fb.data;

    for(int j = 0; j < fb.h; j++)
    {
        uint32_t *pixel = (uint32_t *)row;
        for(int i = 0; i < fb.w; i++)
        {
            uint8_t blue = 0;
            uint8_t green = 0;
            uint8_t red = 0;
            uint8_t alpha = 0;

            if(i >= GRID_BORDER && j >= GRID_BORDER)
            {
                int cx = (i - GRID_BORDER) / CELL_WIDTH;
                int cy = (j - GRID_BORDER) / CELL_HEIGHT;
                if(cx < GRID_WIDTH && cy < GRID_HEIGHT)
                {
                    switch(gs.cells[cx][cy])
                    {
                        case CELL_ON:
                            blue = 255; green = 255; red = 255; alpha = 255;
                            break;
                        case CELL_WALL:
                            blue = 0; green = 0; red = 255;
                            break;
                        case CELL_SHORTCUT:
                            blue = 255; green = 0; red = 255;
                            break;
                        case CELL_HUB_START:
                            blue = 0; green = 255; red = 0;
                            break;
                        case CELL_HUB:
                            blue = 203; green = 192; red = 255;
                            break;
                        case CELL_EXIT:
                            blue = 128; green = 0; red = 128;
                            break;
                        case CELL_AREA1:
                            blue = 204; green = 204; red = 0;
                            break;
                        case CELL_AREA2:
                            blue = 100; green = 100; red = 200;
                            break;
                        case CELL_AREA3:
                            blue = 0; green = 255; red = 255;
                            break;
                        case CELL_AREA4:
                            blue = 50; green = 255; red = 155;
                            break;
                        case CELL_AREA5:
                            blue = 100; green = 155; red = 55;
                            break;
                        case CELL_END:
                            blue = 0; green = 128; red = 255;
                            break;
                        default:
                            break;
                    }
                }
            }

            *pixel++ = ((alpha << 24) | (red << 16) | (green << 8) | blue);
        }
        row += pitch;
    }
}

void drawGrid(framebuffer fb)
{
    int pitch = fb.w*fb.bytesperpixel;
    uint8_t *row = fb.data;

    for(int y = 0; y < fb.h; y++)
    {
        uint32_t *pixel = (uint32_t *)row;
        for(int x = 0; x < fb.w; x++)
        {
            if(((x+GRID_BORDER)%CELL_WIDTH == 0 && x >= GRID_BORDER && x <= fb.w-GRID_BORDER && y >= GRID_BORDER && y <= fb.h-GRID_BORDER) ||
               ((y+GRID_BORDER)%CELL_HEIGHT == 0 && y >= GRID_BORDER && y <= fb.h-GRID_BORDER && x >= GRID_BORDER && x <= fb.w-GRID_BORDER) ||
               (x == fb.w-GRID_BORDER && y >= GRID_BORDER && y <= fb.h-GRID_BORDER) ||
               (y == fb.h-GRID_BORDER && x >= GRID_BORDER && x <= fb.w-GRID_BORDER))
            {
                uint8_t blue = 255;
                uint8_t green = 255;
                uint8_t red = 255;
                uint8_t alpha = 255;

                *pixel = ((alpha << 24) | (red << 16) | (green << 8) | blue);
            }
            pixel++;
        }
        row += pitch;
    }
}

// Checks whether cells in area are CELL_OFF
// Seems to be working...
bool checkAreaIsOff(gridstatus *gs, int x, int y, int w, int h)
{
    if(w < 0)
    {
        x += w+1;
        w = -w;
    }
    if(h < 0)
    {
        y += h+1;
        h = -h;
    }

    // Look for something that isn't CELL_OFF in area
    for(int i = x; i < x+w; i++)
        for(int j = y; j < y+h; j++)
        {
            if(i < 0 || i > GRID_WIDTH-1 || j < 0 || j > GRID_HEIGHT-1 || gs->cells[i][j] != CELL_OFF)
                return false;
        }
    return true;
}

void createExitCells(gridstatus *gs, cellstatus CS)
{
    for(int x = 0; x < GRID_WIDTH; x++)
    {
        for(int y = 0; y < GRID_HEIGHT; y++)
        {
            // Exit point must be able to accomodate corridor
            int corridorLen = 5;
            int corridorWid = 5; // Should be odd
            if(gs->cells[x][y] == CS)
            {
                if(y-1 >= 0 && gs->cells[x][y-1] == CELL_OFF)
                {
                    if(checkAreaIsOff(gs, x-(corridorWid/2), y-corridorLen, corridorWid, corridorLen))
                        gs->cells[x][y-1] = CELL_EXIT;
                }
                if(x-1 >= 0 && gs->cells[x-1][y] == CELL_OFF)
                {
                    if(checkAreaIsOff(gs, x-corridorLen, y-(corridorWid/2), corridorLen, corridorWid))
                        gs->cells[x-1][y] = CELL_EXIT;
                }
                if(x+1 < GRID_WIDTH && gs->cells[x+1][y] == CELL_OFF)
                {
                    if(checkAreaIsOff(gs, x+1, y-(corridorWid/2), corridorLen, corridorWid))
                        gs->cells[x+1][y] = CELL_EXIT;
                }
                if(y+1 < GRID_HEIGHT && gs->cells[x][y+1] == CELL_OFF)
                {
                    if(checkAreaIsOff(gs, x-(corridorWid/2), y+1, corridorWid, corridorLen))
                        gs->cells[x][y+1] = CELL_EXIT;
                }
            }
        }
    }
}

// Creates a random blob of the size
void createBlob(gridstatus *gs, int *inputX, int *inputY, int size, cellstatus CS)
{
    for(int i = 0; i < size; i++)
    {

        // If can eat start cell, eat it
        if(gs->cells[*inputX][*inputY] == CELL_OFF)
        {
            gs->cells[*inputX][*inputY] = CS;
        }
        else
        {
            int x = *inputX;
            int y = *inputY;

            // Check what directions we can eat
            char array[4];
            for(int n = 0; n < 4; n++)
                array[n] = '0';
            int arraylen = 0;

            while(y-1 >= 0 && gs->cells[x][y-1] == CELL_OFF)
            {
                array[arraylen] = 'N';
                arraylen++;
                break;
            }
            while(x-1 >= 0 && gs->cells[x-1][y] == CELL_OFF)
            {
                array[arraylen] = 'W';
                arraylen++;
                break;
            }
            while(x+1 < GRID_WIDTH && gs->cells[x+1][y] == CELL_OFF)
            {
                array[arraylen] = 'E';
                arraylen++;
                break;
            }
            while(y+1 < GRID_HEIGHT && gs->cells[x][y+1] == CELL_OFF)
            {
                array[arraylen] = 'S';
                arraylen++;
                break;
            }

            // If no cell to eat, move in a random direction, only over CS cells
            // Not ideal...
            if(arraylen == 0)
            {
                int direction = stb_rand()%4;
                if(direction == 0 && y-1 >= 0 && gs->cells[x][y-1] == CS)
                    *inputY = y-1;
                if(direction == 1 && y+1 < GRID_HEIGHT && gs->cells[x][y+1] == CS)
                    *inputY = y+1;
                if(direction == 2 && x-1 >= 0 && gs->cells[x-1][y] == CS)
                    *inputX = x-1;
                if(direction == 3 && x+1 < GRID_WIDTH && gs->cells[x+1][y] == CS)
                    *inputX = x+1;
            }
            // Else, choose a random cell to eat and size--
            else
            {
                // Eat current position cell if its CELL_OFF
                int direction = -1;
                if(gs->cells[x][y] == CELL_OFF)
                    gs->cells[x][y] = CS;
                else
                {
                    direction = stb_rand()%arraylen;
                    if(array[direction] == 'N')
                        gs->cells[x][y-1] = CS;
                    if(array[direction] == 'W')
                        gs->cells[x-1][y] = CS;
                    if(array[direction] == 'E')
                        gs->cells[x+1][y] = CS;
                    if(array[direction] == 'S')
                        gs->cells[x][y+1] = CS;
                }

                // Random chance to move
                // If there is only 1 remaining, move
                if(direction != -1 && stb_rand()%arraylen == 0)
                {
                    if(array[direction] == 'N')
                        *inputY = y-1;
                    if(array[direction] == 'W')
                        *inputX = x-1;
                    if(array[direction] == 'E')
                        *inputX = x+1;
                    if(array[direction] == 'S')
                        *inputY = y+1;
                }
            }
        }
    }
}

void createWalls(gridstatus *gs, cellstatus cs)
{
     for(int i = 0; i < GRID_WIDTH; i++)
     {
         for(int j = 0; j < GRID_HEIGHT; j++)
         {
             if(gs->cells[i][j] == cs)
             {

                 if(j-1 >= 0 && gs->cells[i][j-1] == CELL_OFF)
                     gs->cells[i][j-1] = CELL_WALL;
                 if(i-1 >= 0 && gs->cells[i-1][j] == CELL_OFF)
                    gs->cells[i-1][j] = CELL_WALL;
                 if(i+1 < GRID_WIDTH && gs->cells[i+1][j] == CELL_OFF)
                     gs->cells[i+1][j] = CELL_WALL;
                 if(j+1 < GRID_HEIGHT && gs->cells[i][j+1] == CELL_OFF)
                     gs->cells[i][j+1] = CELL_WALL;
             }
         }
     }
}

struct point
{
    int x;
    int y;
};

struct pointList
{
    point *list;
    int count;
    int mallocSize;
};

point makePoint(int x, int y)
{
     point result;
     result.x = x;
     result.y = y;

     return result;
}

void addPoint(pointList *pl, point p)
{
    pl->count++;

    // If more memory is needed, realloc more
    if(pl->mallocSize/pl->count < sizeof(point))
    {
        pl->mallocSize += 4096;
        pl->list = (point *)realloc(pl->list, pl->mallocSize);
        if(pl->list == NULL)
            exit(1);
    }

    pl->list[pl->count-1] = p;
}

void findExitCells(gridstatus *gs, pointList *pl)
{
    pl->count = 0;
    for(int i = 0; i < GRID_WIDTH; i++)
    {
        for(int j = 0; j < GRID_HEIGHT; j++)
        {
            if(gs->cells[i][j] == CELL_EXIT)
                addPoint(pl, makePoint(i, j));
        }
    }
}

void createArea(gridstatus *gs, pointList *pl, cellstatus CS, bool clearWalls)
{
    if(pl->count == 0)
        return;

    point p = pl->list[stb_rand()%pl->count];
    gs->cells[p.x][p.y] = CS;

    createBlob(gs, &p.x, &p.y, hubBlobSize*2, CS);
    createExitCells(gs, CS);
    createWalls(gs, CS);

    for(int i = 0; i < GRID_WIDTH; i++)
    {
        for(int j = 0; j < GRID_HEIGHT; j++)
        {
            if(gs->cells[i][j] == CELL_EXIT)
            {
                if(i-1 >= 0 && i+1 < GRID_WIDTH && j-1 >= 0 && j+1 < GRID_HEIGHT &&
                   gs->cells[i-1][j] != CELL_OFF && gs->cells[i+1][j] != CELL_OFF &&
                   gs->cells[i][j-1] != CELL_OFF && gs->cells[i][j+1] != CELL_OFF)
                {
                    gs->cells[i][j] = CELL_SHORTCUT;
                }
                else if(clearWalls)
                    gs->cells[i][j] = CELL_WALL;
            }
        }
    }

    if(clearWalls)
    {
        for(int i = 0; i < GRID_WIDTH; i++)
            for(int j = 0; j < GRID_HEIGHT; j++)
                if(gs->cells[i][j] == CELL_WALL)
                    gs->cells[i][j] = CELL_OFF;
    }
}

int init(void)
{
    // Prefer EGL over GLX. Some X servers have a broken GLX extension
    // that crashes SDL at context creation.
    SDL_SetHint(SDL_HINT_VIDEO_X11_FORCE_EGL, "1");

    if(SDL_Init(SDL_INIT_VIDEO) != 0)
    {
        printf("Video init failed. SDL_Error: %s\n", SDL_GetError());
        return 1;
    }

    // Create Window
    window = SDL_CreateWindow("Proc Gen",
                                          SDL_WINDOWPOS_UNDEFINED,
                                          SDL_WINDOWPOS_UNDEFINED,
                                          SCREEN_WIDTH,
                                          SCREEN_HEIGHT,
                                          SDL_WINDOW_RESIZABLE);
    if(window == NULL)
    {
        printf("Window creation failed. SDL_Error: %s\n", SDL_GetError());
        return 1;
    }

    renderer = SDL_CreateRenderer(window, -1, 0);
    if(renderer == NULL)
    {
        printf("Renderer creation failed. SDL_Error: %s\n", SDL_GetError());
        return 1;
    }

    SDL_GetWindowSize(window, &buffer.w, &buffer.h);
    buffer.bytesperpixel = 4;
    buffer.data = (uint8_t *)malloc(buffer.w * buffer.h *
                                    buffer.bytesperpixel);
    if(buffer.data == NULL)
    {
        printf("Failed to allocate framebuffer.\n");
        return 1;
    }

    texture = SDL_CreateTexture(renderer,
                                            SDL_PIXELFORMAT_ARGB8888,
                                            SDL_TEXTUREACCESS_STREAMING,
                                            buffer.w, buffer.h);
    if(texture == NULL)
    {
        printf("Texture creation failed. SDL_Error: %s\n", SDL_GetError());
        return 1;
    }
    // Init gridstatus
    grid.w = GRID_WIDTH;
    grid.h = GRID_HEIGHT;

    clear(&grid);

    currentX = GRID_WIDTH/2;
    currentY = GRID_HEIGHT/2;
    grid.cells[currentX][currentY] = CELL_HUB_START;
    hubBlobSize = 100;
    status = 0;

    // Pre loop stuff
    running = 1;
    fullscreen = 0;
    newTime = 0;

    return 0;
}

void mainloop(void)
{
    if(!running) {
        #ifdef __EMSCRIPTEN__
        emscripten_cancel_main_loop();
        #else
        exit(0);
        #endif
    }

    while(SDL_PollEvent(&e) != 0)
    {
        switch(e.type)
        {
            case SDL_KEYDOWN:
            //case SDL_KEYUP:
                switch(e.key.keysym.sym)
                {
                    case SDLK_q:
                        running = 0;
                        break;
                    case SDLK_ESCAPE:
                        running = 0;
                        break;
                    case SDLK_f:
                        if(!fullscreen)
                            SDL_SetWindowFullscreen(window, SDL_WINDOW_FULLSCREEN);
                        else
                            SDL_SetWindowFullscreen(window, 0);

                        fullscreen = !fullscreen;
                        break;
                    case SDLK_c:
                        currentX = GRID_WIDTH/2;
                        currentY = GRID_HEIGHT/2;
                        status = 0;
                        clear(&grid);
                        grid.cells[GRID_WIDTH/2][GRID_HEIGHT/2] = CELL_HUB_START;
                        break;
                } break;
            case SDL_MOUSEMOTION:
                int x, y;
                SDL_GetMouseState(&x, &y);
                //printf("Mouse x:%d y:%d\n", x, y);
                break;
            case SDL_MOUSEBUTTONDOWN:
                // Doesn't work?
                //if(e.button == SDL_BUTTON_LEFT)
                //    printf("Left mouse button!\n");
                break;
            case SDL_WINDOWEVENT:
                switch(e.window.event)
                {
                    case SDL_WINDOWEVENT_RESIZED:
                    case SDL_WINDOWEVENT_SIZE_CHANGED:
                        //printf("Window Width:%d Height:%d\n", e.window.data1, e.window.data2);
                        free(buffer.data);
                        SDL_DestroyTexture(texture);
                        texture = SDL_CreateTexture(renderer, SDL_PIXELFORMAT_ARGB8888,
                                                    SDL_TEXTUREACCESS_STREAMING,
                                                    e.window.data1, e.window.data2);
                        buffer.data = (uint8_t *)malloc(e.window.data1 * e.window.data2 * buffer.bytesperpixel);
                        buffer.w = e.window.data1;
                        buffer.h = e.window.data2;
                        break;
                } break;
            case SDL_QUIT:
                running = 0;
                break;
        }
    }

    // Do procedural generation
    if(status == 0)
    {
        createBlob(&grid, &currentX, &currentY, hubBlobSize, CELL_HUB);
        createExitCells(&grid, CELL_HUB);
        createWalls(&grid, CELL_HUB);
        createWalls(&grid, CELL_HUB_START);

        pointList pl;
        pl.count = 0;
        pl.mallocSize = 4096;
        pl.list = (point *)malloc(pl.mallocSize);

        cellstatus areas[] = { CELL_AREA1, CELL_AREA2, CELL_AREA3, CELL_AREA4 };
        for(int a = 0; a < 4; a++)
        {
            findExitCells(&grid, &pl);
            createArea(&grid, &pl, areas[a], true);

            createExitCells(&grid, CELL_HUB);
            for(int b = 0; b <= a; b++)
                createExitCells(&grid, areas[b]);
            createWalls(&grid, CELL_HUB);
            createWalls(&grid, CELL_HUB_START);
            for(int b = 0; b <= a; b++)
                createWalls(&grid, areas[b]);
        }

        findExitCells(&grid, &pl);
        createArea(&grid, &pl, CELL_AREA5, false);

        free(pl.list);

        status--;
    }

    // Write to buffer
    fillGrid(buffer, grid);
    drawGrid(buffer);

    // Update the screen
    SDL_UpdateTexture(texture, 0, buffer.data,
                      buffer.w * buffer.bytesperpixel);
    SDL_RenderCopy(renderer, texture, 0, 0);
    SDL_RenderPresent(renderer);

    // Cap the framerate at ~60fps
    newTime = (int)SDL_GetTicks();
    deltaTime = newTime - prevTime;
    if(deltaTime < 16)
        SDL_Delay(16 - deltaTime);
    prevTime = (int)SDL_GetTicks();
}

int main(int argc, char* args[])
{
    if(init() != 0) {
        return 1;
    }

    #ifdef __EMSCRIPTEN__
    emscripten_set_main_loop(mainloop, 0, 1);
    #else
    while(running)
    {
        mainloop();
    }
    #endif

    SDL_SetWindowFullscreen(window, 0);

    free(buffer.data);

    // Necessary?
    SDL_DestroyRenderer(renderer);
    SDL_DestroyWindow(window);

    SDL_Quit();

    return 0;
}
