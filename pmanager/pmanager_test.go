package pmanager

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func reviveProcess(path string) {
	command := fmt.Sprintf("(python %s &)", path)
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Start() // we dont wait to finish the command
	if err != nil {
		log.Fatal(err)
	}
}

func TestKeepAlive(t *testing.T) {
	path := "~/t1.py"
	process := []Process{{Command: "python t1.py", Name: "t1.py"}}
	var found bool
	for _, p := range process {
		if strings.Contains(path, p.Name) {
			t.Logf("%s is running", path)
			found = true
		}
	}
	if !found {
		t.Logf("%s is dead, start the reviving operation", path)
		reviveProcess(path)
	}
	return
}

func TestParseProcessInfo(t *testing.T) {
	//Remember ps aux | awk '{print $1, $2, $3, $4, $11, $12}'
	//so the $1 is the user, $2 pid, and soo on...
	info := `bighelmet7 111 4,9 1,6 /Applications/Test/Contents/Test/8.8.8.8/TestGo test.go
	`
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
	t.Skip(Process{User: dividedInfo[0], Pid: pid, CPU: cpu, MEM: mem, Command: dividedInfo[4], Name: dividedInfo[5]})
	return
}

func TestRunningProcess(t *testing.T) {
	command := "ps aux | awk '{print $1, $2, $3, $4, $11, $12}'"
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(output), "\n")
	lines = append(lines[:0], lines[1:len(lines)-1]...)
	var processSlice []Process
	for _, line := range lines {
		processSlice = append(processSlice, Process{User: "Test", Pid: 1, CPU: 1.1, MEM: 1.1, Command: line, Name: "Test.go"})
	}
	t.Skip(processSlice)
	return
}
