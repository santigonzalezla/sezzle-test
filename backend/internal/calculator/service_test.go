package calculator

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name    string
		op      Operation
		a, b    float64
		want    float64
		wantErr bool
	}{
		{name: "know operation", op: OpAdd, a: 2, b: 3, want: 5},
		{name: "unknown operation", op: "modulo", a: 2, b: 3, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.op, tt.a, tt.b)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected an error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tt.want {
				t.Errorf("Calculate(%v, %v, %v) = %v, want %v", tt.op, tt.a, tt.b, got, tt.want)
			}
		})
	}
}
