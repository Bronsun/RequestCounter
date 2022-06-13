package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Bronsun/RequestCounter/api/counter"
	"github.com/Bronsun/RequestCounter/api/models"
)

// RequestHandler holds everything that controller needs
type RequestHandler struct {
	requestRepo models.RequestRepository
}

// NewReuqestHandler return a nen RequestHandler
func NewRequestHandler(requestRepo models.RequestRepository) *RequestHandler {
	return &RequestHandler{
		requestRepo: requestRepo,
	}
}

// RequestCounter logic for main endpoint
// It counts all requests on the instances and save total number of requests to cluster in Redis store
// Return of the logic is "text/plain" response with information about total number of requests to the instance and cluster
func (h *RequestHandler) RequestCounter(w http.ResponseWriter, r *http.Request) {

	host, err := os.Hostname()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	counter := counter.IncrementRequests()

	clustercounter, err := h.requestRepo.SaveRequest(models.Key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, err = fmt.Fprintf(w, "You are talking to instance %s%s.\nThis is request %d to this instance and request %d to the cluster.\n", host, os.Getenv("WEB_PORT"), counter, clustercounter)

}
