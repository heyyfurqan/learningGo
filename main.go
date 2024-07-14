package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var j int = 10

func average(a, b, c int) float32 {
	return float32(a+b+c) / 3
}

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

	c, d := swap("hello", "dawg")
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

	var l int = 5
	var m float64 = float64(l)
	fmt.Println(l, m)

	quiz1, quiz2, quiz3 := 23, 34, 29

	if quiz1 > quiz2 {
		fmt.Println("Quiz 1 has higher marks.")
	} else if quiz2 > quiz1 {
		fmt.Println("Quiz 2 has higher marks.")
	} else {
		fmt.Println("Both are equal")
	}

	switch quiz3 {
	case 29:
		fmt.Println("Noice")
	case 30:
		fmt.Println("Woohoo")
	}

	switch age := 17; {
	case age == 0:
		fmt.Println("Newborn")
	case age >= 1 && age <= 3:
		fmt.Println("Toddler")
	case age >= 4 && age <= 12:
		fmt.Println("Child")
	case age >= 13 && age <= 17:
		fmt.Println("Teenager")
	case age >= 18:
		fmt.Println("Adult")
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Nigga ", i)
	}

	for i := 0; i <= 50; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
