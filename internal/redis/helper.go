package redis

import (
	"github.com/redis/go-redis/v9"
	"os"
	"fmt"
	"context"
)

func verifyConnection() error {
	if redisClient == nil {
		return fmt.Errorf("Redis client is not initialized")
	}
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
