package config

import "os"

type Config struct {
	AppPort string
}

func Load() *Config {
	return &Config{
		AppPort: getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
