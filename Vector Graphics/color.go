package main

import "image/color"

func addClr(a, b color.RGBA) color.RGBA {
	return color.RGBA{a.R + b.R, a.G + b.G, a.B + b.B, a.A + b.A}
}

func multClr(a color.RGBA, b float64) color.RGBA {
	return color.RGBA{uint8(float64(a.R) * b), uint8(float64(a.G) * b), uint8(float64(a.B) * b), uint8(float64(a.A) * b)}
}
