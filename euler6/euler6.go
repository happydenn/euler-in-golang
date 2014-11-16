package main

import (
	"fmt"
	"math"
)

func main() {
	sumsquares := 0
	squaresums := 0

	for i := range make([]int, 100) {
		sumsquares += int(math.Pow(float64(i+1), 2))
		squaresums += i + 1
	}

	squaresums = int(math.Pow(float64(squaresums), 2))
	fmt.Println(squaresums - sumsquares)
}
