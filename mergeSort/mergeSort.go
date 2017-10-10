package main

import "fmt"

func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	result := make([]int, 0, size)
	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			result[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			result[k] = left[i]
			i++
		} else if left[i] < right[i] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
	}
	return result
}

func mergeSort(v []int) []int {
	if len(v) < 2 {
		return v
	}
	mid := len(v) / 2
	return merge(mergeSort(v[:mid]), mergeSort(v[mid:]))
}

func main() {
	v := []int{3, 2, 4, 1, 5, 7, 6, 8, 9, 11}
	v = mergeSort(v)
	fmt.Println(v)
}
