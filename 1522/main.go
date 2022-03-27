package main

// Diameter of N-Ary Tree
// https://leetcode.com/problems/diameter-of-n-ary-tree/

type Node struct {
	Val      int
	Children []*Node
}

var res int

func diameter(root *Node) int {
	res = 0
	checkAndCalculate(root)
	return res
}
func checkAndCalculate(root *Node) int {
	if root == nil {
		return 0
	}
	var first, second int
	for _, child := range root.Children {
		cur := checkAndCalculate(child)
		if first < cur {
			second = first
			first = cur
		} else if second < cur {
			second = cur
		}
	}
	if res < first+second {
		res = first + second
	}
	return first + 1
}
