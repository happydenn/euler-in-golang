package main

import (
	"fmt"
	"math"
)

func main() {
	counter := make([]int, 1000)
	sum := 0

	for i := range counter {
		if math.Mod(float64(i), 3.0) == 0 || math.Mod(float64(i), 5.0) == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
