package main

import (
	"fmt"
	"strconv"
)

func main() {
	largest := 0

	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			p := i * j
			p_str := strconv.Itoa(p)
			reverse_p := reverse(p_str)

			if p_str == reverse_p {
				if p > largest {
					largest = p
				}
			}
		}
	}

	fmt.Println(largest)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
