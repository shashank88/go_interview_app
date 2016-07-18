package main

import (
	"net/http"
	"os"
)

func main() {
	initApp()
	router := NewRouter()
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
