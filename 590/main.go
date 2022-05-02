package main

// N-ary Tree Postorder Traversal

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var res []int
	traversal(root, &res)
	return res
}
func traversal(root *Node, res *[]int) {
	if root == nil {
		return
	}
	for _, child := range root.Children {
		if child != nil {
			traversal(child, res)
		}
	}
	*res = append(*res, root.Val)
}

func postorderIteratively(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	var s []*Node
	s = append(s, root)
	for len(s) > 0 {
		cur := s[len(s)-1]
		s = s[:len(s)-1]
		res = append([]int{cur.Val}, res...)
		for _, child := range cur.Children {
			if child != nil {
				s = append(s, child)
			}
		}
	}
	return res
}
