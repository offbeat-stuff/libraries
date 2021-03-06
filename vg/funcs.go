package vg

import (
	"image"
	"image/color"
)

//Line to draw lines from a,b vectors defining start and end and width a float64
func Line(a, b Vec2, width float64, src *image.RGBA, clr color.Color) {
	// The line
	var road = b.Add(a.Neg())
	//Rotated road
	var rLn = road.Norm()
	rLn = Vec2{-rLn.Y, rLn.X}.Mult(width / 2)
	inc := 1 / road.Mag()
	for i := 0.0; i < 1; i += inc {
		pos := a.Add(road.Mult(i))
		line(pos.Add(rLn), pos.Add(rLn.Neg()), src, clr)
	}
}

func line(a, b Vec2, src *image.RGBA, clr color.Color) {
	// The line
	var road = b.Add(a.Neg())
	inc := 1 / road.Mag()
	for i := 0.0; i < 1; i += inc {
		pos := a.Add(road.Mult(i))
		src.Set(int(pos.X), int(pos.Y), clr)
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
				cPos := centr.Add(Vec2{x, y}).Add(z)
				pClr := src.At(int(cPos.X), int(cPos.Y)).(color.RGBA)
				if z.Add(Vec2{x, y}).MagSq() < rSq {
					cClr = addClr(cClr, multClr(clr, 0.25))
				} else {
					cClr = addClr(cClr, multClr(pClr, 0.25))
				}
			}
			src.SetRGBA(int(centr.X+x), int(centr.Y+y), cClr)
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
			nPts[i] = pts[i+1].Add(pts[i].Neg()).Mult(prog).Add(pts[i])
		}
		pts = nPts
	}
	return pts[0]
}

//Rect Func to draw rects
func Rect(a, b Vec2, src *image.RGBA, clr color.Color) {
	for i := int(a.X); i < int(b.X); i++ {
		for j := int(a.Y); j < int(b.Y); j++ {
			src.Set(i, j, clr)
		}
	}
}
