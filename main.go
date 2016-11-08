package main

import (
	"log"
	"net/http"
)

func main() {
	router := Router()
	log.Printf("Listening...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
