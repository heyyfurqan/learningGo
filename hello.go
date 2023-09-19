package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var j int = 10

func split(sum int) (a, b int) {
	a = sum * 5 / 8
	b = sum - a
	return
}

func add(a, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(math.Pi)

	fmt.Println(add(34, 66))

	c, d := swap("hello", "nigga")
	fmt.Println(c, d)

	fmt.Println(split(25))

	cpp, js, react := true, false, "slow AF!"
	fmt.Println(j, cpp, js, react)

	var z complex128 = cmplx.Sqrt(-5 + 12i)

	fmt.Printf("Type: %T, Value: %v", z, z)

	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
