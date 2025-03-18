package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	ctx := context.Background()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Howareyou&143",
		DB:       2,
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis (DB 2)")
}

// Publish message to Redis channel
func PublishMessage(channel string, message string) error {
	ctx := context.Background()
	return RedisClient.Publish(ctx, channel, message).Err()
}

// Subscribe to a Redis channel
func SubscribeToChannel(channel string) *redis.PubSub {
	ctx := context.Background()
	return RedisClient.Subscribe(ctx, channel)
}
