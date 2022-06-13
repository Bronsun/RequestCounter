package repository

import (
	"github.com/go-redis/redis"
)

// RequestRepo implements models.RequestRepo
type RequestRepo struct {
	db *redis.Client
}

// NewRequestRepo init request repo
func NewRequestRepo(db *redis.Client) *RequestRepo {
	return &RequestRepo{
		db: db,
	}
}

// SaveRequest counts and save number of total requests in claster.
// It uses redis Piplined and Transaction to prevenet from race condition
func (r *RequestRepo) SaveRequest(key string) (int64, error) {
	var clustercount int64

	err := r.db.Watch(func(tx *redis.Tx) error {
		val, err := tx.Get(key).Int64()
		if err != nil && err != redis.Nil {
			return err
		}

		clustercount = val

		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			pipe.Incr(key)
			return nil
		})
		return err

	}, key)

	if err == redis.TxFailedErr {
		return r.SaveRequest(key)
	}

	return clustercount + 1, nil
}
