package main

import (
	"fmt"

	"github.com/offbeat-stuff/libraries/sdl-go/sdl"
	"github.com/offbeat-stuff/libraries/vg"
)

var frames = 0
var running = true

func main() {
	fmt.Println("starting...")
	sdl.Init(sdl.InitVideo)

	res, refreshRate, err := sdl.GetDisplayProperties()
	if err != nil {
		panic(err)
	}
	sdl.CreateWindow("Hello", vg.Vec2{X: res.X / 2, Y: 0}, res.Mult(0.5), sdl.WindowShown)

	sdl.BeginLoop(refreshRate, func() {
		update()
		draw()
	}, &running)
}
