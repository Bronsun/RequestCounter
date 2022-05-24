package db

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var client *redis.Client

func RedisConnect() redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DB_ADDR"),
		Password: os.Getenv("REDIS_DB_PASSWORD"),
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return *client
}
