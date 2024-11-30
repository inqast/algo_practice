package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	key  string
	next *node
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
		case "add":
			var key string
			fmt.Fscan(in, &key)

			ht.put(key)
			if ht.isResizeRequired() {
				ht.resize()
			}
		case "remove":
			var key string
			fmt.Fscan(in, &key)

			ht.remove(key)
		case "contains":
			var key string
			fmt.Fscan(in, &key)

			ok := ht.contains(key)
			if ok {
				writer.WriteString("+")
			} else {
				writer.WriteString("-")
			}
			writer.WriteString("\n")
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

func (ht *hashTable) put(key string) {
	hash := ht.getHash(key)

	head := ht.buckets[hash]

	if head == nil {
		ht.buckets[hash] = &node{
			key: key,
		}
		ht.size++

		return
	}

	for ; ; head = head.next {
		if head.key == key {
			return
		}

		if head.next == nil {
			break
		}
	}

	head.next = &node{
		key: key,
	}
	ht.size++
}

func (ht *hashTable) remove(key string) {
	hash := ht.getHash(key)

	var prev *node
	for head := ht.buckets[hash]; head != nil; head = head.next {
		if head.key == key {
			if prev != nil && head.key != prev.key {
				prev.next = head.next
			} else {
				ht.buckets[hash] = head.next
			}

			ht.size--

			break
		}

		prev = head
	}
}

func (ht *hashTable) contains(key string) bool {
	hash := ht.getHash(key)

	for head := ht.buckets[hash]; head != nil; head = head.next {
		if head.key == key {
			return true
		}
	}

	return false
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

	ht.put(head.key)
}

func (ht *hashTable) isResizeRequired() bool {
	return float64(ht.size)/float64(ht.capacity) >= 0.75
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
			builder.WriteString(fmt.Sprintf(" %s", head.key))
			head = head.next
		}

		builder.WriteString("\n")
	}

	return builder.String()
}
