package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DBHost       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBPort       string
	RedisAddr    string
	RedisPass    string
	TMDBApiKey   string
	OMDBApiKey   string
	JWTSecret    string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Note: .env file not found, using system environment variables")
	}

	return &Config{
		Port:         getEnv("PORT", "8080"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", "postgres"),
		DBName:       getEnv("DB_NAME", "movie_watchlist"),
		DBPort:       getEnv("DB_PORT", "5432"),
		RedisAddr:    getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPass:    getEnv("REDIS_PASSWORD", ""),
		TMDBApiKey:   getEnv("TMDB_API_KEY", ""),
		OMDBApiKey:   getEnv("OMDB_API_KEY", ""),
		JWTSecret:    getEnv("JWT_SECRET", "supersecretkey"),
	}
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
