package main

type Int64Heap []int64

func (h Int64Heap) Len() int {
	return len(h)
}

func (h Int64Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Int64Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Int64Heap) Push(x any) {
	*h = append(*h, x.(int64))
}

func (h *Int64Heap) Pop() any {
	prev := *h
	lastIdx := len(prev) - 1
	res := prev[lastIdx]

	prev = prev[0:lastIdx]
	*h = prev

	return res
}
