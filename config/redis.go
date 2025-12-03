package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisConfig(addr, password string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Redis Connected")

	return rdb
}
