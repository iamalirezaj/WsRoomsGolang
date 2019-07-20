package main

import (
	"os"
	"log"
	"flag"
	"net"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	var (
		socketFile = flag.String("sock", "/restream.app/sockets/websocket.sock", "Socket file")
	)

	os.Remove(*socketFile)

	unixListener, err := net.Listen("unix", *socketFile)
	if err != nil {
		panic(err)
	}

	flag.Parse()
	log.SetFlags(0)

	hub := NewHub()
	router := mux.NewRouter()

	router.HandleFunc("/stream/{stream_hash_id}", hub.HandleWS).Methods("GET")

	http.Handle("/", router)

	defer unixListener.Close()

	log.Printf("http_err: %v", http.Serve(unixListener, nil))
}