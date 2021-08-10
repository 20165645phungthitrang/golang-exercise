package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
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
	return val
}

func main() {
	fmt.Println("Result from func:", Sqrt(2))
	fmt.Println("Result from math.Sqrt:", math.Sqrt(2))
}