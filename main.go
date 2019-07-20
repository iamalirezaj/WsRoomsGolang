package main

import (
	"log"
	"flag"
	"net/http"
	"github.com/gorilla/mux"
)

var addr = flag.String("addr", "0.0.0.0:8080", "http service address")

func main() {

	flag.Parse()
	log.SetFlags(0)

	hub := NewHub()
	router := mux.NewRouter()

	router.HandleFunc("/stream/{stream_hash_id}", hub.HandleWS).Methods("GET")

	http.Handle("/", router)
	log.Printf("http_err: %v", http.ListenAndServe(*addr, nil))
}