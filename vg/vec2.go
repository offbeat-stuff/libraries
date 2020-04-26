package vg

import "math"

//Vec2 Vector 2d
type Vec2 struct {
	X, Y float64
}

func (v Vec2) equals(v2 Vec2) bool {
	return v.X == v2.X && v.Y == v2.Y
}

func (v Vec2) add(v2 Vec2) Vec2 {
	return Vec2{v.X + v2.X, v.Y + v2.Y}
}

func (v Vec2) neg() Vec2 {
	return Vec2{-v.X, -v.Y}
}

func (v Vec2) copy() Vec2 {
	return Vec2{v.X, v.Y}
}

func (v Vec2) mult(m float64) Vec2 {
	return Vec2{v.X * m, v.Y * m}
}

func (v Vec2) rotate(ang float64) Vec2 {
	x := (v.X * math.Cos(ang)) - (v.Y * math.Sin(ang))
	y := (v.X * math.Sin(ang)) + (v.Y * math.Cos(ang))
	return Vec2{x, y}
}

func (v Vec2) dot(v2 Vec2) float64 {
	return (v.X * v2.X) + (v.Y * v2.Y)
}

func (v Vec2) cross(v2 Vec2) float64 {
	return (v.X * v2.Y) - (v2.X * v.Y)
}

func (v Vec2) mag() float64 {
	return math.Sqrt(v.dot(v))
}

func (v Vec2) dir() float64 {
	return math.Atan2(v.Y, v.X)
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
