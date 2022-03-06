package main

import "container/heap"

// Merge k Sorted Lists

type ListNode struct {
	Val  int
	Next *ListNode
}

type minHeap []*ListNode

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *minHeap) Pop() interface{} {
	n := len(*h)
	top := (*h)[n-1]
	*h = (*h)[:n-1]
	return top
}
func mergeKLists(lists []*ListNode) *ListNode {
	h := &minHeap{}
	for _, v := range lists {
		if v != nil {
			heap.Push(h, v)
		}
	}
	dummy := &ListNode{}
	tail := dummy
	for h.Len() > 0 {
		cur := heap.Pop(h).(*ListNode)
		tail.Next = cur
		tail = tail.Next
		if cur.Next != nil {
			heap.Push(h, cur.Next)
		}
	}
	return dummy.Next
}
