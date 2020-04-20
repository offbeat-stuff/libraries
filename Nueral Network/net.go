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

func (n *net) think(input []float64) []float64 {
	for i := 0; i < len(*n); i++ {
		var out = []float64{}
		for j := 0; j < len((*n)[i]); j++ {
			out = append(out, (*n)[i][j].proc(input))
		}
		input = out
	}
	return input
}

func (n *net) backProp(input []float64, dOp []float64) {
	var ops [][]float64
	ops = append(ops, input)
	for i := 0; i < len(*n); i++ {
		ops = append(ops, []float64{})
		for j := 0; j < len((*n)[i]); j++ {
			ops[i+1] = append(ops[i+1], (*n)[i][j].proc(ops[i]))
		}
	}
	for i := range *n {
		for j := range (*n)[i] {
			for k := range (*n)[i][j].weighs {
				(*n)[i][j].weighs[k] -= n.calcWeighDer(i, j, k, ops, dOp)
			}
			(*n)[i][j].bias -= n.calcBiasDer(i, j, ops, dOp)
		}
	}
}

func (n net) calcWeighDer(x, y, z int, ops [][]float64, dOp []float64) float64 {
	adj := ops[x][z] * n[x][y].actf.der(ops[x+1][y])
	nudges := make([]float64, len(ops[x+1]))
	nudges[y] = adj
	x++
	for ; x < len(n); x++ {
		nudges = n.calLayerDer(nudges, ops[x+1], x)
	}
	adj = 0.0
	for i := range nudges {
		adj += nudges[i] * 2 * (ops[x][i] - dOp[i])
	}
	return adj
}

func (n net) calcBiasDer(x, y int, ops [][]float64, dOp []float64) float64 {
	adj := n[x][y].actf.der(ops[x+1][y])
	nudges := make([]float64, len(ops[x+1]))
	nudges[y] = adj
	x++
	for ; x < len(n); x++ {
		nudges = n.calLayerDer(nudges, ops[x+1], x)
	}
	adj = 0.0
	for i := range nudges {
		adj += nudges[i] * 2 * (ops[x][i] - dOp[i])
	}
	return adj
}

func (n net) calLayerDer(nudges, inps []float64, z int) []float64 {
	var ngOp = make([]float64, len(n[z]))
	for i := range n[z] {
		for j := range nudges {
			ngOp[i] += nudges[j] * n[z][i].weighs[j]
		}
		ngOp[i] *= n[z][i].actf.der(inps[i])
	}
	return ngOp
}
