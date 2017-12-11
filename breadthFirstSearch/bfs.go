//You can do it better m8, just read more go, this could be done more
//efficiently ;)
package main

import "fmt"

//Queue is a simple implementation of a queue in golang
type Queue struct {
	Elements []int64
}

//front returns the first element in the queue
func (q *Queue) front() int64 {
	return q.Elements[0]
}

//pop deletes the top element in the queue
func (q *Queue) pop() {
	if len(q.Elements) > 0 {
		q.Elements = q.Elements[1:]
	} else {
		var emp []int64
		q.Elements = emp
	}
}

//push append an element to the queue
func (q *Queue) push(elem int64) {
	q.Elements = append(q.Elements, elem)
}

//empty tell us if the queue is empty
func (q *Queue) empty() bool {
	return len(q.Elements) == 0
}

//bfs for a G = (V , E) where |V| = 8
func bfs(G [8][]int, start int64) []int64 {
	Q := &Queue{} //declares our queue
	var (
		visited [8]bool // visited elements
		L       []int64 //slice of all the nodes that start can go
	)
	Q.push(start)
	visited[start] = true
	for !Q.empty() {
		v := Q.front()
		Q.pop()
		for _, elem := range G[v] {
			if !visited[elem] {
				visited[elem] = true
				Q.push(int64(elem))
				for !Q.empty() {
					m := Q.front()
					Q.pop()
					for _, elemJ := range G[m] {
						if !visited[elemJ] {
							visited[elemJ] = true
							Q.push(int64(elemJ))
						}
					}
					L = append(L, m)
				}
			}
		}
	}
	return L
}

func main() {
	//Exemple graph
	//    0
	//   / \
	//  1 - 2 - 3 - 4
	//      |       |
	//      7   6   5
	//Adjacencies Matrix:
	//   0 1 2 3 4 5 6 7
	// 0 0 1 1 0 0 0 0 0
	// 1 1 0 1 0 0 0 0 0
	// 2 1 1 0 1 0 0 0 1
	// 3 0 0 1 0 1 0 0 0
	// 4 0 0 0 1 0 1 0 0
	// 5 0 0 0 0 1 0 1 0
	// 6 0 0 0 0 0 1 0 0
	// 7 0 0 1 0 0 0 0 0
	graph := [8][]int{}
	graph[0] = []int{1, 2}
	graph[1] = []int{0, 2}
	graph[2] = []int{0, 1, 3, 7}
	graph[3] = []int{2, 4}
	graph[4] = []int{3, 5}
	graph[5] = []int{4}
	graph[6] = []int{}
	graph[7] = []int{2}
	fmt.Println(graph)
	var start int64 = 6
	fmt.Println(bfs(graph, start))
}
