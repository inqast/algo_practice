package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type edge struct {
	weight int
	to     int
}

type node struct {
	dist int
	idx  int
}

func (left node) isMoreImportant(right node) bool {
	return left.dist < right.dist
}

type heap []node

func (h heap) siftUp() {
	idx := len(h) - 1
	for {
		if h[idx].isMoreImportant(h[idx/2]) && idx != 1 {
			h[idx/2], h[idx] = h[idx], h[idx/2]
			idx /= 2
			continue
		}

		break
	}
}

func (h heap) siftDown() {
	idx := 1

	for {
		left := idx * 2
		right := idx*2 + 1

		if left >= len(h) {
			return
		}

		var maxIdx int
		if right < len(h) && h[right].isMoreImportant(h[left]) {
			maxIdx = right
		} else {
			maxIdx = left
		}

		if h[maxIdx].isMoreImportant(h[idx]) {
			h[idx], h[maxIdx] = h[maxIdx], h[idx]
			idx = maxIdx
		} else {
			return
		}
	}
}

func newHeap() heap {
	return make(heap, 1)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var vertexCount, edgeCount int
	fmt.Fscan(in, &vertexCount, &edgeCount)

	var start, end int
	fmt.Fscan(in, &start, &end)

	adjacencyList := make([][]edge, vertexCount+1)
	for i := 0; i < edgeCount; i++ {
		var from, to, weight int
		fmt.Fscan(in, &from, &to, &weight)

		adjacencyList[from] = append(adjacencyList[from], edge{
			weight: weight,
			to:     to,
		})
	}

	cost := dijkstra(adjacencyList, start, end)

	writer.WriteString(strconv.Itoa(cost) + "\n")
}

func dijkstra(adjacencyList [][]edge, from, to int) int {
	minWeights := make([]int, len(adjacencyList))
	for i := 0; i < len(adjacencyList); i++ {
		minWeights[i] = math.MaxInt
	}
	minWeights[from] = 0
	visited := make([]bool, len(adjacencyList))

	nodesHeap := newHeap()
	nodesHeap = addNode(node{
		dist: minWeights[from],
		idx:  from,
	}, nodesHeap)

	var minNode node
	for len(nodesHeap) > 1 {
		nodesHeap, minNode = extractMin(nodesHeap)
		if visited[minNode.idx] {
			continue
		}

		for _, link := range adjacencyList[minNode.idx] {
			if !visited[link.to] {
				relax(minWeights, link, minNode.idx)
				nodesHeap = addNode(node{
					dist: minWeights[link.to],
					idx:  link.to,
				}, nodesHeap)
			}
		}

		visited[minNode.idx] = true
	}

	if visited[to] {
		return minWeights[to]
	} else {
		return -1
	}
}

func addNode(
	n node,
	targetEdges heap,
) heap {
	targetEdges = append(targetEdges, n)
	targetEdges.siftUp()

	return targetEdges
}

func extractMin(edges heap) (heap, node) {
	min := edges[1]

	edges[1] = edges[len(edges)-1]
	edges = edges[:len(edges)-1]

	edges.siftDown()

	return edges, min
}

func relax(minWeights []int, link edge, nodeIdx int) {
	if minWeights[link.to] > minWeights[nodeIdx]+link.weight {
		minWeights[link.to] = minWeights[nodeIdx] + link.weight
	}
}
