package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := ListNode{}

	for head != nil {
		cur := &dummy
		for cur.Next != nil && cur.Next.Val < head.Val {
			cur = cur.Next
		}
		temp := head.Next
		head.Next = cur.Next
		cur.Next = head
		head = temp
	}
	return dummy.Next
}
