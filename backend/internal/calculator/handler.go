package calculator

import (
	"encoding/json"
	"net/http"

	"github.com/santigonzalezla/sezzle-calculator/internal/apperror"
	"github.com/santigonzalezla/sezzle-calculator/internal/httpserver"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/calculate/add", HandleAdd)
	mux.HandleFunc("POST /api/calculate/subtract", HandleSubtract)
	mux.HandleFunc("POST /api/calculate/multiply", HandleMultiply)
	mux.HandleFunc("POST /api/calculate/divide", HandleDivide)
	mux.HandleFunc("POST /api/calculate/power", HandlePower)
	mux.HandleFunc("POST /api/calculate/sqrt", HandleSqrt)
	mux.HandleFunc("POST /api/calculate/percentage", HandlePercentage)
}

// HandleAdd godoc
// @Summary Add two numbers
// @Description Returns a + b
// @Tags calculator
// @Accept json
// @Produce json
// @Param body body CalculateRequest true "Operands"
// @Success 200 {object} CalculateResponse
// @Failure 400 {object} httpserver.ErrorBody
// @Router /api/calculate/add [post]
func HandleAdd(writer http.ResponseWriter, request *http.Request) {
	handleOperation(OpAdd)(writer, request)
}

// HandleSubtract godoc
// @Summary Subtract two numbers
// @Description  Returns a - b
// @Tags calculator
// @Accept json
// @Produce json
// @Param body body CalculateRequest true "Operands"
// @Success 200 {object} CalculateResponse
// @Failure 400 {object} httpserver.ErrorBody
// @Router /api/calculate/subtract [post]
func HandleSubtract(w http.ResponseWriter, r *http.Request) {
	handleOperation(OpSubtract)(w, r)
}

// HandleMultiply godoc
// @Summary Multiply two numbers
// @Description  Returns a * b
// @Tags calculator
// @Accept json
// @Produce json
// @Param body body CalculateRequest true "Operands"
// @Success 200 {object} CalculateResponse
// @Failure 400 {object} httpserver.ErrorBody
// @Router /api/calculate/multiply [post]
func HandleMultiply(w http.ResponseWriter, r *http.Request) {
	handleOperation(OpMultiply)(w, r)
}

// HandleDivide godoc
// @Summary Divide two numbers
// @Description  Returns a / b. Fails with 400 if b is 0.
// @Tags calculator
// @Accept json
// @Produce json
// @Param body  body      CalculateRequest  true  "Operands"
// @Success 200   {object}  CalculateResponse
// @Failure 400   {object}  httpserver.ErrorBody  "b must not be zero"
// @Router /api/calculate/divide [post]
func HandleDivide(w http.ResponseWriter, r *http.Request) {
	handleOperation(OpDivide)(w, r)
}

// HandlePower godoc
// @Summary Raise a to the power of b
// @Description  Returns a^b. Fails with 400 if the result is not finite.
// @Tags calculator
// @Accept json
// @Produce json
// @Param body  body      CalculateRequest  true  "Base (a) and exponent (b)"
// @Success 200   {object}  CalculateResponse
// @Failure 400   {object}  httpserver.ErrorBody
// @Router /api/calculate/power [post]
func HandlePower(w http.ResponseWriter, r *http.Request) {
	handleOperation(OpPower)(w, r)
}

// HandleSqrt godoc
// @Summary Square root of a number
// @Description  Returns sqrt(a). B is ignored. Fails with 400 if a is negative.
// @Tags calculator
// @Accept json
// @Produce json
// @Param body  body      CalculateRequest  true  "Only a is used"
// @Success 200   {object}  CalculateResponse
// @Failure 400   {object}  httpserver.ErrorBody  "a must not be negative"
// @Router /api/calculate/sqrt [post]
func HandleSqrt(w http.ResponseWriter, r *http.Request) {
	handleOperation(OpSqrt)(w, r)
}

// HandlePercentage godoc
// @Summary Percentage
// @Description  Returns a percent of b (e.g. a=50, b=200 -> 100)
// @Tags calculator
// @Accept json
// @Produce json
// @Param body  body CalculateRequest  true  "Operands"
// @Success 200 {object}  CalculateResponse
// @Failure 400 {object}  httpserver.ErrorBody
// @Router /api/calculate/percentage [post]
func HandlePercentage(w http.ResponseWriter, r *http.Request) {
	handleOperation(OpPercentage)(w, r)
}

func handleOperation(op Operation) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var req CalculateRequest
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			httpserver.WriteError(writer, request, apperror.BadRequest("INVALID_BODY", "request body must be a valid JSON"))
			return
		}

		result, err := Calculate(op, req.A, req.B)

		if err != nil {
			httpserver.WriteError(writer, request, apperror.BadRequest("CALCULATION_ERROR", err.Error()))
			return
		}

		httpserver.WriteJSON(writer, http.StatusOK, CalculateResponse{
			Operation: op,
			A:         req.A,
			B:         req.B,
			Result:    result,
		})
	}

}
