package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	var addr string
	if port == "" {
		addr = "localhost:8080"
	} else {
		addr = ":" + port
	}
	mux := mux.NewRouter()
	mux.HandleFunc("/", helloHandle)
	mux.HandleFunc("/{name}", sayHi)
	log.Println(http.ListenAndServe(addr, mux))
}
