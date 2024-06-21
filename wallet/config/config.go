package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DatabaseURL string
	Port        string
}

func LoadConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		panic("No .env file found")
	}
	return &AppConfig{
		DatabaseURL: getEnvString("DATABASE_URL", "mongodb://localhost:27017"),
		Port:        getEnvString("PORT", "8080"),
	}
}

func getEnvString(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
