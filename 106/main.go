package main

// Construct Binary Tree from Inorder and Postorder Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var inorderCache map[int]int

func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderCache = make(map[int]int)
	for i, v := range inorder {
		inorderCache[v] = i
	}
	return build(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}
func build(inorder []int, inStart, inEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	val := postorder[postEnd]
	inorderPos := inorderCache[val]
	leftLen := inorderPos - inStart
	node := TreeNode{Val: val}
	node.Left = build(inorder, inStart, inorderPos-1, postorder, postStart, postStart+leftLen-1)
	node.Right = build(inorder, inorderPos+1, inEnd, postorder, postStart+leftLen, postEnd-1)
	return &node
}
