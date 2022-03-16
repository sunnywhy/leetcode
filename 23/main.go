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
func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	v := (*h)[n-1]
	*h = (*h)[:n-1]
	return v
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	h := &minHeap{}
	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}

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
