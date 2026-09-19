package calculator

type Operation string

const (
	OpAdd        Operation = "add"
	OpSubtract   Operation = "subtract"
	OpMultiply   Operation = "multiply"
	OpDivide     Operation = "divide"
	OpPower      Operation = "power"
	OpSqrt       Operation = "sqrt"
	OpPercentage Operation = "percentage"
)

type CalculateRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type CalculateResponse struct {
	Operation Operation `json:"operation"`
	A         float64   `json:"a"`
	B         float64   `json:"b,omitempty"`
	Result    float64   `json:"result"`
}
