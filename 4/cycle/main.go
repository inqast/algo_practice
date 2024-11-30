package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	white = iota
	gray
	black
)

type stackValue struct {
	node int
	prev int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var edgeCount int
	fmt.Fscan(in, &edgeCount)

	graph := make(map[int][]int, edgeCount)
	for i := 0; i < edgeCount; i++ {
		var from, to int
		fmt.Fscan(in, &from, &to)

		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	cycle := findCycle(graph)
	sort.Slice(cycle, func(i, j int) bool {
		return cycle[i] < cycle[j]
	})

	if len(cycle) > 0 {
		for i, vertex := range cycle {
			if i != 0 {
				writer.WriteString(" ")
			}
			writer.WriteString(strconv.Itoa(vertex))
		}
	} else {
		writer.WriteString("-1")
	}
	writer.WriteString("\n")
}

func findCycle(graph map[int][]int) []int {
	color := make(map[int]int, len(graph))
	ancestors := make(map[int]int, len(graph))

	stack := []stackValue{}

	loopedEdge := -1
loop:
	for node := range graph {
		if color[node] != white {
			continue
		}

		stack = append(stack, stackValue{
			node: node,
			prev: -1,
		})
		for len(stack) != 0 {
			from := stack[len(stack)-1].node
			ancestors[from] = stack[len(stack)-1].prev

			stack = stack[:len(stack)-1]

			if color[from] == white {
				color[from] = gray
				stack = append(stack, stackValue{
					node: node,
					prev: ancestors[from],
				})

				for _, to := range graph[from] {

					if color[to] == white {
						stack = append(stack, stackValue{
							node: to,
							prev: from,
						})
					} else if color[to] == gray && ancestors[from] != to {
						loopedEdge = from
						break loop
					}
				}
			} else if color[from] == gray {
				color[from] = black
			}
		}
	}

	result := []int{}
	curEdge := ancestors[loopedEdge]
	for curEdge != loopedEdge && loopedEdge >= 0 {
		result = append(result, curEdge)
		curEdge = ancestors[curEdge]
	}
	result = append(result, loopedEdge)

	return result
}
