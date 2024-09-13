package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

// InitializeRedis sets up the Redis client connection
func InitializeRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		DB:       Config.RedisDB,
		Addr:     Config.RedisHost + ":" + Config.RedisPort,
		Password: Config.RedisPassword,
	})

	// Test the connection
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return
	}

	fmt.Println("Connected to Redis")
}
