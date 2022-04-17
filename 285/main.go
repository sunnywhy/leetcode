package main

// Inorder Successor in BST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var prev *TreeNode

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if p.Right != nil { // find the left most node in the right tree
		return findMostLeftNode(p.Right)
	} else {
		prev = nil
		checkPathForNode(root, p)
		return prev
	}
}

func findMostLeftNode(node *TreeNode) *TreeNode {
	cur := node
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

func checkPathForNode(root *TreeNode, p *TreeNode) {
	cur := root
	for cur != p {
		if p.Val < cur.Val {
			prev = cur
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
}
