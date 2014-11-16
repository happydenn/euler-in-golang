package main

import (
	"fmt"
	"math"
)

func main() {
	n := 0
	sum := 0

	for {
		term := fib(n)

		if term > 4000000 {
			break
		}

		if math.Mod(float64(term), 2) == 0 {
			sum += term
		}
		n++
	}

	fmt.Println(sum)
}

func fib(i int) int {
	a := 0

	switch i {
	case 0:
		a = 0
	case 1:
		a = 1
	default:
		a = fib(i-1) + fib(i-2)
	}

	return a
}
