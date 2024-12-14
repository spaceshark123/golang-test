package math_test

import (
	"testing"
	"github.com/spaceshark123/golang-test/src/utils/math"
)

func TestSquareRoot(t *testing.T) {
	// use a table driven test
	var tests = []struct{
		name string
		input float64
		output float64
		hasError bool
	}{
		{"normal input 1", 4, 2, false},
		{"normal input 2", 36, 6, false},
		{"decimal input", 2.25, 1.5, false},
		{"input: 1", 1, 1, false},
		{"input: 0", 0, 0, false},
		{"negative integer input", -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := math.SquareRoot(tt.input)
			if ans != tt.output {
				t.Errorf("expected: %f, recieved: %f", tt.output, ans)
			}
			if (err != nil) != tt.hasError {
				t.Errorf("expected: %v, recieved: %v", tt.hasError, err)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	// use a table driven test
	var tests = []struct{
		name string
		input int
		output int
		hasError bool
	}{
		{"normal input 1", 4, 24, false},
		{"normal input 2", 5, 120, false},
		{"input: 1", 1, 1, false},
		{"input: 0", 0, 0, false},
		{"negative integer input", -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := math.Factorial(tt.input)
			if ans != tt.output {
				t.Errorf("expected: %d, recieved: %d", tt.output, ans)
			}
			if (err != nil) != tt.hasError {
				t.Errorf("expected: %v, recieved: %v", tt.hasError, err)
			}
		})
	}
}

func TestGCD(t *testing.T) {
	// use a table driven test
	var tests = []struct{
		name string
		input1 int
		input2 int
		output int
	}{
		{"normal input 1", 4, 8, 4},
		{"normal input 2", 5, 10, 5},
		{"input: 1", 1, 1, 1},
		{"input: 0", 0, 0, 0},
		{"input: 0 and 1", 0, 1, 1},
		{"input: 1 and 0", 1, 0, 1},
		{"1 negative integer input", -1, 1, 1},
		{"2 negative integer input", -3, -6, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := math.GCD(tt.input1, tt.input2)
			if ans != tt.output {
				t.Errorf("expected: %d, recieved: %d", tt.output, ans)
			}
		})
	}
}