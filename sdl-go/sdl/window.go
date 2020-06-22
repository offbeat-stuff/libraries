package sdl

/*
#include "main.h"

#cgo amd64 CFLAGS: -I include
#cgo LDFLAGS: -lSDL2
*/
import "C"

// type window struct {
// 	win       *C.SDL_Window
// 	pos, size vg.Vec2
// }

//CreateWindow creates a window
func CreateWindow(title string, x, y, w, h int32, flags ...C.uint) {
	winFlag := C.uint(0)
	for i := range flags {
		winFlag |= flags[i]
	}
	C.SDL_CreateWindow(C.CString(title), C.int(x), C.int(y), C.int(w), C.int(h), winFlag)
}
