//Package calculator contains pure arithmetic logic
//decoupled from any transport layer (HTTP, CLI, etc)
//Keeping it pure makes it trivial to unit test and reuse

package calculator

import (
	"errors"
	"math"
)

// Like sentinel errors let callers distinguish between
// different error conditions and handle them appropriately
var (
	ErrDivisionByZero    = errors.New("division by zero")
	ErrNegativeSqrt      = errors.New("cannot take square root of a negative number")
	ErrInvalidPercentage = errors.New("percentage base cannot be zero")
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

func Power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, ErrNegativeSqrt
	}
	return math.Sqrt(a), nil
}

// Percentage calculates the percentage of a base value
func Percentage(part, base float64) (float64, error) {
	if base == 0 {
		return 0, ErrInvalidPercentage
	}
	return (part / base) * 100, nil
}
