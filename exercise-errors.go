package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %d\n", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := x / 2
	val := x
	for i := 0; i < 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
		if val-z < 1e-10 {
			break
		}
		val = z
	}
	return val, nil

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
