package config

import "os"

var AppPort = getEnv("APP_PORT", "3000")

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}
