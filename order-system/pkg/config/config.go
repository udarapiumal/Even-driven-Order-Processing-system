package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	RedisAddr   string
}

func Load() Config {
	return Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgresql://localhost:5432/OrderProcessing"),
		RedisAddr:   getEnv("REDIS_ADDR", "localhost:6379"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
