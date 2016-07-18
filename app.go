package main

import (
	"log"
	"net/http"
)

func main() {
	initApp()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":5000", router))
}
