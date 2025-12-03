package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr string
	JWTSecret string
}

func LoadConfig() *Config {
	// Load .env using godotenv
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	redisAddr := os.Getenv("REDIS_ADDR")

	return &Config{
		RedisAddr: redisAddr,
		JWTSecret: jwtSecret,
	}
}
