package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Port           string
	AllowedOrigins []string
}

func Load() (Config, error) {
	port := getEnv("PORT", "8080")
	if strings.TrimSpace(port) == "" {
		return Config{}, fmt.Errorf("PORT environment variable not set")
	}

	originsRaw := getEnv("CORS_ORIGINS", "http://localhost:5173")
	origins := strings.Split(originsRaw, ",")

	for i, origin := range origins {
		origins[i] = strings.TrimSpace(origin)
	}

	return Config{
		Port:           port,
		AllowedOrigins: origins,
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}

	return fallback
}
