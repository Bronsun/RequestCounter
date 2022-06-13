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

var mutex sync.Mutex

// RequestHanlder logic for main endpoint
// It counts all requests on the instances and save total number of requests to cluster in Redis store
// Return of the logic is "text/plain" response with information about total number of requests to the instance and cluster
func RequestHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	store, err := db.RedisConnect()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	hosts, err := store.Get(db.HostnameKey).Result()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	host, err := hostname.OverwriteHostname(hosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	counter.IncrementRequests()

	clustercounter, err := db.SaveRequests(db.ClustercountKey, store)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, err = fmt.Fprintf(w, "You are talking to instance %s%s.\nThis is request %d to this instance and request %d to the cluster.\n", host, os.Getenv("WEB_PORT"), counter.GetRequests(), clustercounter)

	mutex.Unlock()
}
