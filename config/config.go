package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr     string
	RedisPassword string
	JWTSecret     string
	QDrantURI     string
	QDrantAPIKey  string
	QDrantPort    int
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
	qDrantURI := os.Getenv("QDRANT_URI")
	qDrantAPIKey := os.Getenv("QDRANT_API_KEY")
	qDrantPort := os.Getenv("QDRANT_PORT")

	qDrantPortInt, err := strconv.Atoi(qDrantPort)
	if err != nil {
		panic("Invalid QDRANT_PORT value")
	}

	return &Config{
		RedisAddr:     redisAddr,
		RedisPassword: redisPassword,
		JWTSecret:     jwtSecret,
		QDrantURI:     qDrantURI,
		QDrantAPIKey:  qDrantAPIKey,
		QDrantPort:    qDrantPortInt,
	}
}
