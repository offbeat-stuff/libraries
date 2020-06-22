package main

import (
	"fmt"

	"github.com/offbeat-stuff/libraries/sdl-go/sdl"
)

func update() {
	frames++
	if frames%60 == 0 {
		fmt.Println(frames / 60)
		if frames/60 == 10 {
			running = false
			sdl.Quit()
		}
	}
}

func draw() {

}
