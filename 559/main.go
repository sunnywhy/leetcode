package main

// Maximum Depth of N-ary Tree

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var maxChild int
	for _, v := range root.Children {
		cur := maxDepth(v)
		if cur > maxChild {
			maxChild = cur
		}
	}
	return maxChild + 1
}
