package main

import "math"

type perc struct {
	weighs []float64
	bias   float64
	actf   actFunc
}

type actFunc struct {
	fun func(float64) float64
	der func(float64) float64
}

// |> Normal Activation Functions
var (
	sigmoid actFunc = actFunc{
		fun: func(x float64) float64 {
			return 1 / (1 + math.Pow(math.E, -x))
		},
		der: func(x float64) float64 {
			return x * (1 - x)
		},
	}
	relu actFunc = actFunc{
		fun: func(x float64) float64 {
			if x > 0 {
				return x
			}
			return 0
		},
		der: func(x float64) float64 {
			if x > 0 {
				return 1
			}
			return 0
		},
	}
)

// |> Process the inputs to give output
func (n *perc) proc(inputs []float64) float64 {
	var sum = dotArr(inputs, n.weighs)
	return n.actf.fun(sum + n.bias)
}

func dotArr(x, y []float64) float64 {
	var z float64
	for i := range x {
		z += x[i] * y[i]
	}
	return z
}
