package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DatabaseURL string
	Port        string
	WalletURL   string
}

func LoadConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		panic("No .env file found")
	}
	return &AppConfig{
		DatabaseURL: getEnvString("DATABASE_URL", "root:your_password@tcp(localhost:3306)/test1-go?charset=utf8mb4&parseTime=True&loc=Local"),
		Port:        getEnvString("PORT", ":8080"),
		WalletURL:   getEnvString("WALLET_URL", "http://localhost:8081"),
	}
}

func getEnvString(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
