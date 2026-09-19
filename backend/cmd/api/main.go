package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/santigonzalezla/sezzle-calculator/internal/calculator"
	"github.com/santigonzalezla/sezzle-calculator/internal/config"
	"github.com/santigonzalezla/sezzle-calculator/internal/httpserver"
)

// @title Sezzle Calculator API
// @version 1.0
// @description REST API for a full-stack calculator. Exposes basic and advanced
// arithmetic operations (add, subtract, multiply, divide, power, sqrt,
// percentage).
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	cfg, err := config.Load()

	if err != nil {
		slog.Error("invalid configuration", "error", err)
		os.Exit(1)
	}

	mux := httpserver.New()
	calculator.RegisterRoutes(mux)
	handler := httpserver.Wrap(mux, cfg.AllowedOrigins)

	slog.Info("starting server", "port", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		slog.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
