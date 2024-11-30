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

	var boxCount int
	fmt.Fscan(in, &boxCount)

	var (
		boxes         = make([]int, boxCount)
		maxCostForBox = make([]int, boxCount)
	)

	for i := range boxes {
		fmt.Fscan(in, &boxes[i])
	}

	absoluteMax := 0
	for i, currentCost := range boxes {
		prevEvenMaxCost := getPrevOdd(maxCostForBox, i, 2)
		prevOddMaxCost := getPrevOdd(maxCostForBox, i, 1)

		maxCostForBox[i] = maxVal(prevEvenMaxCost+currentCost, prevOddMaxCost)
		if maxCostForBox[i] > absoluteMax {
			absoluteMax = maxCostForBox[i]
		}
	}

	writer.WriteString(strconv.Itoa(absoluteMax) + "\n")
}

func getPrevOdd(maxCostForBox []int, i, shift int) int {
	if i < shift {
		return 0
	}

	return maxCostForBox[i-shift]
}

func maxVal(left, right int) int {
	if left > right {
		return left
	}

	return right
}
