package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var primeNumbers = []int32{2, 3, 5, 7, 13, 17, 19, 31, 61, 89, 107, 127, 521, 607, 1279, 2203, 2281, 3217, 4253, 4423, 9689,
	9941, 11213, 19937, 21701, 23209, 44497, 86243, 110503, 132049, 216091, 756839, 859433, 1257787,
	1398269, 2976221, 3021377, 6972593, 13466917, 20996011, 24036583, 25964951, 30402457, 32582657,
	37156667, 42643801, 43112609, 57885161, 74207281, 77232917, 82589933}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanLines)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	bSets := getSets(scanner)
	aSets := getSets(scanner)

	aSignatures := getSignatures(aSets)
	bSignatures := getSignatures(bSets)

	similarities := getSimilarities(aSignatures, bSignatures)

	printSimilarities(writer, similarities)
}

func getSets(scanner *bufio.Scanner) [][]string {
	scanner.Scan()
	setsCount, _ := strconv.Atoi(scanner.Text())

	sets := make([][]string, setsCount)
	for i := range sets {
		scanner.Scan()
		for _, str := range strings.Split(scanner.Text(), " ") {
			sets[i] = append(sets[i], str)
		}
	}

	return sets
}

func getSignatures(sets [][]string) [][]int32 {
	signatures := make([][]int32, len(sets))

	for i, set := range sets {
		signatures[i] = getSignature(set)
	}

	return signatures
}

func getSignature(set []string) []int32 {
	signature := make([]int32, len(primeNumbers))

	for i, primeNumber := range primeNumbers {
		signature[i] = getMinHash(set, primeNumber)
	}

	return signature
}

func getMinHash(set []string, seed int32) int32 {
	minHash := int32(math.MaxInt32)
	for _, item := range set {
		hash := getHash(item, seed)

		if hash < minHash {
			minHash = hash
		}
	}

	return minHash
}

func getHash(item string, seed int32) int32 {
	hash := int32(0)

	for _, char := range item {
		hash += int32(char) - 33
		hash *= seed
	}

	return hash
}

func getSimilarities(aSignatures, bSignatures [][]int32) [][]float64 {
	similarities := make([][]float64, len(aSignatures))
	for i, aSignature := range aSignatures {
		similarities[i] = make([]float64, len(bSignatures))

		for j, bSignature := range bSignatures {
			similarities[i][j] = getSimilarity(aSignature, bSignature)
		}
	}

	return similarities
}

func getSimilarity(aSignature, bSignature []int32) float64 {
	matchCount := 0.
	for i := range primeNumbers {
		if aSignature[i] == bSignature[i] {
			matchCount++
		}
	}

	return matchCount / float64(len(primeNumbers))
}

func printSimilarities(writer *bufio.Writer, similarities [][]float64) {
	for i := range similarities {
		for j := range similarities[i] {
			if j != 0 {
				writer.WriteString(" ")
			}
			writer.WriteString(fmt.Sprintf("%.3f", math.Round(similarities[i][j]*1000)/1000))
		}
		writer.WriteString("\n")
	}
}
