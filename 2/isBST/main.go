package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type node struct {
	value int
	left  *node
	right *node
}

func main() {
	in := bufio.NewReader(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	root := createTree(in)

	isBST := checkIsBST(root, math.MinInt, math.MaxInt)

	if isBST {
		writer.WriteString("yes\n")
	} else {
		writer.WriteString("no\n")
	}
}

func checkIsBST(root *node, min, max int) bool {
	if root == nil {
		return true
	}

	if root.value >= min && root.value < max {
		return checkIsBST(root.left, min, root.value) && checkIsBST(root.right, root.value, max)
	}

	return false
}

func createTree(in *bufio.Reader) *node {
	var nodeCount int
	fmt.Fscan(in, &nodeCount)

	nodes := make([]*node, nodeCount)
	for i := range nodes {
		nodes[i] = &node{}
	}

	for _, n := range nodes {
		var value, left, right int
		fmt.Fscan(in, &value, &left, &right)

		n.value = value

		if left != -1 {
			n.left = nodes[left]
		}

		if right != -1 {
			n.right = nodes[right]
		}
	}

	return nodes[0]
}
