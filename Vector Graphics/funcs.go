package main

import (
	"image"
	"image/color"
)

func line(a, b vec2, src *image.RGBA, clr color.Color) {
	// The line
	var road = b.add(a.neg())
	inc := 1 / road.mag()
	for i := 0.0; i < 1; i += inc {
		pos := a.add(road.mult(i))
		src.Set(int(pos.x), int(pos.y), clr)
	}
}

func showCurve(pts []vec2, src *image.RGBA, clr color.Color) {
	var cPt = pts[0]
	for i := 0.01; i < 1; i += 0.01 {
		nxt := curveLerp(pts, i)
		line(cPt, nxt, src, clr)
		cPt = nxt
	}
}

func curveLerp(pts []vec2, prog float64) vec2 {
	for len(pts) > 1 {
		nPts := make([]vec2, len(pts)-1)
		for i := 0; i < len(nPts); i++ {
			nPts[i] = pts[i+1].add(pts[i].neg()).mult(prog).add(pts[i])
		}
		pts = nPts
	}
	return pts[0]
}

func rect(a, b vec2, src *image.RGBA, clr color.Color) {
	for i := int(a.x); i < int(b.x); i++ {
		for j := int(a.y); j < int(b.y); j++ {
			src.Set(i, j, clr)
		}
	}
}
