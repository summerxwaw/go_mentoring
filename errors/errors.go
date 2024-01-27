package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {

	return fmt.Sprint("you got an error in your response: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := x

	var prev float64

	for i := 0; i < 10; i++ {
		z -= ((z * z) - x) / (2 * z)

		if prev == z {
			break
		}

		prev = z
	}

	if z < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	return z, nil
}

func main() {
	num := -40.0

	fmt.Println(Sqrt(num))
	fmt.Println(Sqrt(num * num))
}
