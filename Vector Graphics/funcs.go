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

var _samplePoints = [4]vec2{vec2{0.25, 0.25}, vec2{0.75, 0.25}, vec2{0.25, 0.75}, vec2{0.75, 0.75}}

func circle(centr vec2, r float64, src *image.RGBA, clr color.RGBA) {
	var rSq = sqr(r)
	for x := -r; x < r; x++ {
		for y := -r; y < r; y++ {
			var cClr color.RGBA
			for _, z := range _samplePoints {
				cPos := centr.add(vec2{x, y}).add(z)
				pClr := src.At(int(cPos.x), int(cPos.y)).(color.RGBA)
				if z.add(vec2{x, y}).magSq() < rSq {
					cClr = addClr(cClr, multClr(clr, 0.25))
				} else {
					cClr = addClr(cClr, multClr(pClr, 0.25))
				}
			}
			src.SetRGBA(int(centr.x+x), int(centr.y+y), cClr)
		}
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
