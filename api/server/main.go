package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Bronsun/RequestCounter/api/handlers"
	"github.com/Bronsun/RequestCounter/db"
)

func main() {

	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	store := db.RedisConnect()
	store.Append("hostname", host+",")

	s := http.Server{Addr: os.Getenv("WEB_PORT")}

	http.HandleFunc("/", handlers.RequestHandler)

	log.Fatal(s.ListenAndServe())

}
