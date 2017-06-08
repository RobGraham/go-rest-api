package main

import (
	"log"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
