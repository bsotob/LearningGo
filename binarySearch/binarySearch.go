package main

import "fmt"

func binarySearch(v []int, x, left, right int) int {
	if left <= right {
		middle := (left + right) / 2
		if x < v[middle] {
			return binarySearch(v, x, left, middle-1)
		} else if x > v[middle] {
			return binarySearch(v, x, middle+1, right)
		} else {
			return middle
		}
	}
	return -1
}

func main() {
	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := 3
	if binarySearch(v, x, 0, len(v)-1) != -1 {
		fmt.Println("EXIST")
	}
}
