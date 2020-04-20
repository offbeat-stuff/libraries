package main

import "math/rand"

type net [][]perc

func sqr(x float64) float64 {
	return x * x
}

func randFloat() float64 {
	return rand.Float64()*2 - 1
}

func randArr(len int) []float64 {
	out := make([]float64, len)
	for i := range out {
		out[i] = randFloat()
	}
	return out
}

func createNet(inps int, layout []int) net {
	var oP net = make([][]perc, len(layout))
	for i := range oP {
		oP[i] = make([]perc, layout[i])
		for j := range oP[i] {
			var len int
			if i == 0 {
				len = inps
			} else {
				len = layout[i-1]
			}
			oP[i][j] = perc{
				randArr(len), randFloat(), sigmoid,
			}
		}
	}
	return oP
}
