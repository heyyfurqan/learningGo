package main

import (
	"fmt"
	"math"
)

func add(a, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(44, 55))
	fmt.Println(swap("hello", "nigga"))
}
