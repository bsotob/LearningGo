package main

import (
	"fmt"
)

type WorkerStatus string

//State englobs all the possible status of a worker
const ( //Could be done with iota
	Ready WorkerStatus = "Ready"
	Run WorkerStatus = "Run"
	Block WorkerStatus = "Block"
)

/*
type WorkerStatus int64
const State (
	Ready WorkerStatus = iota
	Run
	Block
)
*/

type Worker struct {
	Name string
	ID int64
	Task string
	Status WorkerStatus
}

func main() {
	w := Worker{
		Name: "structEnum.go",
		ID: 9,
		Task: "go run structEnum.go",
		Status: Ready,
	}
	fmt.Printf("This is your worker: %v\n", w)
}
