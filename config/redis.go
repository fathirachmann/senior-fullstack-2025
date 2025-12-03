package config

import "github.com/redis/go-redis/v9"

func NewRedisConfig(addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	return rdb
}
