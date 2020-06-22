package vg

import "math"

//Vec2 Vector 2d
type Vec2 struct {
	X, Y float64
}

func (v Vec2) Equals(v2 Vec2) bool {
	return v.X == v2.X && v.Y == v2.Y
}

func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v.X + v2.X, v.Y + v2.Y}
}

func (v Vec2) Neg() Vec2 {
	return Vec2{-v.X, -v.Y}
}

func (v Vec2) Copy() Vec2 {
	return Vec2{v.X, v.Y}
}

func (v Vec2) Mult(m float64) Vec2 {
	return Vec2{v.X * m, v.Y * m}
}

func (v Vec2) Rotate(ang float64) Vec2 {
	x := (v.X * math.Cos(ang)) - (v.Y * math.Sin(ang))
	y := (v.X * math.Sin(ang)) + (v.Y * math.Cos(ang))
	return Vec2{x, y}
}

func (v Vec2) Dot(v2 Vec2) float64 {
	return (v.X * v2.X) + (v.Y * v2.Y)
}

func (v Vec2) Cross(v2 Vec2) float64 {
	return (v.X * v2.Y) - (v2.X * v.Y)
}

func (v Vec2) Mag() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vec2) Dir() float64 {
	return math.Atan2(v.Y, v.X)
}

func (v Vec2) Norm() Vec2 {
	m := v.Mag()
	if m == 0 {
		return Vec2{}
	}
	return v.Mult(1 / m)
}

func (v Vec2) MagSq() float64 {
	return v.Dot(v)
}

func sqr(x float64) float64 {
	return x * x
}
