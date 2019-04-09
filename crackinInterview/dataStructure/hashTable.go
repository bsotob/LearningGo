package main

import "log"

func main() {
	// m its a map internally golang implements it as a hash table.
	m := make(map[string]int)
	m["test"] = 0
	if val, ok := m["test"]; ok {
		log.Printf("The key searched is %s and has a value of %d\n", "test", val)
	} else {
		log.Fatalln("The key doesnt exists")
	}
}
