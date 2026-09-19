package httpserver

import (
	"net/http"

	"github.com/santigonzalezla/sezzle-calculator/internal/apperror"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/santigonzalezla/sezzle-calculator/docs"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handleHealth)
	mux.Handle("GET /docs/", httpSwagger.WrapHandler)
	mux.HandleFunc("/", handleNotFound)
	return mux
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	WriteJSON(writer, http.StatusOK, map[string]string{"status": "ok"})
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	WriteError(w, r, apperror.NotFound("NOT_FOUND", "resource not found"))
}

func Wrap(handler http.Handler, allowedOrigins []string) http.Handler {
	return Chain(handler, Recover, Logging, CORS(allowedOrigins))
}
