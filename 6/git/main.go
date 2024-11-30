package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	leftLen, _ := strconv.Atoi(scanner.Text())

	left := make([]string, leftLen+1)
	for i := 1; i < len(left); i++ {
		scanner.Scan()
		left[i] = scanner.Text()
	}

	scanner.Scan()
	rightLen, _ := strconv.Atoi(scanner.Text())

	right := make([]string, rightLen+1)
	for i := 1; i < len(right); i++ {
		scanner.Scan()
		right[i] = scanner.Text()
	}

	dp := make([][]int, 2)

	dp[0] = make([]int, len(right))
	dp[1] = make([]int, len(right))
	for j := 0; j < len(right); j++ {
		dp[0][j] = j
	}

	curr := 1
	prev := 0
	for i := 1; i < len(left); i++ {
		dp[curr][0] = i

		for j := 1; j < len(right); j++ {
			prevVal := dp[prev][j-1]
			if left[i] != right[j] {
				prevVal++
			}

			dp[curr][j] = getMin(dp[prev][j]+1, dp[curr][j-1]+1, prevVal)
		}

		curr = 1 - curr
		prev = 1 - prev
	}

	writer.WriteString(strconv.Itoa(dp[prev][len(right)-1]))
}

func getMin(upper, left, prev int) int {
	min := upper

	if left < min {
		min = left
	}

	if prev < min {
		min = prev
	}

	return min
}
