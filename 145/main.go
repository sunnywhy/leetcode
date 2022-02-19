package main

// Binary Tree Postorder Traversal
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var res []int
	postorder(root, &res)
	return res
}

func postorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, res)
	postorder(root.Right, res)
	*res = append(*res, root.Val)
}

func postorderTraversalIteratively(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	s := []*TreeNode{root}
	for len(s) > 0 {
		cur := s[len(s)-1]
		s = s[:len(s)-1]
		res = append([]int{cur.Val}, res...)
		if cur.Left != nil {
			s = append(s, cur.Left)
		}
		if cur.Right != nil {
			s = append(s, cur.Right)
		}
	}
	return res
}
