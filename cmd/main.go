package main

import (
	"fmt"
	"taskQ/internal/redis"
)

func main() {
	if err := redis.Connect(); err != nil {
		fmt.Printf("Error connecting to Redis: %v\n", err)
		return
	}

	defer redis.Close()

}
