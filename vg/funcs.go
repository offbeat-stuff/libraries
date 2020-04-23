package vg

import (
	"image"
	"image/color"
)

func line(a, b Vec2, src *image.RGBA, clr color.Color) {
	// The line
	var road = b.add(a.neg())
	inc := 1 / road.mag()
	for i := 0.0; i < 1; i += inc {
		pos := a.add(road.mult(i))
		src.Set(int(pos.x), int(pos.y), clr)
	}
}

var samplePoints = [4]Vec2{{0.25, 0.25}, {0.75, 0.25}, {0.25, 0.75}, {0.75, 0.75}}

//Circle function to draw circles based on center and radius
func Circle(centr Vec2, r float64, src *image.RGBA, clr color.RGBA) {
	var rSq = sqr(r)
	for x := -r; x < r; x++ {
		for y := -r; y < r; y++ {
			var cClr color.RGBA
			for _, z := range samplePoints {
				cPos := centr.add(Vec2{x, y}).add(z)
				pClr := src.At(int(cPos.x), int(cPos.y)).(color.RGBA)
				if z.add(Vec2{x, y}).magSq() < rSq {
					cClr = addClr(cClr, multClr(clr, 0.25))
				} else {
					cClr = addClr(cClr, multClr(pClr, 0.25))
				}
			}
			src.SetRGBA(int(centr.x+x), int(centr.y+y), cClr)
		}
	}
}

//Curve function to draw curves based on control points
func Curve(pts []Vec2, src *image.RGBA, clr color.Color) {
	var cPt = pts[0]
	for i := 0.01; i < 1; i += 0.01 {
		nxt := curveLerp(pts, i)
		line(cPt, nxt, src, clr)
		cPt = nxt
	}
}

func curveLerp(pts []Vec2, prog float64) Vec2 {
	for len(pts) > 1 {
		nPts := make([]Vec2, len(pts)-1)
		for i := 0; i < len(nPts); i++ {
			nPts[i] = pts[i+1].add(pts[i].neg()).mult(prog).add(pts[i])
		}
		pts = nPts
	}
	return pts[0]
}

//Rect Func to draw rects
func Rect(a, b Vec2, src *image.RGBA, clr color.Color) {
	for i := int(a.x); i < int(b.x); i++ {
		for j := int(a.y); j < int(b.y); j++ {
			src.Set(i, j, clr)
		}
	}
}
