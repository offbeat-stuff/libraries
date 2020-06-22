package sdl

/*
#include "main.h"

#cgo amd64 CFLAGS: -I include
#cgo LDFLAGS: -lSDL2
*/
import "C"
import "errors"

//Init an sdl subsystem
func Init(flags ...C.int) error {
	var flag = C.int(0)
	for i := range flags {
		flag |= flags[i]
	}
	if C.SDL_Init(C.uint(flag)) != 0 {
		return GetError()
	}
	return nil
}

//GetError return sdl error that occurred recently
func GetError() error {
	if err := C.SDL_GetError(); err != nil {
		gostr := C.GoString(err)
		// SDL_GetError returns "an empty string if there hasn't been an error message"
		if len(gostr) > 0 {
			return errors.New(gostr)
		}
	}
	return nil
}
