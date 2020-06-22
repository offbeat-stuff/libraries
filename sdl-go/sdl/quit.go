package sdl

/*
#include "main.h"

#cgo amd64 CFLAGS: -I include
#cgo LDFLAGS: -lSDL2
*/
import "C"

//Quit sdl
func Quit() {
	C.SDL_Quit()
}
