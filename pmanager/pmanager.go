package pmanager

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)


type Process struct {
	User    string
	Pid     int64
	CPU     float64
	MEM     float64
	Command string
	Name    string
	Alive   bool
}

//ReviveProcess restart the process by his path
func ReviveProcess(path string) {
	//TODO DOESNT WORK
	command := fmt.Sprintf("(python %s &)", path)
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Start() // we dont wait to finish the command
	if err != nil {
		log.Fatal(err)
	}
}

//isAlive returns true if the name process was specified to be alive
func isAlive(paths []string, name string) bool {
	for _, path := range paths {
		if strings.Contains(path, name) {
			return true
		}
	}
	return false
}

//ParseProcessInfo parses the information catched from the ps command
//and fills the Process struct to be return
func ParseProcessInfo(info string, paths []string) Process {
	//Remember ps aux | awk '{print $1, $2, $3, $4, $11}'
	//so the $1 is the user, $2 pid, and soo on...
	info = strings.Trim(info, "\n")
	dividedInfo := strings.Split(info, " ")
	if len(dividedInfo) <= 0 || len(dividedInfo) <= 4 {
		log.Fatalf("Not enough Process information\n")
	}
	pid, err := strconv.ParseInt(dividedInfo[1], 0, 64)
	if err != nil {
		log.Fatal(err)
	}
	cpu, err := strconv.ParseFloat(strings.Replace(dividedInfo[2], ",", ".", -1), 64)
	if err != nil {
		log.Fatal(err)
	}
	mem, err := strconv.ParseFloat(strings.Replace(dividedInfo[3], ",", ".", -1), 64)
	if err != nil {
		log.Fatal(err)
	}
	return Process{
		User:    dividedInfo[0],
		Pid:     pid,
		CPU:     cpu,
		MEM:     mem,
		Command: dividedInfo[4],
		Name:    dividedInfo[5],
		Alive:   isAlive(paths, dividedInfo[5]),
	}
}

//RunningProcess returns all the running process that were given by
//the ps command, and returns a slice with all the active Process.
//the paths parameters is a slice of all the process that must be
//running, if the Process was found the KeepAlive atribute will be
//true, otherwise will be set to false
func RunningProcess(paths []string) []Process {
	command := "ps aux | awk '{print $1, $2, $3, $4, $11, $12}'"
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(output), "\n")
	//Ignore the headers given by ps, and the last print line
	lines = append(lines[:0], lines[1:len(lines)-1]...)
	var processSlice []Process
	for _, line := range lines {
		processSlice = append(processSlice, ParseProcessInfo(line, paths))
	}
	return processSlice
}
