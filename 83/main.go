package main

import "math"

// 删除排序链表中的重复元素
type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: math.MinInt, Next: head}
	prev, cur := dummy, head
	for cur != nil {
		if prev.Val == cur.Val {
			cur = cur.Next
			continue
		}
		prev.Next = cur
		prev = cur
	}
	prev.Next = nil // handle duplicate in the end
	return dummy.Next
}