package calculator

import (
	"errors"
	"testing"
)

func TestOperations(t *testing.T) {
	tests := []struct {
		name    string
		fn      func(a, b float64) (float64, error)
		a, b    float64
		want    float64
		wantErr error
	}{
		{name: "add positives", fn: Add, a: 2, b: 3, want: 5},
		{name: "add negatives", fn: Add, a: -2, b: -3, want: -5},
		{name: "subtract", fn: Subtract, a: 10, b: 4, want: 6},
		{name: "multiply", fn: Multiply, a: 3, b: 4, want: 12},
		{name: "multiply by zero", fn: Multiply, a: 3, b: 0, want: 0},
		{name: "divide", fn: Divide, a: 10, b: 2, want: 5},
		{name: "divide by zero", fn: Divide, a: 10, b: 0, wantErr: ErrDivisionByZero},
		{name: "power", fn: Power, a: 2, b: 10, want: 1024},
		{name: "sqrt", fn: Sqrt, a: 9, want: 3},
		{name: "sqrt of negative", fn: Sqrt, a: -4, wantErr: ErrNegativeSqrt},
		{name: "percentage", fn: Percentage, a: 50, b: 200, want: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fn(tt.a, tt.b)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("got error %v, want %v", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tt.want {
				t.Errorf("%s(%v, %v) = %v, want %v", tt.name, tt.a, tt.b, got, tt.want)
			}

		})
	}
}
