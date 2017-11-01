/*Selection sort in golang*/
package main

import "fmt"

func maxPosition(v []int, m int) int {
	k := 0
	for i := 1; i <= m; i++ {
		if v[i] > v[k] {
			k = i
		}
	}
	return k
}

//selectionSort has a complexity of Theta(n^2)
func selectionSort(v []int, n int) {
	for i := n - 1; i >= 0; i-- {
		maxPos := maxPosition(v, i)
		v[maxPos], v[i] = v[i], v[maxPos]
	}
}

func main() {
	v := []int{3, 1, 2, 4}
	selectionSort(v, len(v))
	fmt.Println(v)
}
