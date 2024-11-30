package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	white = iota
	gray
	black
)

func main() {
	in := bufio.NewReader(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var start, sword, finish int
	fmt.Fscan(in, &start, &sword, &finish)

	var pathCount int
	fmt.Fscan(in, &pathCount)

	graph := make(map[int][]int, pathCount)
	for i := 0; i < pathCount; i++ {
		var from, to int
		fmt.Fscan(in, &from, &to)

		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	toSword, found := bfs(graph, start, sword, finish)
	if !found {
		writer.WriteString("-1")

		return
	}

	toFinish, found := bfs(graph, sword, finish, -1)
	if !found {
		writer.WriteString("-1")

		return
	}

	writer.WriteString(strconv.Itoa(toSword + 1 + toFinish))
}

func bfs(graph map[int][]int, start, finish, excluded int) (int, bool) {
	queue := []int{start}
	color := make(map[int]int, len(graph))
	color[start] = gray

	distances := make(map[int]int, len(graph))

	for i := 0; i < len(queue); i++ {
		curGlade := queue[i]

		if curGlade == finish {
			return distances[curGlade], true
		}

		for _, neighbour := range graph[curGlade] {
			if color[neighbour] == white && neighbour != excluded {
				queue = append(queue, neighbour)
				color[neighbour] = gray
				distances[neighbour] = distances[curGlade] + 1
			}
		}

		color[curGlade] = black
	}

	return 0, false
}
