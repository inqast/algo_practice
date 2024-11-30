package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var colLen int
	fmt.Fscan(in, &colLen)

	col := make([]int, colLen)
	for i := range col {
		fmt.Fscan(in, &col[i])
	}

	var needle int
	fmt.Fscan(in, &needle)

	result := findRecursive(col, needle-1)

	writer.WriteString(strconv.Itoa(result) + "\n")
}

func findRecursive(col []int, needle int) int {
	if len(col) == 0 {
		return 0
	}

	left, right := partition(col)
	if needle >= left && needle <= right {
		return col[left]
	}

	if needle < left {
		return findRecursive(col[:left], needle)
	} else {
		return findRecursive(col[right+1:], needle-(right+1))
	}
}

func partition(col []int) (int, int) {
	pivot := col[len(col)-1]
	left, right := 0, len(col)-1

	for i := left; i <= right; i++ {
		if col[i] > pivot {
			col[i], col[left] = col[left], col[i]
			left++
		} else if col[i] < pivot {
			col[i], col[right] = col[right], col[i]
			right--
			i--
		}
	}

	return left, right
}
