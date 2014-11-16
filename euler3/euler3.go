package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestPrimeFactor(600851475143))
}

func largestPrimeFactor(num int) int {
	for i := int(math.Floor(math.Sqrt(float64(num)))); i > 1; i-- {
		if math.Mod(float64(num), float64(i)) == 0 && isPrime(i) {
			return i
		}
	}

	return -1
}

func isPrime(num int) bool {
	prime := true

	for i := 2; i < num; i++ {
		switch i {
		case 0, 1:
			continue
		default:
			if math.Mod(float64(num), float64(i)) == 0 {
				prime = false
				break
			}
		}
	}

	return prime
}
