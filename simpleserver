package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	/*
		Multiplexor.
		Address.
		Handler.
		Routing.
		Request & Response.
		Other Parameters.
	*/
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)

	log.Println("Setting server with config...")
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Println("Service running")
	server.ListenAndServe()
}

// ping is a dummy function to response ping.
func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong\n")
}
