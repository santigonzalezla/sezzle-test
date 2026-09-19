package calculator

import "fmt"

type calcFunc func(a, b float64) (float64, error)

var operations = map[Operation]calcFunc{
	OpAdd:        Add,
	OpSubtract:   Subtract,
	OpMultiply:   Multiply,
	OpDivide:     Divide,
	OpPower:      Power,
	OpSqrt:       Sqrt,
	OpPercentage: Percentage,
}

func Calculate(op Operation, a, b float64) (float64, error) {
	fn, ok := operations[op]

	if !ok {
		return 0, fmt.Errorf("unsuporte operation: %q", op)
	}

	return fn(a, b)
}
