package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(faster(20))
}

func naive(n int) int {
	i := 1

	for {
		x := n * i
		found := true

		for k := range make([]int, 20) {
			if math.Mod(float64(x), float64(k+1)) != 0 {
				found = false
				break
			}
		}

		if found {
			return x
		}

		i++
	}
}

func faster(n int) int {
	if n == 1 {
		return 1
	} else {
		return lcm(n, faster(n-1))
	}
}

func gcd(i, j int) int {
	if j == 0 {
		return i
	} else {
		return gcd(j, int(math.Mod(float64(i), float64(j))))
	}
}

func lcm(i, j int) int {
	return i * j / gcd(i, j)
}
