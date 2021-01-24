package main

// Runtime: 96 ms, faster than 99.07% of Go online submissions for Find Median from Data Stream.
// Memory Usage: 16.3 MB, less than 76.64% of Go online submissions for Find Median from Data Stream.
// https://leetcode.com/submissions/detail/447202722/

type heap struct {
	data []int
	less func(i, j int) bool
}

func (h *heap) Push(val int) {
	h.data = append(h.data, val)
	index := len(h.data) - 1
	for parent := (index - 1) / 2; index > 0 && h.less(h.data[index], h.data[parent]); index, parent = parent, (parent-1)/2 {
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
	}
}

func (h *heap) Pop() int {
	last := len(h.data) - 1
	val := h.data[0]
	h.data[0], h.data[last] = h.data[last], h.data[0]
	h.data = h.data[:last]
	for index, left := 0, 1; left < last; left = 2*index + 1 {
		min, right := left, left+1
		if right < last && h.less(h.data[right], h.data[left]) {
			min = right
		}
		if !h.less(h.data[min], h.data[index]) {
			break
		}
		h.data[index], h.data[min] = h.data[min], h.data[index]
		index = min
	}
	return val
}

// MedianFinder ...
type MedianFinder struct {
	minHeap, maxHeap heap
}

// Constructor : initialize your data structure here.
func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: heap{less: func(i, j int) bool { return i < j }},
		maxHeap: heap{less: func(i, j int) bool { return i > j }},
	}
}

// AddNum ...
func (mf *MedianFinder) AddNum(num int) {
	if len(mf.minHeap.data) < len(mf.maxHeap.data) {
		mf.maxHeap.Push(num)
		mf.minHeap.Push(mf.maxHeap.Pop())
	} else {
		mf.minHeap.Push(num)
		mf.maxHeap.Push(mf.minHeap.Pop())
	}
}

// FindMedian ...
func (mf *MedianFinder) FindMedian() float64 {
	if len(mf.minHeap.data) == len(mf.maxHeap.data) {
		return (float64(mf.minHeap.data[0]) + float64(mf.maxHeap.data[0])) / 2
	}
	return float64(mf.maxHeap.data[0])
}
