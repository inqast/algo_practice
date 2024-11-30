package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	inOrderValues := inOrder(root)
	printRow(writer, inOrderValues)

	preOrderValues := preOrder(root)
	printRow(writer, preOrderValues)

	postOrderValues := postOrder(root)
	printRow(writer, postOrderValues)
}

func inOrder(root *node) []int {
	if root == nil {
		return []int{}
	}

	var values []int

	values = append(values, inOrder(root.left)...)

	values = append(values, root.value)

	values = append(values, inOrder(root.right)...)

	return values
}

func preOrder(root *node) []int {
	if root == nil {
		return []int{}
	}

	var values []int

	values = append(values, root.value)

	values = append(values, preOrder(root.left)...)
	values = append(values, preOrder(root.right)...)

	return values
}

func postOrder(root *node) []int {
	if root == nil {
		return []int{}
	}

	var values []int

	values = append(values, postOrder(root.left)...)
	values = append(values, postOrder(root.right)...)

	values = append(values, root.value)

	return values
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

func printRow(out *bufio.Writer, row []int) {
	for i, num := range row {
		if i != 0 {
			out.WriteString(" ")
		}

		out.WriteString(strconv.Itoa(num))
	}

	out.WriteString("\n")
}
