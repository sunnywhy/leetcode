package main

// Convert Sorted List to Binary Search Tree
// https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Method 1: Recursion, O(N*logN) time, O(1) space
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	rightStart := slow.Next
	slow.Next = nil
	prev.Next = nil
	left := sortedListToBST(head)
	right := sortedListToBST(rightStart)
	return &TreeNode{slow.Val, left, right}
}

// Method 2: Convert to pure data first, O(N) time, O(N) space
func sortedListToBST2(head *ListNode) *TreeNode {
	var data []int
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}
	return buildTree(data, 0, len(data))
}
func buildTree(data []int, start, end int) *TreeNode {
	if start >= end {
		return nil
	}
	mid := start + (end-start)/2
	left := buildTree(data, start, mid)
	right := buildTree(data, mid+1, end)
	return &TreeNode{data[mid], left, right}
}
