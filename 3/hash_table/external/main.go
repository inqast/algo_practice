package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	run(in, out)
}

// ----------- HashMap ----------

const hashCode = 31

type ListNode struct {
	key, value string
	next       *ListNode
}

type HashMap struct {
	buckets []*ListNode
}

func (m *HashMap) Put(key, value string) {
	var prev *ListNode
	k := m.hash(key)

	v := m.buckets[k]
	if v == nil {
		m.buckets[k] = &ListNode{key: key, value: value}
		goto ext
	}

	for v != nil {
		if v.key == key {
			v.value = value
			return
		}
		prev = v
		v = v.next
	}

	prev.next = &ListNode{key: key, value: value}

ext:
	if m.Len() >= m.Cap()*3/4 {
		m.extend()
	}
}

func (m *HashMap) Get(key string) (string, bool) {
	k := m.hash(key)

	v := m.buckets[k]
	for v != nil {
		if v.key == key {
			return v.value, true
		}
		v = v.next
	}

	return "", false
}

func (m *HashMap) String() string {
	r := &strings.Builder{}

	fmt.Fprintln(r, m.Len(), m.Cap())
	for _, v := range m.buckets {
		for v != nil {
			fmt.Fprintf(r, "\t%s %s", v.key, v.value)
			v = v.next
		}
		fmt.Fprintln(r)
	}

	return r.String()
}

func (m *HashMap) Len() int {
	var total int
	for _, v := range m.buckets {
		for v != nil {
			total++
			v = v.next
		}
	}

	return total
}

func (m *HashMap) Cap() int {
	return len(m.buckets)
}

func (m *HashMap) hash(key string) int {
	hash := int(key[0] - 33)
	for i := 1; i < len(key); i++ {
		hash = hash*hashCode + int(key[i]-33)
	}

	return hash % m.Cap()
}

func (m *HashMap) extend() {
	buckets := m.buckets
	m.buckets = make([]*ListNode, m.Cap()*2)

	for _, v := range buckets {
		if v == nil {
			continue
		}

		for v != nil {
			h := m.hash(v.key)
			next := v.next

			currentNode := m.buckets[h]
			if currentNode == nil {
				v.next = nil
			} else {
				v.next = currentNode
			}

			m.buckets[h] = v
			v = next
		}
	}
}

// ------------------------------

func run(in io.Reader, out io.Writer) {
	var n int
	fmt.Fscan(in, &n)

	m := &HashMap{buckets: make([]*ListNode, 4)}
	for i := 0; i < n; i++ {
		var op string
		fmt.Fscan(in, &op)

		switch op {
		case "put":
			var k, v string
			fmt.Fscan(in, &k, &v)

			m.Put(k, v)
		case "get":
			var k string
			fmt.Fscan(in, &k)

			v, ok := m.Get(k)
			if !ok {
				fmt.Fprintln(out, "-")
			} else {
				fmt.Fprintf(out, "+%s\n", v)
			}
		case "print":
			fmt.Fprint(out, m)
		}
	}
}
