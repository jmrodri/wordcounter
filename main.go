package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Listening on http://localhost:9090")
	log.Fatal(http.ListenAndServe(":9090", NewHandler()))
}
