package main

import (
	"log"
	"net/http"
	"welcome_robot/go"
)

func main() {
	const PORT string = ":8080"
	log.Printf("Server started at localhost" + PORT + "/wr/v1")
	router := _go.NewRouter()

	log.Fatal(http.ListenAndServe(PORT, router))

}