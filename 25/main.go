package main

// Reverse Nodes in k-Group
type ListNode struct {
	Val int 
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
    for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	
	// reverse first k
	dummy := &ListNode{}
	tail := head
	for i := 0; i < k; i++ {
		temp := head.Next
		head.Next = dummy.Next
		dummy.Next = head
		head = temp
	}
	tail.Next = reverseKGroup(head, k)
	return dummy.Next
}