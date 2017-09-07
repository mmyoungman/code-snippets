/*
  To run, you need:
    - gcc/g++
    - SDL2 (Ubuntu 16.04, "sudo apt install libsdl2-2.0-0 libsdl2-dev")
    - SDL2-mixer (Ubuntu 16.04, "sudo apt install libsdl2-mixer-2.0-0 libsdl2-mixer-dev")

  To build: g++ -ggdb procgen.cpp -lSDL2 -lSDL2_mixer -o procgen
*/

#include <SDL2/SDL.h>
#include <SDL2/SDL_events.h>
#include <SDL2/SDL_mixer.h>

#include <stdio.h>
#include <stdlib.h> // malloc, free

#include "procgen-lib.h"

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

void fillGrid(framebuffer buffer, gridstatus gs)
{
    int pitch = buffer.w * buffer.bytesperpixel;
    uint8_t *row = buffer.data;

    for(int j = 0; j < buffer.h; j++) 
    {
        uint32_t *pixel = (uint32_t *)row;
        for(int i = 0; i < buffer.w; i++) 
        {
            uint8_t blue = 0;
            uint8_t green = 0;
            uint8_t red = 0;
            uint8_t alpha = 0;

            if(i >= GRID_BORDER && i < buffer.w-GRID_BORDER && j >= GRID_BORDER && j < buffer.h-GRID_BORDER)
            {
                if(gs.cells[ ((i-GRID_BORDER)/CELL_WIDTH) ][ ((j-GRID_BORDER)/CELL_HEIGHT) ] == CELL_ON)
                {
                    blue = 255;
                    green = 255;
                    red = 255;
                    alpha = 255;
                }
                
                if(gs.cells[ ((i-GRID_BORDER)/CELL_WIDTH) ][ ((j-GRID_BORDER)/CELL_HEIGHT) ] == CELL_WALL)
                {
                    blue = 0;
                    green = 0;
                    red = 255;
                    alpha = 0;
                }

                if(gs.cells[ ((i-GRID_BORDER)/CELL_WIDTH) ][ ((j-GRID_BORDER)/CELL_HEIGHT) ] == CELL_SHORTCUT)
                {
                    blue = 255;
                    green = 0;
                    red = 255;
                    alpha = 0;
                }

                if(gs.cells[ ((i-GRID_BORDER)/CELL_WIDTH) ][ ((j-GRID_BORDER)/CELL_HEIGHT) ] == CELL_HUB_START)
                {
                    blue = 0;
                    green = 255;
                    red = 0;
                    alpha = 0;
                }

                if(gs.cells[ ((i-GRID_BORDER)/CELL_WIDTH) ][ ((j-GRID_BORDER)/CELL_HEIGHT) ] == CELL_HUB)
                {
                     blue = 203;
                     green = 192;
                     red = 255;
                     alpha = 0;
                }

                if(gs.cells[ ((i-GRID_BORDER)/CELL_WIDTH) ][ ((j-GRID_BORDER)/CELL_HEIGHT) ] == CELL_EXIT)
                {
                     blue = 128;
                     green = 0;
                     red = 128;
                     alpha = 0;
                }

                if(gs.cells[ ((i-GRID_BORDER)/CELL_WIDTH) ][ ((j-GRID_BORDER)/CELL_HEIGHT) ] == CELL_AREA1)
                {
                     blue = 204;
                     green = 204;
                     red = 0;
                     alpha = 0;
                }

                if(gs.cells[ ((i-GRID_BORDER)/CELL_WIDTH) ][ ((j-GRID_BORDER)/CELL_HEIGHT) ] == CELL_AREA2)
                {
                     blue = 100;
                     green = 100;
                     red = 200;
                     alpha = 0;
                }

                if(gs.cells[ ((i-GRID_BORDER)/CELL_WIDTH) ][ ((j-GRID_BORDER)/CELL_HEIGHT) ] == CELL_AREA3)
                {
                     blue = 0;
                     green = 255;
                     red = 255;
                     alpha = 0;
                }

                if(gs.cells[ ((i-GRID_BORDER)/CELL_WIDTH) ][ ((j-GRID_BORDER)/CELL_HEIGHT) ] == CELL_AREA4)
                {
                     blue = 50;
                     green = 255;
                     red = 155;
                     alpha = 0;
                }

                if(gs.cells[ ((i-GRID_BORDER)/CELL_WIDTH) ][ ((j-GRID_BORDER)/CELL_HEIGHT) ] == CELL_AREA5)
                {
                     blue = 100;
                     green = 155;
                     red = 55;
                     alpha = 0;
                }

                if(gs.cells[ ((i-GRID_BORDER)/CELL_WIDTH) ][ ((j-GRID_BORDER)/CELL_HEIGHT) ] == CELL_END)
                {
                    blue = 0;
                    green = 255;
                    red = 0;
                    alpha = 0;
                }
            }

            *pixel++ = ((alpha << 24) | (red << 16) | (green << 8) | blue);
        }
        row += pitch;
    }
}

