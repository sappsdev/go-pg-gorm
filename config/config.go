package config

import (
	"os"
)

func Debug() bool {
	value := os.Getenv("GIN_MODE")
	if len(value) == 0 {
		return true
	}
	return false
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
