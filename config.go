package main

import (
	"os"
)

type Config struct {
	IsDev bool
}

func Load() Config {
	return Config{
		IsDev: os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "",
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
