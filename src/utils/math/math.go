package math

import (
	"errors"
)

func SquareRoot(num float64) (float64, error) {
	const ITERATIONS int = 10

	if num < 0 {
		return 0, errors.New("Cannot take square root of a negative number")
	}
	estimate := num / 2
	for i := 0; i < ITERATIONS; i++ {
		div := num / estimate
		estimate = (estimate + div) / 2
	}
	return estimate, nil
}

func Factorial(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("Cannot take factorial of a negative number")
	}
	if num <= 1 {
		return num, nil
	}
	prev, _ := Factorial(num - 1)
	return num * prev, nil
}

func GCD(n1 int, n2 int) int {
	if(n2 < n1) {
		return GCD(n2, n1)
	}
	if n1 == n2 {
		return n1
	}
	return GCD(n1, n2 - n1)
}