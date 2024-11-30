package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	nArr := make([]int, n)

	for i := range nArr {
		fmt.Fscan(in, &nArr[i])
	}
	mArr := make([]int, m)
	for i := range mArr {
		fmt.Fscan(in, &mArr[i])
	}

	cN := 0
	cM := 0

	if nArr[n-1] <= mArr[0] {
		for i := 0; i < n; i++ {
			fmt.Fprint(out, nArr[i], " ")
		}
		for i := 0; i < m; i++ {
			fmt.Fprint(out, mArr[i], " ")
		}
	} else if mArr[n-1] <= nArr[0] {
		for i := 0; i < m; i++ {
			fmt.Fprint(out, mArr[i], " ")
		}
		for i := 0; i < n; i++ {
			fmt.Fprint(out, nArr[i], " ")
		}
	} else {
		for {
			if nArr[cN] < mArr[cM] {
				fmt.Fprint(out, nArr[cN], " ")

				if cN == n-1 { //attach second
					for i := cM; i < m; i++ {
						fmt.Fprint(out, mArr[i], " ")
					}
					break
				}

				cN++

			} else {

				fmt.Fprint(out, mArr[cM], " ")
				if cM == m-1 { //attach second
					for i := cN; i < n; i++ {
						fmt.Fprint(out, nArr[i], " ")
					}
					break
				}
				cM++
			}

		}
	}

	//fmt.Fprint(out, nArr)
	//.Fprint(out, mArr)
}
