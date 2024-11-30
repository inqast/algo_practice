package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var (
	MNums = []int32{2, 3, 5, 7, 13, 17, 19, 31, 61, 89, 107, 127, 521, 607, 1279, 2203, 2281, 3217, 4253, 4423, 9689,
		9941, 11213, 19937, 21701, 23209, 44497, 86243, 110503, 132049, 216091, 756839, 859433, 1257787,
		1398269, 2976221, 3021377, 6972593, 13466917, 20996011, 24036583, 25964951, 30402457, 32582657,
		37156667, 42643801, 43112609, 57885161, 74207281, 77232917, 82589933}
)

// Signature ...
type Signature []int32

// NewSignature ...
func NewSignature(values []string) Signature {

	sign := make(Signature, 0, len(MNums))

	for _, M := range MNums {
		sign = append(sign, minHash(values, M))
	}

	return sign
}

// Compare ...
func (s Signature) Compare(in Signature) float64 {
	var (
		equalCount uint64
		similarity float64
	)

	if s == nil || in == nil {
		return 0
	}

	for i := range s {
		if s[i] == in[i] {
			equalCount++
		}
	}

	if equalCount == 0 {
		return 0
	}

	similarity = float64(equalCount) / float64(len(s))

	return similarity
}

func minHash(values []string, M int32) int32 {

	var minHash *int32

	for _, value := range values {
		hash := hashString(value, M)
		if minHash == nil || *minHash > hash {
			minHash = &hash
		}
	}

	if minHash == nil {
		return 0
	}

	return *minHash
}

func hashChar(value rune) int32 {
	return value - 33
}

func hashString(value string, M int32) int32 {

	var (
		hash int32
	)

	for _, c := range []rune(value) {
		hash = hash*M + hashChar(c)
	}

	return hash * M
}

func main() {
	var (
		n, m int
	)

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := out.Flush(); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := fmt.Fscanln(in, &n); err != nil {
		log.Fatalln(err)
	}

	r := bufio.NewReader(in)

	setsB := make([]Signature, 0, n)

	for i := 0; i < n; i++ {
		str, err := r.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		str = strings.TrimRight(str, "\n")
		setsB = append(setsB, NewSignature(trim(strings.Split(str, " "))))
	}

	if _, err := fmt.Fscanln(in, &m); err != nil {
		log.Fatalln(err)
	}

	setsA := make([]Signature, 0, m)

	for j := 0; j < m; j++ {
		str, err := r.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		str = strings.TrimRight(str, "\n")
		setsA = append(setsA, NewSignature(trim(strings.Split(str, " "))))
	}

	for i := 0; i < len(setsA); i++ {
		for j := 0; j < len(setsB); j++ {
			if j != 0 {
				if _, err := out.WriteString(" "); err != nil {
					log.Fatalln(err)
				}
			}

			if _, err := out.WriteString(
				fmt.Sprintf(
					"%.3f",
					math.Round(setsA[i].Compare(setsB[j])*1000)/1000)); err != nil {
				log.Fatalln(err)
			}
		}
		if _, err := out.WriteString("\n"); err != nil {
			log.Fatalln(err)
		}
	}
}

func trim(values []string) []string {

	for i := range values {
		newValue := strings.Trim(values[i], " \n")
		values[i] = newValue
	}

	return values
}
