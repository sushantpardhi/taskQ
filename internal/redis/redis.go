package redis

import (
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func Connect() error {

	err := getConnection()
	if err != nil {
		return err
	}

	err = verifyConnection()
	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	if redisClient != nil {
		redisClient.Close()
	}
	return nil
}