void drawGrid(framebuffer buffer)
{
    int pitch = buffer.w*buffer.bytesperpixel;
    uint8_t *row = buffer.data;

    for(int y = 0; y < buffer.h; y++)
    {
        uint32_t *pixel = (uint32_t *)row;
        for(int x = 0; x < buffer.w; x++)
        {
            if(((x+GRID_BORDER)%CELL_WIDTH == 0 && x >= GRID_BORDER && x <= buffer.w-GRID_BORDER && y >= GRID_BORDER && y <= buffer.h-GRID_BORDER) || 
               ((y+GRID_BORDER)%CELL_HEIGHT == 0 && y >= GRID_BORDER && y <= buffer.h-GRID_BORDER && x >= GRID_BORDER && x <= buffer.w-GRID_BORDER) || 
               (x == buffer.w-GRID_BORDER && y >= GRID_BORDER && y <= buffer.h-GRID_BORDER) || 
               (y == buffer.h-GRID_BORDER && x >= GRID_BORDER && x <= buffer.w-GRID_BORDER)) 
            {
                uint8_t blue = 255;
                uint8_t green = 255;
                uint8_t red = 255;
                uint8_t alpha = 255;

                *pixel = ((alpha << 24) | (red << 16) | (green << 8) | blue);
            }
            *pixel++;
        }
        row += pitch;
    }
}

