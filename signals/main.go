package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	select {
	case <- c:
		log.Println("SIGINT received")
	}
}
