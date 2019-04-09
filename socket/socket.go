package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	var (
		network = "tcp"
		addr    = ":8080"
	)
	conn, err := net.Dial(network, addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	body := bytes.NewBufferString("GET /home HTTP/1.0\r\n\r\n")
	if _, err := conn.Write(body.Bytes()); err != nil {
		log.Fatalln(err)
	}

	var data bytes.Buffer

	io.Copy(&data, conn)
	fmt.Printf("%s\n", data.String())
}
