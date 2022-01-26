package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	p := dummy
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			continue
		}
		p = p.Next
	}
	return dummy.Next
}

func main() {
	data := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
			},
		},
	}
	result := removeElements(data, 6)
	for result != nil {
		fmt.Println(result)
		result = result.Next
	}
}