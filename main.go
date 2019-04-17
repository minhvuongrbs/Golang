package main

import (
	"log"
	"net/http"
	"welcomerobot-api/route"
)

func main() {
	const PORT string = ":8000"
	log.Printf("Server started at localhost" + PORT + "/wr/v1")
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(PORT, router))
}
