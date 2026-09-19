package httpserver

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/santigonzalezla/sezzle-calculator/internal/apperror"
)

type Middleware func(handler http.Handler) http.Handler

type statusWriter struct {
	http.ResponseWriter
	status int
}

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				slog.Error("recovered from panic", "error", rec, "path", request.URL.Path)
				WriteError(writer, request, apperror.Internal("internal server error"))
			}
		}()
		next.ServeHTTP(writer, request)
	})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		stWriter := &statusWriter{ResponseWriter: writer, status: http.StatusOK}

		next.ServeHTTP(stWriter, request)

		slog.Info("request",
			"method", request.Method,
			"path", request.URL.Path,
			"status", stWriter.status,
			"duration", time.Since(start),
		)
	})
}

func (stWriter *statusWriter) WriteHeader(status int) {
	stWriter.status = status
	stWriter.ResponseWriter.WriteHeader(status)
}

func CORS(allowedOrigins []string) Middleware {
	allowed := make(map[string]bool, len(allowedOrigins))

	for _, origin := range allowedOrigins {
		allowed[origin] = true
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			origin := request.Header.Get("Origin")

			if allowed[origin] {
				writer.Header().Set("Access-Control-Allow-Origin", origin)
				writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			}

			if request.Method == http.MethodOptions {
				writer.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(writer, request)
		})
	}
}

func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	return handler
}
