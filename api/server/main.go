package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Bronsun/RequestCounter/api/handlers"
	"github.com/Bronsun/RequestCounter/api/repository"
	"github.com/Bronsun/RequestCounter/db"
)

func main() {

	db, err := db.RedisConnect()

	if err != nil {
		fmt.Println(err)
	}

	requestRepo := repository.NewRequestRepo(db)
	h := handlers.NewRequestHandler(requestRepo)

	s := http.Server{Addr: os.Getenv("WEB_PORT")}
	http.HandleFunc("/", h.RequestCounter)
	log.Fatal(s.ListenAndServe())

}
