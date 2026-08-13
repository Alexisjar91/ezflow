package config

import "os"

type Config struct {
	DatabaseURL string
}

func Load() Config {
	return Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://admin:1234@localhost:5433/ezflow?sslmode=disable"),
	}
}
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
