package main

// Binary Search Tree Iterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	Stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var s []*TreeNode
	for root != nil {
		s = append(s, root)
		root = root.Left
	}
	return BSTIterator{s}
}

func (this *BSTIterator) Next() int {
	n := len(this.Stack)
	cur := this.Stack[n-1]
	this.Stack = this.Stack[:n-1]
	node := cur.Right
	for node != nil {
		this.Stack = append(this.Stack, node)
		node = node.Left
	}
	return cur.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.Stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
