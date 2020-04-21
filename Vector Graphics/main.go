package main

import (
	"image"
	"image/color"
	"strconv"
	"sync"
)

var cPts = []vec2{{125, 125}, {375, 125}, {375, 375}, {125, 375}}

func main() {
	var wg sync.WaitGroup
	wg.Add(500)
	for x := 0; x < 500; x++ {
		canv := createImg(500, 500)
		rect(vec2{}, vec2{500, 500}, canv, color.Black)
		circle(vec2{250, 250}, 30, canv, color.RGBA{255, 60, 120, 255})
		drawFrame(x, canv, &wg)
	}
	wg.Wait()
}

func drawFrame(x int, src *image.RGBA, wg *sync.WaitGroup) {
	var xPts = make([]vec2, 4)
	for i := range xPts {
		xPts[i] = cPts[i].add(vec2{100, 0}.rotate((float64(x) / 40) + (float64(i) * 0.5)))
	}
	for i := 0; i < 3; i++ {
		line(xPts[i], xPts[i+1], src, color.RGBA{100, 244, 255, 255})
	}
	showCurve(xPts, src, color.White)
	saveImg(src, "images/"+strconv.FormatInt(int64(x), 10)+".png")
	wg.Done()
}
