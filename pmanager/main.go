package main

import (
	"fmt"
	"strings"

	"github.com/bighelmet7/pmanager"
)

func main() {
	paths := []string{"~/deamon/loop.py"}
	runningProcess := pmanager.RunningProcess(paths)
	for _, process := range runningProcess {
		if strings.Contains(process.Name, ".py") {
			fmt.Println(process)
		}
	}
}
