package main

import (
	"codingChallenge/controllers"
	"log"
	"net/http"
)

func main() {
	mux := controllers.NewRouter()
	log.Println("Server is starting on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
