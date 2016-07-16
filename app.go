package main

import (
	"log"
	"net/http"
)

func main() {
	initApp()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8081", router))
}
