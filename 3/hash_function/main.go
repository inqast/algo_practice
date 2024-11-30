package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type (
	hashable interface {
		getHash() uint64
	}

	hashableInt64  int64
	hashableByte   byte
	hashableString []hashableByte
)

func (hi hashableInt64) getHash() uint64 {
	if hi < 0 {
		return uint64(math.MaxUint64) - uint64(-hi) + 1
	}

	return uint64(hi)
}

func (hb hashableByte) getHash() uint64 {
	return uint64(hb - 33)
}

func (hs hashableString) getHash() uint64 {
	var hash uint64

	for _, char := range hs {
		hash *= 127
		hash += char.getHash()
		hash %= uint64(math.MaxUint64)
	}

	return hash
}

func getHashableType(valueType, value string) hashable {
	switch valueType {
	case "number":
		intValue, _ := strconv.ParseInt(value, 10, 64)

		return hashableInt64(intValue)
	case "character":
		return hashableByte(value[0])
	case "string":
		hs := make(hashableString, len(value))

		for i := range hs {
			hs[i] = hashableByte(value[i])
		}

		return hs
	}

	return nil
}

func main() {
	in := bufio.NewReader(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var valuesCount int
	fmt.Fscan(in, &valuesCount)

	for i := 0; i < valuesCount; i++ {
		var valueType, value string
		fmt.Fscan(in, &valueType, &value)

		hash := getHashableType(valueType, value).getHash()

		writer.WriteString(strconv.FormatUint(hash, 10) + "\n")
	}
}
