package main

import (
	"fmt"
)

type Worker struct {
	Name	string
	Id	int64
	Task	string
}

func main() {
	w := Worker {
		Name: "structExemple.go",
		Id: 7,
		Task: "go run structExemple.go",
	}
	fmt.Printf("This is your worker:\n%v\n", w)
}
