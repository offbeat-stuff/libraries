package main

import "fmt"

func main() {
	fmt.Println("Testing File")
	nuNet := createNet(3,[]int{2,2})
	for i:=range nuNet{
		for j:=range nuNet[i]{
			fmt.Println(i,j,">>",nuNet[i][j])
		}
	}
}
