package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type node struct {
	children []*node
}

func main() {
	in := bufio.NewReader(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	root := createTree(in)

	height := countRecursive(root)

	writer.WriteString(strconv.Itoa(height) + "\n")
}

func countRecursive(root *node) int {
	if len(root.children) == 0 {
		return 0
	}

	var maxHeight int
	for _, child := range root.children {
		childHeight := countRecursive(child)
		if childHeight > maxHeight {
			maxHeight = childHeight
		}
	}

	return maxHeight + 1
}

func createTree(in *bufio.Reader) *node {
	var nodeCount int
	fmt.Fscan(in, &nodeCount)

	nodes := make([]*node, nodeCount)
	for i := range nodes {
		nodes[i] = &node{}
	}

	var root *node
	for _, n := range nodes {
		var parent int
		fmt.Fscan(in, &parent)

		if parent == -1 {
			root = n
			continue
		}

		nodes[parent].children = append(nodes[parent].children, n)
	}

	return root
}
