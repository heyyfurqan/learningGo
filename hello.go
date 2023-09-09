package main

import (
	"fmt"
	"math"
)

func add(a, b int) int {
	return a + b
}

func main() {
	//fmt.Printf("Hello, World!\n")
	fmt.Printf("Hello, You have %g problems.\n", math.Sqrt(7))
	//fmt.Println("Pie's value is ", math.Pi)
	fmt.Println(add(13, 12))
}
