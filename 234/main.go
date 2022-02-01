package main

// Palindrome Linked List
type ListNode struct {
	Val int //
	Next *ListNode //
}
func isPalindrome(head *ListNode) bool {
	dummy := &ListNode{}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		temp := slow.Next
		slow.Next = dummy.Next
		dummy.Next = slow
		slow = temp
	}
	if fast != nil { // odd numbers
		slow = slow.Next
	}
	for p:= dummy.Next; slow != nil; slow, p = slow.Next, p.Next {
		if slow.Val != p.Val {
			return false
		}
	}
	return true
}