package main

import (
	"container/heap"
	"sort"
)

type hp []int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
func (h *hp) push(v int) { heap.Push(h, v) }
func (h *hp) pop() int   { return heap.Pop(h).(int) }

func isPossible(nums []int) bool {
	lens := map[int]*hp{}
	for _, v := range nums {
		if lens[v] == nil {
			lens[v] = new(hp)
		}
		if h, ok := lens[v-1]; ok {
			prevLen := h.pop()
			if h.Len() == 0 {
				delete(lens, v-1)
			}
			lens[v].push(prevLen + 1)
		} else {
			lens[v].push(1)
		}
	}

	for _, h := range lens {
		if h.pop() < 3 {
			return false
		}
	}
	return true
}
