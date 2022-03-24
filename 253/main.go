package main

import (
	"container/heap"
	"sort"
)

// Meeting Rooms II

type MinHeap [][]int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1] // sort by the meeting end time
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
func (h *MinHeap) Pop() interface{} {
	n := h.Len()
	val := (*h)[n-1]
	*h = (*h)[:n-1]
	return val
}
func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	h := &MinHeap{}
	for _, v := range intervals {
		if h.Len() != 0 && (*h)[0][1] <= v[0] {
			heap.Pop(h)
		}
		heap.Push(h, []int{v[0], v[1]})
	}
	return h.Len()
}

func minMeetingRooms2(intervals [][]int) int {
	n := len(intervals)
	starts := make([]int, n)
	ends := make([]int, n)
	for i, v := range intervals {
		starts[i] = v[0]
		ends[i] = v[1]
	}

	sort.Ints(starts)
	sort.Ints(ends)
	var res, endIdx int
	for _, start := range starts {
		if start < ends[endIdx] {
			res++
		} else {
			endIdx++
		}
	}
	return res
}
