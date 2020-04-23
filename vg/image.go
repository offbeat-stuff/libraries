package vg

import (
	"image"
	"image/png"
	"os"
)

func createImg(width, height int) *image.RGBA {
	oP := image.NewRGBA(image.Rect(0, 0, width, height))
	return oP
}

func saveImg(src image.Image, oP string) {
	file, err := os.Create(oP)
	if err != nil {
		panic(err)
	}
	png.Encode(file, src)
}
