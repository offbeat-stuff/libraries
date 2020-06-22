#include <SDL2/SDL.h>
#include <SDL2/SDL_events.h>
#include <SDL2/SDL_render.h>
#include <SDL2/SDL_timer.h>
#include <SDL2/SDL_video.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

// #include "sdl_init.h"
// #include "loop.h"

// int exitCode = 0;

// int init()
// {
// targetFPS = 60;
// sdl_init(SDL_INIT_VIDEO | SDL_INIT_EVENTS);
// updateRes();
// printf("%i\n%i\n", res_width, res_height);
// sdl_createWindowAndRenderer("Hello World", 0, 0, res_width / 2, res_height / 2, SDL_WINDOW_SHOWN, SDL_RENDERER_ACCELERATED | SDL_RENDERER_PRESENTVSYNC);
//
// loop();
//
// SDL_DestroyRenderer(win.renderer);
// SDL_DestroyWindow(win.window);
// SDL_Quit();
// return exitCode;
// }