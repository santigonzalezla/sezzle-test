package calculator

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	tests := []struct {
		name       string
		handler    http.HandlerFunc
		body       string
		wantStatus int
		wantResult float64
	}{
		{name: "add", handler: HandleAdd, body: `{"a":2,"b":3}`, wantStatus: http.StatusOK, wantResult: 5},
		{name: "subtract", handler: HandleSubtract, body: `{"a":10,"b":4}`, wantStatus: http.StatusOK, wantResult: 6},
		{name: "sqrt", handler: HandleSqrt, body: `{"a":9}`, wantStatus: http.StatusOK, wantResult: 3},
		{name: "divide by zero", handler: HandleDivide, body: `{"a":10,"b":0}`, wantStatus: http.StatusBadRequest},
		{name: "sqrt negative", handler: HandleSqrt, body: `{"a":-4}`, wantStatus: http.StatusBadRequest},
		{name: "invalid json", handler: HandleAdd, body: `not json`, wantStatus: http.StatusBadRequest},
		{name: "multiply", handler: HandleMultiply, body: `{"a":3,"b":4}`, wantStatus: http.StatusOK, wantResult: 12},
		{name: "divide", handler: HandleDivide, body: `{"a":10,"b":2}`, wantStatus: http.StatusOK, wantResult: 5},
		{name: "power", handler: HandlePower, body: `{"a":2,"b":10}`, wantStatus: http.StatusOK, wantResult: 1024},
		{name: "percentage", handler: HandlePercentage, body: `{"a":50,"b":200}`, wantStatus: http.StatusOK, wantResult: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/calculate/x", bytes.NewBufferString(tt.body))
			rec := httptest.NewRecorder()

			tt.handler(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d, body = %s", rec.Code, tt.wantStatus, rec.Body.String())
			}

			if tt.wantStatus != http.StatusOK {
				return
			}

			var resp CalculateResponse
			if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if resp.Result != tt.wantResult {
				t.Errorf("result = %v, want %v", resp.Result, tt.wantResult)
			}
		})
	}
}
