package main

import (
	"os"
	"log"
	"flag"
	"net"
	"net/http"
	"github.com/gorilla/mux"
)

var sockFile = flag.String("sock", "/restream.web/sockets/websocket.sock", "Socket file")

func main() {

	os.Remove(*sockFile)

	unixListener, err := net.Listen("unix", *sockFile)
	if err != nil {
		panic(err)
	}

	flag.Parse()
	log.SetFlags(0)

	hub := NewHub()
	router := mux.NewRouter()

	router.HandleFunc("/stream/{stream_hash_id}", hub.HandleWS).Methods("GET")

	http.Handle("/", router)
	log.Printf("http_err: %v", http.Serve(unixListener, nil))
}