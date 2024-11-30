package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type scanData struct {
	Type  string
	Value string
}

// Scan ...
func (sd *scanData) Scan(state fmt.ScanState, _ rune) error {
	t, err := state.Token(true, func(r rune) bool {
		return r != '\n'
	})
	if err != nil {
		return err
	}

	arrStr := strings.Split(string(t), " ")
	if len(arrStr) != 2 {
		return errors.New("invalid input data")
	}

	sd.Type = arrStr[0]
	sd.Value = arrStr[1]

	return nil
}

type hasher struct{}

func newHasher() *hasher {
	return &hasher{}
}

func (h *hasher) hash(t string, value string) uint64 {
	switch strings.ToLower(t) {
	case "number":
		return h.hashNumber(value)
	case "character":
		if len(value) < 1 {
			return 0
		}
		return h.hashChar(value[0])
	case "string":
		return h.hashString(value, 127)
	default:
		return 0
	}
}

func (h *hasher) hashNumber(value string) uint64 {

	n, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}

	return uint64(n)
}

func (h *hasher) hashChar(value byte) uint64 {
	return uint64(value - 33)
}

func (h *hasher) hashString(value string, M uint64) uint64 {

	var (
		hash uint64
	)

	if len(value) == 0 {
		return 0
	}

	b := []byte(value)

	for i := 0; i < len(value); i++ {
		hash = (hash*M + h.hashChar(b[i])) % uint64(math.MaxUint64)
	}

	return hash
}

func main() {
	var (
		n  int
		sd scanData
	)

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := out.Flush(); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := fmt.Fscan(in, &n); err != nil {
		log.Fatalln(err)
	}

	sds := make([]scanData, 0, n)

	for i := 0; i < n; i++ {
		if _, err := fmt.Fscan(in, &sd); err != nil {
			log.Fatalln(err)
		}
		sds = append(sds, sd)
	}

	h := newHasher()

	for _, sd := range sds {
		hash := h.hash(sd.Type, sd.Value)
		if _, err := fmt.Fprintln(out, hash); err != nil {
			log.Fatalln(err)
		}
	}
}
