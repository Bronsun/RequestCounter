package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Bronsun/RequestCounter/api/handlers"
)

func main() {
	s := http.Server{Addr: os.Getenv("WEB_PORT")}

	http.HandleFunc("/", handlers.RequestHandler)

	log.Fatal(s.ListenAndServe())
}
