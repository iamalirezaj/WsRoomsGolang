package main

import (
	"log"
	"flag"
	"net/http"
	"github.com/gorilla/mux"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {

	flag.Parse()
	log.SetFlags(0)

	hub := NewHub()
	router := mux.NewRouter()

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	router.HandleFunc("/ws/{room}", hub.HandleWS).Methods("GET")

	http.Handle("/", router)
	log.Printf("http_err: %v", http.ListenAndServe(*addr, nil))
}