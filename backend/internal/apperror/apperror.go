package apperror

import "net/http"

type AppError struct {
	Code    string
	Message string
	Status  int
}

func (e *AppError) Error() string {
	return e.Message
}

func BadRequest(code, message string) *AppError {
	return &AppError{Code: code, Message: message, Status: http.StatusBadRequest}
}

func NotFound(code, message string) *AppError {
	return &AppError{Code: code, Message: message, Status: http.StatusNotFound}
}

func Internal(message string) *AppError {
	return &AppError{Code: "INTERNAL_ERROR", Message: message, Status: http.StatusInternalServerError}
}
