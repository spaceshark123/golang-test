package math

import (
	"errors"
)

func SquareRoot(num float64) (float64, error) {
	const ITERATIONS int = 10

	if num < 0 {
		return 0, errors.New("Cannot take square root of a negative number")
	}
	if num == 0 {
		return 0, nil
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
	// in case of negative numbers, make them positive
	if(n1 < 0) {
		n1 = -n1
	}
	if(n2 < 0) {
		n2 = -n2
	}
	// recusively find the GCD
	if(n2 == 0) {
		return n1
	}
	return GCD(n2, n1 % n2)
}