// Checks whether cells in area are CELL_OFF
// Seems to be working...
bool checkAreaIsOff(gridstatus *gs, int x, int y, int w, int h)
{
    // If w/h are negative...
    //if(w < 0 && x+w+1 > 0)
    //{
    //     x += w+1;
    //     w = -w;
    //}
    //if(h < 0 && y+h+1 > 0)
    //{
    //     y += h+1;
    //     h = -h;
    //}

    // Look for something that isn't CELL_OFF in area 
    for(int i = x; i < x+w; i++)
        for(int j = y; j < y+h; j++)
        {
            if(i < 0 || i > GRID_WIDTH-1 || j < 0 || j > GRID_HEIGHT-1 || gs->cells[i][j] != CELL_OFF)
                return false;
            //gs->cells[i][j] = CELL_END; // To test it works
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
    for(int i = 0; i < size; size--) 
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
            for(int i = 0; i < 4; i++) 
                array[i] = '0';
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

struct travelCell 
{
    int x, y;
    int prevx, prevy;
};

struct travelCellList 
{
    travelCell *travelCells;
    int count;
    uint32_t mallocSize;
};

travelCell makeTravelCell(int x, int y, int prevx, int prevy)
{
     travelCell tc;
     tc.x = x;
     tc.y = y;
     tc.prevx = prevx;
     tc.prevy = prevy;

     return tc;
}

void addCell(travelCellList *list, int x, int y, int prevx, int prevy) 
{
    list->count += 1;

    // If more memory is needed, realloc more
    if(list->mallocSize/list->count < sizeof(travelCell))
    {
        list->mallocSize += 4096;
        list->travelCells = (travelCell *)realloc(list->travelCells, list->mallocSize);
        //printf("malloc resized: %d\n", list->mallocSize);
    }

    travelCell tc;
    tc.x = x;
    tc.y = y;
    tc.prevx = prevx;
    tc.prevy = prevy;
    list->travelCells[list->count-1] = tc;
}

travelCell findCell(travelCellList *list, int x, int y)
{
     for(int i = 0; i < list->count; i++) 
     {
         if(list->travelCells[i].x == x && list->travelCells[i].y == y)
             return list->travelCells[i];
     }

     return makeTravelCell(-1, -1, -1, -1); // cell not found in list, should never happen
}

bool inList(travelCellList *list, int x, int y)
{
     for(int i = 0; i < list->count; i++) 
     {
         if(list->travelCells[i].x == x && list->travelCells[i].y == y)
             return true;
     }

     return false;
}

bool canTravelBetween(gridstatus *gs, int x1, int y1, int x2, int y2, cellstatus CS)
{
    travelCellList list;
    list.mallocSize = 4096;
    list.travelCells = (travelCell *)malloc(list.mallocSize);
    list.count = 0;
    addCell(&list, x1, y1, -1, -1); // -1 for prevx/y means origin cell
    
    travelCell current = list.travelCells[list.count-1];

    do 
    {
        // If starting cell is not CS, return false
        //if(gs->cells[x1][y1] != CS)
        //    return false;
        // If we've reached destination, return true
        if(x2 == current.x && y2 == current.y)
        {
            free(list.travelCells);
            return true;
        }
        // If destination is to the right && we can move right && that cell hasn't been explored, move
        else if(x2 > current.x && gs->cells[current.x+1][current.y] == CS && !inList(&list, current.x+1, current.y))
        {
            addCell(&list, current.x + 1, current.y, current.x, current.y);
            current = list.travelCells[list.count-1];
        } 
        // If destination is to the left... etc.
        else if(x2 < current.x && gs->cells[current.x-1][current.y] == CS && !inList(&list, current.x-1, current.y))
        {
            addCell(&list, current.x - 1, current.y, current.x, current.y); 
            current = list.travelCells[list.count-1];
        }
        else if(y2 > current.y && gs->cells[current.x][current.y+1] == CS && !inList(&list, current.x, current.y+1))
        {
            addCell(&list, current.x, current.y + 1, current.x, current.y);
            current = list.travelCells[list.count-1];
        }
        else if(y2 < current.y && gs->cells[current.x][current.y-1] == CS && !inList(&list, current.x, current.y-1))
        {
             addCell(&list, current.x, current.y - 1, current.x, current.y);
             current = list.travelCells[list.count-1];
        }
        // If we cannot move closer, move in an availble direction instead
        else if(current.x+1 < GRID_WIDTH && gs->cells[current.x+1][current.y] == CS && !inList(&list, current.x+1, current.y))
        {
            addCell(&list, current.x + 1, current.y, current.x, current.y);
            current = list.travelCells[list.count-1];
        }
        else if(current.x-1 >= 0 && gs->cells[current.x-1][current.y] == CS && !inList(&list, current.x-1, current.y))
        {
            addCell(&list, current.x - 1, current.y, current.x, current.y);
            current = list.travelCells[list.count-1];
        }
        else if(current.y+1 < GRID_HEIGHT && gs->cells[current.x][current.y+1] == CS && !inList(&list, current.x, current.y+1))
        {
            addCell(&list, current.x, current.y + 1, current.x, current.y);
            current = list.travelCells[list.count-1];
        }
        else if(current.y-1 >= 0 && gs->cells[current.x][current.y-1] == CS && !inList(&list, current.x, current.y-1))
        {
            addCell(&list, current.x, current.y - 1, current.x, current.y);
            current = list.travelCells[list.count-1];
        }
        // If we cannot move anywhere else and we're not at the origin, move to previously explored point
        else if(current.prevx != -1)
        {
            current = findCell(&list, current.prevx, current.prevy);
        }
        // Continue while we're not at the origin... 
    } while(current.prevx != -1 || 
            // or we are at the origin but still have somewhere to move
            (current.x-1 >= 0 && gs->cells[current.x-1][current.y] == CS && !inList(&list, current.x-1, current.y)) ||
            (current.x+1 < GRID_WIDTH && gs->cells[current.x+1][current.y] == CS && !inList(&list, current.x+1, current.y)) ||
            (current.y-1 >= 0 && gs->cells[current.x][current.y-1] == CS && !inList(&list, current.x, current.y-1)) ||
            (current.y+1 < GRID_HEIGHT && gs->cells[current.x][current.y+1] == CS && !inList(&list, current.x, current.y+1)));
    
    free(list.travelCells);
    return false;
}

void joinTwoCells(gridstatus *gs, int x1, int y1, int x2, int y2, cellstatus CS) 
{
    // Pick a cell between the two points
    int x, y; 
    if(x1 > x2) 
        x = x2 + ((x1-x2)/2);
    else
        x = x1 + ((x2-x1)/2);
    if(y1 > y2)
        y = y2 + ((y1-y2)/2);
    else
        y = y1 + ((y2-y1)/2);

    // Create a blob between the two points
    if(!canTravelBetween(gs, x1, y1, x2, y2, CS))
    {
        // TODO: Make blobSize a function argument?
        int blobSize = 10;
        while(blobSize > 0)
        {
            // NOTE: This changes x and y, giving path some character
            createBlob(gs, &x, &y, blobSize, CS); 
        }
    }

    // Test whether you can travel between the two points
    while(!canTravelBetween(gs, x1, y1, x2, y2, CS))
    {
        // If not, recursively call joinTwoCells() twice
        joinTwoCells(gs, x, y, x2, y2, CS);
        joinTwoCells(gs, x1, y1, x, y, CS);
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
    }

    pl->list[pl->count-1] = p;
}

bool canDrawLineBetween(gridstatus *gs, point a, point b, int maxLength)
{
     
}

point pickRandomCell(gridstatus *gs, point exitCell, int maxLength, cellstatus CS)
{
    point result;
    // Pick a random point
    // Is there a line between the point and exitCell? -- bresenham?
    // Is the line not exceed maxLength??
    // If so, return the point
    // If not, pick another point

    // Place a random point

    while(1)
    {
        result.x = stb_rand()%GRID_WIDTH;
        result.y = stb_rand()%GRID_HEIGHT;
        if(gs->cells[result.x][result.y] == CELL_OFF)
        {
            gs->cells[result.x][result.y] = CS;
            break;
        }
    }

    return result;
}

int main(int argc, char* args[])
{
    if(SDL_Init(SDL_INIT_VIDEO | SDL_INIT_AUDIO) != 0)
    {
        printf("Vid/Sound init failed. SDL_Error: %s\n", SDL_GetError());
    }

    // Create Window
    SDL_Window *window = SDL_CreateWindow("Proc Gen",
                                          SDL_WINDOWPOS_UNDEFINED,
                                          SDL_WINDOWPOS_UNDEFINED,
                                          SCREEN_WIDTH,
                                          SCREEN_HEIGHT,
                                          SDL_WINDOW_RESIZABLE);

    SDL_Renderer *renderer = SDL_CreateRenderer(window, -1, 0);

    framebuffer buffer;
    SDL_GetWindowSize(window, &buffer.w, &buffer.h);
    buffer.bytesperpixel = 4;

    buffer.data = (uint8_t *)malloc(buffer.w * buffer.h * 
                                    buffer.bytesperpixel);
    
    SDL_Texture *texture = SDL_CreateTexture(renderer,
                                            SDL_PIXELFORMAT_ARGB8888,
                                            SDL_TEXTUREACCESS_STREAMING,
                                            buffer.w, buffer.h);
    // Init gridstatus
    gridstatus grid;
    grid.w = GRID_WIDTH;
    grid.h = GRID_HEIGHT;

    clear(&grid);

    int currentX = GRID_WIDTH/2;
    int currentY = GRID_HEIGHT/2;
    grid.cells[currentX][currentY] = CELL_HUB_START;
    int hubBlobSize = 100;
    int size = hubBlobSize;
    int status = 0;

    // Pre loop stuff
    int running = 1;
    int fullscreen = 0;
    int newTime = 0;
    int prevTime;
    int deltaTime;
    SDL_Event e;

    while(running) 
    {
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
                            //printf("size: %d\n", size);
                            size = hubBlobSize;
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


        // Test for joinTwoCells and canTravelBetween
        //int num = 0;
        //while(!canTravelBetween(&grid, num, num, GRID_WIDTH-1, GRID_HEIGHT-1))
        //{
        //    joinTwoCells(&grid, num, num, GRID_WIDTH-1, GRID_HEIGHT-1);
        //    //printf("canTravel0,0: %d\n", canTravelBetween(&grid, 0, 0, GRID_WIDTH-1, GRID_HEIGHT-1));
        //    //printf("canTravelnum,num: %d\n", canTravelBetween(&grid, num, num, GRID_WIDTH-1, GRID_HEIGHT-1));
        //}
        
        // Do procedural generation
        if(status == 0)
        {
            createBlob(&grid, &currentX, &currentY, size, CELL_HUB);
            createExitCells(&grid, CELL_HUB);
            createWalls(&grid, CELL_HUB); 
            createWalls(&grid, CELL_HUB_START); 

            // Find all CELL_EXITs
            pointList pl;
            pl.count = 0;
            pl.mallocSize = 4096;
            pl.list = (point *)malloc(pl.mallocSize);
            
            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_EXIT)
                       addPoint(&pl, makePoint(i, j));
                }
            }

            // Chooses a random CELL_EXIT for area1 
            point p = pl.list[stb_rand()%pl.count];
            grid.cells[p.x][p.y] = CELL_AREA1;
            
            // Create area1
            createBlob(&grid, &p.x, &p.y, size*2, CELL_AREA1);
            createExitCells(&grid, CELL_AREA1);
            createWalls(&grid, CELL_AREA1); 

            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_EXIT)
                    {
                        // If none of the surrounding cells is CELL_OFF
                        // grid.cells[i[j] = CELL_SHORTCUT;
                        if(i-1 >= 0 && i+1 < GRID_WIDTH && j-1 >= 0 && j+1 < GRID_HEIGHT &&
                           grid.cells[i-1][j] != CELL_OFF && grid.cells[i+1][j] != CELL_OFF &&
                           grid.cells[i][j-1] != CELL_OFF && grid.cells[i][j+1] != CELL_OFF)
                        {
                             grid.cells[i][j] = CELL_SHORTCUT;
                        }
                        else
                            grid.cells[i][j] = CELL_WALL;

                    }
                }
            }

            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_WALL)
                        grid.cells[i][j] = CELL_OFF;
                }
            }
            createExitCells(&grid, CELL_HUB);
            createExitCells(&grid, CELL_AREA1);
            createWalls(&grid, CELL_HUB); 
            createWalls(&grid, CELL_HUB_START); 
            createWalls(&grid, CELL_AREA1); 

            // Empty pl.list
            free(pl.list);
            pl.mallocSize = 4096;
            pl.count = 0;
            pl.list = (point *)malloc(pl.mallocSize);

            // Add current CELL_EXITs to list
            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_EXIT)
                       addPoint(&pl, makePoint(i, j));
                }
            }
            
            // Choose a random CELL_EXIT for area2
            p = pl.list[stb_rand()%pl.count];
            grid.cells[p.x][p.y] = CELL_AREA2;
            
            // Create area2
            createBlob(&grid, &p.x, &p.y, size*2, CELL_AREA2);
            createExitCells(&grid, CELL_AREA2);
            createWalls(&grid, CELL_AREA2); 

            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_EXIT)
                    {
                        // If none of the surrounding cells is CELL_OFF
                        // grid.cells[i[j] = CELL_SHORTCUT;
                        if(i-1 >= 0 && i+1 < GRID_WIDTH && j-1 >= 0 && j+1 < GRID_HEIGHT &&
                           grid.cells[i-1][j] != CELL_OFF && grid.cells[i+1][j] != CELL_OFF &&
                           grid.cells[i][j-1] != CELL_OFF && grid.cells[i][j+1] != CELL_OFF)
                        {
                             grid.cells[i][j] = CELL_SHORTCUT;
                        }
                        else
                            grid.cells[i][j] = CELL_WALL;

                    }
                }
            }

            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_WALL)
                        grid.cells[i][j] = CELL_OFF;
                }
            }
            createExitCells(&grid, CELL_HUB);
            createExitCells(&grid, CELL_AREA1);
            createExitCells(&grid, CELL_AREA2);
            createWalls(&grid, CELL_HUB); 
            createWalls(&grid, CELL_HUB_START); 
            createWalls(&grid, CELL_AREA1); 
            createWalls(&grid, CELL_AREA2); 

            // Empty pl.list
            free(pl.list);
            pl.mallocSize = 4096;
            pl.count = 0;
            pl.list = (point *)malloc(pl.mallocSize);

            // Add current CELL_EXITs to list
            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_EXIT)
                       addPoint(&pl, makePoint(i, j));
                }
            }
            
            // Choose a random CELL_EXIT for area3
            p = pl.list[stb_rand()%pl.count];
            grid.cells[p.x][p.y] = CELL_AREA3;
            
            // Create area2
            createBlob(&grid, &p.x, &p.y, size*2, CELL_AREA3);
            createExitCells(&grid, CELL_AREA3);
            createWalls(&grid, CELL_AREA3); 

            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_EXIT)
                    {
                        // If none of the surrounding cells is CELL_OFF
                        // grid.cells[i[j] = CELL_SHORTCUT;
                        if(i-1 >= 0 && i+1 < GRID_WIDTH && j-1 >= 0 && j+1 < GRID_HEIGHT &&
                           grid.cells[i-1][j] != CELL_OFF && grid.cells[i+1][j] != CELL_OFF &&
                           grid.cells[i][j-1] != CELL_OFF && grid.cells[i][j+1] != CELL_OFF)
                        {
                             grid.cells[i][j] = CELL_SHORTCUT;
                        }
                        else
                            grid.cells[i][j] = CELL_WALL;

                    }
                }
            }
            
            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_WALL)
                        grid.cells[i][j] = CELL_OFF;
                }
            }
            createExitCells(&grid, CELL_HUB);
            createExitCells(&grid, CELL_AREA1);
            createExitCells(&grid, CELL_AREA2);
            createExitCells(&grid, CELL_AREA3);
            createWalls(&grid, CELL_HUB); 
            createWalls(&grid, CELL_HUB_START); 
            createWalls(&grid, CELL_AREA1); 
            createWalls(&grid, CELL_AREA2); 
            createWalls(&grid, CELL_AREA3); 

            // Empty pl.list
            free(pl.list);
            pl.mallocSize = 4096;
            pl.count = 0;
            pl.list = (point *)malloc(pl.mallocSize);

            // Add current CELL_EXITs to list
            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_EXIT)
                       addPoint(&pl, makePoint(i, j));
                }
            }
            
            // Choose a random CELL_EXIT for area3
            p = pl.list[stb_rand()%pl.count];
            grid.cells[p.x][p.y] = CELL_AREA4;
            
            // Create area2
            createBlob(&grid, &p.x, &p.y, size*2, CELL_AREA4);
            createExitCells(&grid, CELL_AREA4);
            createWalls(&grid, CELL_AREA4); 

            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_EXIT)
                    {
                        // If none of the surrounding cells is CELL_OFF
                        // grid.cells[i[j] = CELL_SHORTCUT;
                        if(i-1 >= 0 && i+1 < GRID_WIDTH && j-1 >= 0 && j+1 < GRID_HEIGHT &&
                           grid.cells[i-1][j] != CELL_OFF && grid.cells[i+1][j] != CELL_OFF &&
                           grid.cells[i][j-1] != CELL_OFF && grid.cells[i][j+1] != CELL_OFF)
                        {
                             grid.cells[i][j] = CELL_SHORTCUT;
                        }
                        else
                            grid.cells[i][j] = CELL_WALL;

                    }
                }
            }

            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_WALL)
                        grid.cells[i][j] = CELL_OFF;
                }
            }
            createExitCells(&grid, CELL_HUB);
            createExitCells(&grid, CELL_AREA1);
            createExitCells(&grid, CELL_AREA2);
            createExitCells(&grid, CELL_AREA3);
            createExitCells(&grid, CELL_AREA4);
            createWalls(&grid, CELL_HUB); 
            createWalls(&grid, CELL_HUB_START); 
            createWalls(&grid, CELL_AREA1); 
            createWalls(&grid, CELL_AREA2); 
            createWalls(&grid, CELL_AREA3); 
            createWalls(&grid, CELL_AREA4); 

            // Empty pl.list
            free(pl.list);
            pl.mallocSize = 4096;
            pl.count = 0;
            pl.list = (point *)malloc(pl.mallocSize);

            // Add current CELL_EXITs to list
            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_EXIT)
                       addPoint(&pl, makePoint(i, j));
                }
            }
            
            // Choose a random CELL_EXIT for area3
            p = pl.list[stb_rand()%pl.count];
            grid.cells[p.x][p.y] = CELL_AREA5;
            
            // Create area2
            createBlob(&grid, &p.x, &p.y, size*2, CELL_AREA5);
            createExitCells(&grid, CELL_AREA5);
            createWalls(&grid, CELL_AREA5); 

            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_EXIT)
                    {
                        // If none of the surrounding cells is CELL_OFF
                        // grid.cells[i[j] = CELL_SHORTCUT;
                        if(i-1 >= 0 && i+1 < GRID_WIDTH && j-1 >= 0 && j+1 < GRID_HEIGHT &&
                           grid.cells[i-1][j] != CELL_OFF && grid.cells[i+1][j] != CELL_OFF &&
                           grid.cells[i][j-1] != CELL_OFF && grid.cells[i][j+1] != CELL_OFF)
                        {
                             grid.cells[i][j] = CELL_SHORTCUT;
                        }
                        //else
                        //    grid.cells[i][j] = CELL_WALL;

                    }
                }
            }
            //int lengthFromP = 30;
            //point area1 = pickRandomCell(&grid, p, lengthFromP, CELL_AREA1);


            //// Create path between EXIT and random point
            //joinTwoCells(&grid, area1.x, area1.y, p.x, p.y, CELL_AREA1);
            //createExitCells(&grid, CELL_AREA1);
            //createWalls(&grid, CELL_AREA1); 

            //free(pl.list);

            status--;
        }






        if(size == -2)
        {
            pointList pl;
            pl.count = 0;
            pl.mallocSize = 4096;
            pl.list = (point *)malloc(pl.mallocSize);

            // Add new exit cells 
            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_EXIT)
                       addPoint(&pl, makePoint(i, j));
                }
            }

            // Random Area2 cells
            point area2, p;
            int count = 0;
            do 
            {
                count++;
                p = pl.list[stb_rand()%pl.count];
                while(1)
                {
                    area2.x = stb_rand()%GRID_WIDTH;
                    area2.y = stb_rand()%GRID_HEIGHT;
                    if(grid.cells[area2.x][area2.y] == CELL_OFF)
                        break;
                }
                 
            // p.x and p.y is not CELL_OFF and therefore must be first canTravelBetween starting point
            } while(!canTravelBetween(&grid, p.x, p.y, area2.x, area2.y, CELL_OFF) && count < 30);

            if(count == 30)
            {
                // We're done? 
                size = -10;
            }
            else 
            {
                grid.cells[area2.x][area2.y] = CELL_AREA2;
                grid.cells[p.x][p.y] = CELL_AREA2;
                joinTwoCells(&grid, area2.x, area2.y, p.x, p.y, CELL_AREA2);
                createExitCells(&grid, CELL_AREA2);
                createWalls(&grid, CELL_AREA2); 
                size = -3;
            }
        } 

        if(size == -3)
        {
            pointList pl;
            pl.count = 0;
            pl.mallocSize = 4096;
            pl.list = (point *)malloc(pl.mallocSize);

            // Add new exit cells 
            for(int i = 0; i < GRID_WIDTH; i++) 
            {
                for(int j = 0; j < GRID_HEIGHT; j++) 
                {
                    if(grid.cells[i][j] == CELL_EXIT)
                       addPoint(&pl, makePoint(i, j));
                }
            }

            // Random Area3 cells
            point area3, p;
            int count = 0;
            do 
            {
                count++;
                p = pl.list[stb_rand()%pl.count];
                while(1)
                {
                    area3.x = stb_rand()%GRID_WIDTH;
                    area3.y = stb_rand()%GRID_HEIGHT;
                    if(grid.cells[area3.x][area3.y] == CELL_OFF)
                        break;
                }
                 
            // p.x and p.y is not CELL_OFF and therefore must be canTravelBetween starting point
            // because if cTB() end point isn't CELL_OFF, it always fails
            } while(!canTravelBetween(&grid, p.x, p.y, area3.x, area3.y, CELL_OFF) && count < 30);

            if(count == 30)
            {
                // We're done? 
                size = -10;
            }
            else 
            {
                grid.cells[area3.x][area3.y] = CELL_AREA3;
                grid.cells[p.x][p.y] = CELL_AREA3;
                joinTwoCells(&grid, area3.x, area3.y, p.x, p.y, CELL_AREA3);
                createExitCells(&grid, CELL_AREA3);
                createWalls(&grid, CELL_AREA3); 
                size = -4;
            }
        } 

        // Write to buffer
        fillGrid(buffer, grid);
        drawGrid(buffer);

        // Update the screen
        SDL_UpdateTexture(texture, 0, buffer.data, 
                          buffer.w * buffer.bytesperpixel);
        SDL_RenderCopy(renderer, texture, 0, 0);
        SDL_RenderPresent(renderer);

        // Calling SDL_GetTicks() twice for accurate deltaTime. Necessary?
        newTime = (int)SDL_GetTicks();
        deltaTime = newTime - prevTime;
        //printf("prevTime: %d newTime: %d deltaTime: %d\n", prevTime, newTime, deltaTime);
        //if(deltaTime < 20)
        //    SDL_Delay(20 - deltaTime);
        prevTime = (int)SDL_GetTicks();
    }

    SDL_SetWindowFullscreen(window, 0);

    // Necessary?
    SDL_DestroyRenderer(renderer);
    SDL_DestroyWindow(window);

    Mix_Quit();
    SDL_Quit();

    return 0;
}
