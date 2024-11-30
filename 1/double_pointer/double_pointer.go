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

	var col1Len, col2Len int
	fmt.Fscan(in, &col1Len, &col2Len)

	col1 := make([]int, col1Len)
	for i := range col1 {
		fmt.Fscan(in, &col1[i])
	}

	col2 := make([]int, col2Len)
	for i := range col2 {
		fmt.Fscan(in, &col2[i])
	}

	var (
		col1Idx int
		col2Idx int
		result  = make([]int, 0, col1Len+col2Len)
	)
	for col1Idx < len(col1) && col2Idx < len(col2) {
		if col1[col1Idx] <= col2[col2Idx] {
			result = append(result, col1[col1Idx])
			col1Idx++
		} else {
			result = append(result, col2[col2Idx])
			col2Idx++
		}
	}

	if col1Idx < len(col1) {
		result = append(result, col1[col1Idx:]...)
	} else {
		result = append(result, col2[col2Idx:]...)
	}

	for i, num := range result {
		writer.WriteString(strconv.Itoa(num))
		if i < len(result)-1 {
			writer.WriteString(" ")
		}
	}
}
