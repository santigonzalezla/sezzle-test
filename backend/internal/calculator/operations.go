package calculator

import (
	"errors"
	"fmt"
	"math"
)

var (
	ErrDivisionByZero = errors.New("division by zero")
	ErrNegativeSqrt   = errors.New("cannot calculate negative square root")
)

func Add(a, b float64) (float64, error) {
	return a + b, nil
}

func Subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func Multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

func Power(a, b float64) (float64, error) {
	result := math.Pow(a, b)

	if math.IsNaN(result) || math.IsInf(result, 0) {
		return 0, fmt.Errorf("%v to the power of %v is not a finite number", a, b)
	}

	return result, nil
}

func Sqrt(a, _ float64) (float64, error) {
	if a < 0 {
		return 0, ErrNegativeSqrt
	}

	return math.Sqrt(a), nil
}

func Percentage(a, b float64) (float64, error) {
	return (a / 100) * b, nil
}
