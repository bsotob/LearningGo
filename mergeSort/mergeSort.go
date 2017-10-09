package main

import "fmt"

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}

func mergeSort(v []int) []int {
	if len(v) < 2 {
		return v
	}
	median := len(v) / 2
	left := mergeSort(v[:median])
	right := mergeSort(v[median:])
	return merge(left, right)
}

func main() {
	v := []int{3, 2, 4, 1, 5, 7, 6, 8, 9, 11}
	v = mergeSort(v)
	fmt.Println(v)
}
