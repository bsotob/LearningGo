package main

import "fmt"

//inserctionSort has a time complexity of Theta(n^2)
func inserctionSort(v []int, n int) {
	for k := 1; k < n; k++ {
		for t := k - 1; t >= 0 && v[t+1] < v[t]; t-- {
			v[t+1], v[t] = v[t], v[t+1]
		}
	}
}

func main() {
	v := []int{4, 2, 3, 1}
	inserctionSort(v, len(v))
	fmt.Println(v)
}
