package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type node struct {
	key   string
	value string
	next  *node
}

type hashTable struct {
	capacity int
	size     int
	buckets  []*node
}

func main() {
	in := bufio.NewReader(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	ht := newHashTable()

	var commandsCount int
	fmt.Fscan(in, &commandsCount)
	for commandIdx := 0; commandIdx < commandsCount; commandIdx++ {
		var command string
		fmt.Fscan(in, &command)

		switch command {
		case "put":
			var key, value string
			fmt.Fscan(in, &key, &value)

			ht.put(key, value)
			if ht.isResizeRequired() {
				ht.resize()
			}
		case "get":
			var key string
			fmt.Fscan(in, &key)

			value, ok := ht.get(key)
			if ok {
				writer.WriteString("+")
			} else {
				writer.WriteString("-")
			}
			writer.WriteString(value + "\n")
		case "print":
			writer.WriteString(ht.String())
		}
	}
}

func newHashTable() *hashTable {
	return &hashTable{
		capacity: 4,
		buckets:  make([]*node, 4),
	}
}

func (ht *hashTable) put(key, value string) {
	hash := ht.getHash(key)

	head := ht.buckets[hash]

	if head == nil {
		ht.buckets[hash] = &node{
			key:   key,
			value: value,
		}
		ht.size++

		return
	}

	for ; ; head = head.next {
		if head.key == key {
			head.value = value

			return
		}

		if head.next == nil {
			break
		}
	}

	head.next = &node{
		key:   key,
		value: value,
	}
	ht.size++
}

func (ht *hashTable) get(key string) (string, bool) {
	hash := ht.getHash(key)

	for head := ht.buckets[hash]; head != nil; head = head.next {
		if head.key == key {
			return head.value, true
		}
	}

	return "", false
}

func (ht *hashTable) resize() {
	oldBuckets := ht.buckets

	ht.capacity *= 2
	ht.buckets = make([]*node, ht.capacity)
	ht.size = 0

	for _, bucket := range oldBuckets {
		ht.evacuateRecursive(bucket)
	}
}

func (ht *hashTable) evacuateRecursive(head *node) {
	if head == nil {
		return
	}

	ht.evacuateRecursive(head.next)

	ht.put(head.key, head.value)
}

func (ht *hashTable) isResizeRequired() bool {
	return math.Ceil((float64(ht.size)/float64(ht.capacity))*100)/100 >= 0.75
}

func (ht *hashTable) getHash(key string) int {
	var hash int

	for _, char := range key {
		hash *= 31
		hash += int(char) - 33
		hash %= ht.capacity
	}

	return hash
}

func (ht *hashTable) String() string {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("%d %d\n", ht.size, ht.capacity))
	for _, bucket := range ht.buckets {
		head := bucket
		for head != nil {
			builder.WriteString(fmt.Sprintf("\t%s %s", head.key, head.value))
			head = head.next
		}

		builder.WriteString("\n")
	}

	return builder.String()
}
