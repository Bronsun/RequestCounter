package db

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var (
	addr     = os.Getenv("REDIS_DB_ADDR")
	password = os.Getenv("REDIS_DB_PASSWORD")
	db, _    = strconv.Atoi(os.Getenv("REDIS_DB"))
)

// RedisConnect connect to Redis database
func RedisConnect() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
