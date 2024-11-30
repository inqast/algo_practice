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

	writer.WriteString(strconv.Itoa(countInversions(col)))
}

func countInversions(col []int) int {
	if len(col) <= 1 {
		return 0
	}

	left := col[:len(col)/2]
	right := col[len(col)/2:]

	leftInversionsCount := countInversions(left)
	rightInversionsCount := countInversions(right)

	sortedRight := merge(right[:len(right)/2], right[len(right)/2:])
	for i := range right {
		right[i] = sortedRight[i]
	}

	var currentInversions int
	for _, leftNum := range left {
		currentInversions += countBinary(leftNum, right)
	}

	return leftInversionsCount + rightInversionsCount + currentInversions
}

func merge(left, right []int) []int {
	var leftIdx, rightIdx int

	sorted := make([]int, 0, len(left)+len(right))
	for leftIdx < len(left) && rightIdx < len(right) {
		if left[leftIdx] <= right[rightIdx] {
			sorted = append(sorted, left[leftIdx])
			leftIdx++
		} else {
			sorted = append(sorted, right[rightIdx])
			rightIdx++
		}
	}

	if leftIdx < len(left) {
		sorted = append(sorted, left[leftIdx:]...)
	} else {
		sorted = append(sorted, right[rightIdx:]...)
	}

	return sorted
}

func countBinary(target int, col []int) int {
	leftIdx := 0
	rightIdx := len(col) - 1

	if target <= col[leftIdx] {
		return 0
	} else if target > col[rightIdx] {
		return rightIdx - leftIdx + 1
	}

	for rightIdx-leftIdx > 1 {
		mid := (leftIdx + rightIdx) / 2
		if col[mid] < target {
			leftIdx = mid
		} else {
			rightIdx = mid
		}
	}

	return leftIdx + 1
}
