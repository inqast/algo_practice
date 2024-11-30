package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type task struct {
	id       int
	priority int
}

func (left task) isMoreImportant(right task) bool {
	if left.priority != right.priority {
		return left.priority < right.priority
	}

	return left.id < right.id
}

type heap []task

func (h heap) siftUp() {
	idx := len(h) - 1
	for {
		if h[idx].isMoreImportant(h[idx/2]) && idx != 1 {
			h[idx/2], h[idx] = h[idx], h[idx/2]
			idx /= 2
			continue
		}

		break
	}
}

func (h heap) siftDown() {
	idx := 1

	for {
		left := idx * 2
		right := idx*2 + 1

		if left >= len(h) {
			return
		}

		var maxIdx int
		if right < len(h) && h[right].isMoreImportant(h[left]) {
			maxIdx = right
		} else {
			maxIdx = left
		}

		if h[maxIdx].isMoreImportant(h[idx]) {
			h[idx], h[maxIdx] = h[maxIdx], h[idx]
			idx = maxIdx
		} else {
			return
		}
	}
}

func newHeap() heap {
	return make(heap, 1)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var commandsCount int
	fmt.Fscan(in, &commandsCount)

	h := newHeap()

	for commandIdx := 0; commandIdx < commandsCount; commandIdx++ {
		var id int
		fmt.Fscan(in, &id)

		if id < 0 {
			h[1] = h[len(h)-1]
			h = h[:len(h)-1]

			h.siftDown()
		} else {
			var priority int
			fmt.Fscan(in, &priority)

			h = append(h, task{
				id:       id,
				priority: priority,
			})
			h.siftUp()
		}
	}

	for len(h) > 1 {
		extremum := h[1]

		h[1] = h[len(h)-1]
		h = h[:len(h)-1]

		h.siftDown()

		writer.WriteString(strconv.Itoa(extremum.id) + " ")
	}
}
