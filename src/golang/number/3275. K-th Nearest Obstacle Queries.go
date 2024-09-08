package number

import (
	"container/heap"
)

type IntegerHeap []int

// Len 返回堆的长度
func (h IntegerHeap) Len() int {
	return len(h)
}

// Less 比较堆中的两个元素 (i, j) 的大小
// 对于最小堆，我们使用 `h[i] < h[j]`
func (h IntegerHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Swap 交换堆中的两个元素 (i, j)
func (h IntegerHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push 将一个元素添加到堆中
func (h *IntegerHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 从堆中移除并返回最小元素
func (h *IntegerHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func ResultsArray(queries [][]int, k int) []int {
	var l int = len(queries)
	var res []int = make([]int, l)
	h := &IntegerHeap{}
	heap.Init(h)
	for i := 0; i < l; i++ {
		heap.Push(h, abs_int(queries[i][0])+abs_int(queries[i][1]))
		if h.Len() > k {
			heap.Pop(h)
		}
		if h.Len() < k {
			res[i] = -1
		} else {
			res[i] = (*h)[0]
		}
	}
	return res
}
