 package main

 // Reverse Linked List
 type ListNode struct {
	 Val int //
	 Next *ListNode //
 }

 func reverseList1(head *ListNode) *ListNode {
	 dummy := &ListNode{}
	 for head != nil {
		temp := head.Next
		head.Next = dummy.Next
		dummy.Next = head
		head = temp
	 }
	 return dummy.Next
 }

 // recursion
 func reverseList(head *ListNode) *ListNode {
	 if head == nil || head.Next == nil {
		 return head
	 }
	 p := reverseList(head.Next)
	 head.Next.Next = head
	 head.Next = nil
	 return p
 }