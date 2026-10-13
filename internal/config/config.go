package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

func LoadConfig() *DBConfig {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env file, using system defaults.")
	}

	return &DBConfig{
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASSWORD"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}
}

func Get(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
