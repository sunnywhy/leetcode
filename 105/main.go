package main

// Construct Binary Tree from Preorder and Inorder Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var inorderCache map[int]int

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderCache = make(map[int]int)
	for i, v := range inorder {
		inorderCache[v] = i
	}
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	val := preorder[preStart]
	node := TreeNode{Val: val}
	inorderPos := inorderCache[val]
	leftLength := inorderPos - inStart
	node.Left = build(preorder, preStart+1, preStart+leftLength, inorder, inStart, inorderPos-1)
	node.Right = build(preorder, preStart+leftLength+1, preEnd, inorder, inorderPos+1, inEnd)
	return &node
}
