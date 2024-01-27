package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x

	var prev float64

	for i := 0; i < 10; i++ {
		z -= ((z * z) - x) / (2 * z)

		if prev == z {
			break
		}

		prev = z
	}

	return z
}

func main() {
	num := 2.0
	fmt.Println("calculated by sqrt:", Sqrt(num))
	fmt.Println("calculated by math.Sqrt:", math.Sqrt(num))
}
