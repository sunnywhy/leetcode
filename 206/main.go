 package main

 // Reverse Linked List
 type ListNode struct {
	 Val int //
	 Next *ListNode //
 }

 func reverseList(head *ListNode) *ListNode {
	 dummy := &ListNode{}
	 for head != nil {
		temp := head.Next
		head.Next = dummy.Next
		dummy.Next = head
		head = temp
	 }
	 return dummy.Next
 }