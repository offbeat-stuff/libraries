package vg

import (
	"image"
	"image/png"
	"os"
)

//CreateImg creates a src image based on width and height
// |> You can make your own source too
func CreateImg(width, height int) *image.RGBA {
	oP := image.NewRGBA(image.Rect(0, 0, width, height))
	return oP
}

//SaveImg saves img to destination
func SaveImg(src image.Image, oP string) {
	file, err := os.Create(oP)
	if err != nil {
		panic(err)
	}
	png.Encode(file, src)
}
