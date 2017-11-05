package main

import "fmt"

//quickPow is Theta(log n) time complexity
func quickPow(x, n int) int {
	if n == 0 {
		return 1
	}
	k := quickPow(x, n/2)
	if n%2 == 0 {
		return k * k
	}
	return x * k * k
}

func main() {
	x, n := 3, 2
	fmt.Printf("%v times %v is %v\n", n, x, quickPow(x, n))
}
