package db

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var Client *redis.Client

// RedisConnect connect to Redis database
func RedisConnect() redis.Client {
	Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DB_ADDR"),
		Password: os.Getenv("REDIS_DB_PASSWORD"),
		DB:       0,
	})
	pong, err := Client.Ping().Result()
	fmt.Println(pong, err)

	return *Client
}
