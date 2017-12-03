package main

import (
	"fmt"
)

func partition(v []int, left, right int) int {
	pivot := v[right]
	for j := left; j < right; j++ {
		if v[j] < pivot {
			v[j], v[left] = v[left], v[j]
			left++
		}
	}
	v[left], v[right] = v[right], v[left]
	return left
}

func quickSort(v []int, left, right int) {
	if left > right {
		return
	}
	q := partition(v, left, right)
	quickSort(v, left, q-1)
	quickSort(v, q+1, right)
}

func main() {
	v := []int{3, 1, 2, 4}
	quickSort(v, 0, len(v)-1)
	fmt.Println(v)
}
