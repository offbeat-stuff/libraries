package vg

import "math"

//Vec2 Vector 2d
type Vec2 struct {
	x, y float64
}

func (v Vec2) equals(v2 Vec2) bool {
	return v.x == v2.x && v.y == v2.y
}

func (v Vec2) add(v2 Vec2) Vec2 {
	return Vec2{v.x + v2.x, v.y + v2.y}
}

func (v Vec2) neg() Vec2 {
	return Vec2{-v.x, -v.y}
}

func (v Vec2) copy() Vec2 {
	return Vec2{v.x, v.y}
}

func (v Vec2) mult(m float64) Vec2 {
	return Vec2{v.x * m, v.y * m}
}

func (v Vec2) rotate(ang float64) Vec2 {
	x := (v.x * math.Cos(ang)) - (v.y * math.Sin(ang))
	y := (v.x * math.Sin(ang)) + (v.y * math.Cos(ang))
	return Vec2{x, y}
}

func (v Vec2) dot(v2 Vec2) float64 {
	return (v.x * v2.x) + (v.y * v2.y)
}

func (v Vec2) cross(v2 Vec2) float64 {
	return (v.x * v2.y) - (v2.x * v.y)
}

func (v Vec2) mag() float64 {
	return math.Sqrt(v.dot(v))
}

func (v Vec2) dir() float64 {
	return math.Atan2(v.y, v.x)
}

func (v Vec2) normalize() Vec2 {
	m := v.mag()
	if m == 0 {
		return Vec2{}
	}
	return v.mult(1 / m)
}

func (v Vec2) magSq() float64 {
	return v.dot(v)
}

func sqr(x float64) float64 {
	return x * x
}
