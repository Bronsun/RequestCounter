package handlers

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/Bronsun/RequestCounter/api/counter"
	"github.com/Bronsun/RequestCounter/api/hostname"
	"github.com/Bronsun/RequestCounter/db"
)

var mutex = &sync.Mutex{}

// RequestHanlder logic for main endpoint
// It counts all requests on the instances and save total number of requests to cluster in Redis store
// Return of the logic is "text/plain" response with information about total number of requests to the instance and cluster
func RequestHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	store := db.RedisConnect()

	hosts, err := store.Get("hostname").Result()
	if err != nil {
		fmt.Println(err)
	}

	host := hostname.OverwriteHostname(hosts)

	counter.IncrementRequests()

	errCounter := store.Incr("clustercount").Err()
	if errCounter != nil {
		fmt.Println(errCounter)
	}

	clustercount, err := store.Get("clustercount").Result()
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "text/plain")
	_, err = fmt.Fprintf(w, "You are talking to instance %s%s.\nThis is request %d to this instance and request %s to the cluster.\n", host, os.Getenv("WEB_PORT"), counter.GetRequests(), clustercount)

	mutex.Unlock()
}
