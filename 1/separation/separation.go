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

	pivot := col[len(col)-1]
	left, right := 0, len(col)-1

	for i := left; i <= right; i++ {
		if col[i] < pivot {
			col[i], col[left] = col[left], col[i]
			left++
		} else if col[i] > pivot {
			col[i], col[right] = col[right], col[i]
			right--
			i--
		}
	}

	writer.WriteString(strconv.Itoa(left) + " " + strconv.Itoa(right) + "\n")
}
