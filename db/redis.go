package db

import (
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var Client *redis.Client

var (
	addr     = os.Getenv("REDIS_DB_ADDR")
	password = os.Getenv("REDIS_DB_PASSWORD")
	db, _    = strconv.Atoi(os.Getenv("REDIS_DB"))
)

const (
	ClustercountKey = "clustercount"
	HostnameKey     = "hostname"
)

// RedisConnect connect to Redis database
func RedisConnect() (*redis.Client, error) {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := Client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return Client, nil
}

// SaveRequests counts and save number of total requests in claster.
// It uses redis Piplined and Transaction to prevenet from race condition
func SaveRequests(key string, store *redis.Client) (int64, error) {

	var clustercount int64

	err := store.Watch(func(tx *redis.Tx) error {
		val, err := tx.Get(key).Int64()
		if err != nil && err != redis.Nil {
			return err
		}

		clustercount = val

		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			pipe.Set(key, strconv.FormatInt(clustercount+1, 10), 0)
			return nil
		})
		return err

	}, key)

	if err == redis.TxFailedErr {
		return SaveRequests(key, store)
	}

	return clustercount + 1, err
}
