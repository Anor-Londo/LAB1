package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", helloHandle)
	mux.HandleFunc("/{name}", sayHi)
	log.Println(http.ListenAndServe("localhost:8080", mux))
}
