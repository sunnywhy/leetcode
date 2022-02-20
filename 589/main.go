package main

// N-ary Tree Preorder Traversal
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var res []int
	traversal(root, &res)
	return res
}

func traversal(root *Node, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	for _, child := range root.Children {
		traversal(child, res)
	}
}

func preorder_interatively(root *Node) []int {
	var res []int
	var s []*Node
	if root == nil {
		return res
	}
	s = append(s, root)
	for len(s) > 0 {
		cur := s[len(s)-1]
		s = s[:len(s)-1]
		res = append(res, cur.Val)
		for i := len(cur.Children) - 1; i >= 0; i-- {
			child := cur.Children[i]
			if child == nil {
				continue
			}
			s = append(s, child)
		}
	}
	return res
}
