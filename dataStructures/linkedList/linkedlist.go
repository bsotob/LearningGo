//awful exemple should recode this with a real exemple
package main

import "fmt"

//Node is a node
type Node struct {
	value string
	next  *Node
}

//List of students where the firts in the line be always be the teacher
func main() {
	student := &Node{"student1", nil}
	teacher := &Node{"Teacher", student}
	fmt.Println(teacher)
	fmt.Println(student)
}
