package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(prime(10001))
}

func prime(n int) int {
	ans := 1

	for i := 0; i < n; i++ {
		ans++

		for {
			breakpoint := int(math.Floor(math.Sqrt(float64(ans))))
			isPrime := true

			for j := 2; j <= breakpoint; j++ {
				if ans%j == 0 {
					isPrime = false
					break
				}
			}

			if isPrime {
				break
			} else {
				ans++
			}
		}
	}

	return ans
}
