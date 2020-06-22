package sdl

import "github.com/offbeat-stuff/libraries/vg"

/*
#include "main.h"

#cgo amd64 CFLAGS: -I include
#cgo LDFLAGS: -lSDL2
*/
import "C"

//GetDisplayProperties |-> returns default desktop resolution && refresh rate
func GetDisplayProperties() (vg.Vec2, int, error) {
	var dm C.SDL_DisplayMode

	if C.SDL_GetDesktopDisplayMode(0, &dm) != 0 {
		return vg.Vec2{0, 0}, int(0), GetError()
	}
	return vg.Vec2{float64(dm.w), float64(dm.h)}, int(dm.refresh_rate), nil
}
