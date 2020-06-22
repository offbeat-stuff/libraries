package sdl

/*
#include "main.h"

void handleInput(SDL_Event *event){
	if(event->type==SDL_QUIT){
		SDL_Quit();
	}
}

#cgo amd64 CFLAGS: -I include
#cgo LDFLAGS: -lSDL2
*/
import "C"

//BeginLoop begin main render loop
func BeginLoop(fps int, loop func(), running *bool) {
	var mpf = C.Uint32(1000 / fps)
	var event C.SDL_Event
	for *running {
		var start = C.SDL_GetTicks()

		C.SDL_PollEvent(&event)

		C.handleInput(&event)

		loop()

		var passed = C.SDL_GetTicks() - start
		if passed > mpf {
			continue
		} else {
			C.SDL_Delay(mpf - passed)
		}
	}
}
