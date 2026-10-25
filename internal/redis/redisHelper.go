package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func verifyConnection() error {
	ctx := context.Background()
	rdb := redisClient

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	fmt.Println("Connected to Redis!")

	// Store a value for 10 minutes.
	err := rdb.Set(ctx, "user:1:name", "Sushant", 1*time.Minute).Err()
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the value.
	name, err := rdb.Get(ctx, "user:1:name").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User name:", name)
	return nil
}

func getConnection() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     getConnectionString(),
		Password: getPassword(),
		DB:       0,
	})

	ctx := context.Background()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	err := verifyConnection()
	if err != nil {
		return err
	}

	return nil
}

func getConnectionString() string {
	address := os.Getenv("REDIS_ADDR")
	if address == "" {
		return "localhost:6379"
	}
	return address
}

func getPassword() string {
	password := os.Getenv("REDIS_PASSWORD")
	if password == "" {
		return ""
	}
	return password
}
