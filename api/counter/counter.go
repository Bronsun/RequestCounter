package counter

import "sync/atomic"

var requests int64 = 0

// IncrementRequests increments the number of requests and returns the new value
func IncrementRequests() int64 {
	return atomic.AddInt64(&requests, 1)
}
