package main

import "fmt"

type Queue []int

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(i int) {
	*q = append(*q, i)
}

func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}

func bfs(graph [][]int, start int) {
	visited := make([]bool, len(graph))
	queue := Queue{}

	queue.Enqueue(start)
	visited[start] = true

	for !queue.IsEmpty() {
		vertex, _ := queue.Dequeue()
		fmt.Print(vertex, " ")

		for i := 0; i < len(graph); i++ {
			if graph[vertex][i] == 1 && !visited[i] {
				queue.Enqueue(i)
				visited[i] = true
			}
		}
	}
}
