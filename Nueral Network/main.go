package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	tIns := [][]float64{{1, 0, 0, 1}, {0, 1, 1, 0}, {1, 1, 0, 0}, {0, 0, 1, 1}}
	tOps := [][]float64{{1}, {1}, {0}, {0}}
	testInps := [][]float64{{1, 1, 1, 1}}
	// Create a nueral net
	var myNet = createNet(4, []int{3, 3, 1})
	//myNet[1][0].actf = relu

	fmt.Println(">> Before Training")
	for i := range tIns {
		thought := myNet.think(tIns[i])
		fmt.Print("|>>	", thought, "<|>")
		for j := range thought {
			fmt.Print(nearTo(thought[j]), ",")
		}
		fmt.Println(tOps[i], "	<<|	")
	}
	fmt.Print("\n\n")

	fmt.Println(">>Begining Training>>")
	for x := int64(0); x < 200; x++ {
		for y := 0; y < 200; y++ {
			for i := range tIns {
				myNet.backProp(tIns[i], tOps[i])
			}
		}
	}
	fmt.Println()
	fmt.Println(">> After Training")
	fmt.Println("\n<|> What are the Answers")
	for i := range tIns {
		thought := myNet.think(tIns[i])
		fmt.Print("|>>	", thought, "<|>	")
		for j := range thought {
			fmt.Print(nearTo(thought[j]), ",")
		}
		fmt.Println(tOps[i], "	<<|	")
	}
	fmt.Println("\n<|> What are Answers to Test")
	for i := range testInps {
		thought := myNet.think(testInps[i])
		fmt.Print("|>>	", thought, "<|>	")
		for j := range thought {
			fmt.Print(nearTo(thought[j]), ",")
		}
		fmt.Println("	<<|	")
	}
}

func nearTo(x float64) float64 {
	if x < 0.5 {
		return 0
	}
	return 1
}

func calcCost(op []float64, dop []float64) float64 {
	var z float64
	for i := range op {
		z += costFunc(op[i], dop[i])
	}
	return z
}

func costFunc(op float64, dop float64) float64 {
	return sqr(op - dop)
}

func costFuncDer(op float64, dop float64) float64 {
	return (2*op - 2*dop)
}
