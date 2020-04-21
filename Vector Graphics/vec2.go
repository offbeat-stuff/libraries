package main

import "math"

//vec2 Vector 2d
type vec2 struct {
	x, y float64
}

func (v vec2) equals(v2 vec2) bool {
	return v.x == v2.x && v.y == v2.y
}

func (v vec2) add(v2 vec2) vec2 {
	return vec2{v.x + v2.x, v.y + v2.y}
}

func (v vec2) neg() vec2 {
	return vec2{-v.x, -v.y}
}

func (v vec2) copy() vec2 {
	return vec2{v.x, v.y}
}

func (v vec2) mult(m float64) vec2 {
	return vec2{v.x * m, v.y * m}
}

func (v vec2) rotate(ang float64) vec2 {
	x := (v.x * math.Cos(ang)) - (v.y * math.Sin(ang))
	y := (v.x * math.Sin(ang)) + (v.y * math.Cos(ang))
	return vec2{x, y}
}

func (v vec2) dot(v2 vec2) float64 {
	return (v.x * v2.x) + (v.y * v2.y)
}

func (v vec2) cross(v2 vec2) float64 {
	return (v.x * v2.y) - (v2.x * v.y)
}

func (v vec2) mag() float64 {
	return math.Sqrt(v.dot(v))
}

func (v vec2) dir() float64 {
	return math.Atan2(v.y, v.x)
}

func (v vec2) normalize() vec2 {
	m := v.mag()
	if m == 0 {
		return vec2{}
	}
	return v.mult(1 / m)
}

func (v vec2) magSq() float64 {
	return v.dot(v)
}

func sqr(x float64) float64 {
	return x * x
}
