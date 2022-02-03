package main

// Odd Even Linked List
type ListNode struct {
	Val int 
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyOdd, dummyEven := &ListNode{}, &ListNode{}
	odd, even := dummyOdd, dummyEven
	var idx int
	for cur := head; cur != nil; {
		temp := cur.Next
		cur.Next = nil
		if idx % 2 == 0 {
			odd.Next = cur
			odd = odd.Next
		}else {
			even.Next = cur
			even = even.Next
		}
		idx++
		cur = temp
	}
	odd.Next = dummyEven.Next
	return dummyOdd.Next
}