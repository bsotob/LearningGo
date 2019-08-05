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
	mux.HandleFunc("/ping", logger(ping))

	log.Println("Setting server with config...")
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Println("Service running")
	server.ListenAndServe()
}

func logger(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.String(), r.Proto)
	}
}

// ping is a dummy function to response ping.
func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong\n")
}
