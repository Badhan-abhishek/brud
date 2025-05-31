package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
}

var config *Config

func Load() *Config {
	if config != nil {
		return config
	}

	_ = godotenv.Load()

	dbUrl := getEnv("DB_URL", "")

	if dbUrl == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	config = &Config{
		DBUrl: dbUrl,
	}

	return config
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	log.Printf("Environment variable %s not set, using fallback: %s", key, fallback)
	return fallback
}
