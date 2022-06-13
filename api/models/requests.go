package models

// Key const type
const Key = "clustercount"

// RequestRepository interface for defining methods
type RequestRepository interface {
	SaveRequest(key string) (int64, error)
}
