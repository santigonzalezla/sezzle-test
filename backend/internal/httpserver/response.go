package httpserver

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/santigonzalezla/sezzle-calculator/internal/apperror"
)

type ErrorDetail struct {
	Code      string    `json:"code"`
	Message   string    `json:"message"`
	Path      string    `json:"path"`
	Timestamp time.Time `json:"timestamp"`
}

type ErrorBody struct {
	Error ErrorDetail `json:"error"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("failed to encode JSON response", "error", err)
	}
}

func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	appErr, ok := err.(*apperror.AppError)

	if !ok {
		appErr = apperror.Internal(err.Error())
	}

	WriteJSON(w, appErr.Status, ErrorBody{
		Error: ErrorDetail{
			Code:      appErr.Code,
			Message:   appErr.Message,
			Path:      r.URL.Path,
			Timestamp: time.Now().UTC(),
		},
	})
}
