package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr     string
	RedisPassword string
	JWTSecret     string
}

func LoadConfig() *Config {
	// Load .env using godotenv
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	return &Config{
		RedisAddr:     redisAddr,
		RedisPassword: redisPassword,
		JWTSecret:     jwtSecret,
	}
}
