package main

import (
	"math"
	"sort"
)

// Binary Tree Vertical Order Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type data struct {
	Row int
	Col int
	Val int
}

var list []data

// method one, get all data and sort, O(NlogN)
func verticalOrder(root *TreeNode) [][]int {
	list = []data{}
	traversal(root, 0, 0)
	sort.SliceStable(list, func(i, j int) bool {
		if list[i].Col == list[j].Col {
			return list[i].Row < list[j].Row
		} else {
			return list[i].Col < list[j].Col
		}
	})
	var res [][]int
	var cols []int
	prevCol := math.MinInt
	for _, v := range list {
		if v.Col > prevCol {
			if prevCol != math.MinInt {
				res = append(res, append([]int{}, cols...))
			}
			cols = []int{}
			prevCol = v.Col
		}
		cols = append(cols, v.Val)
	}
	if prevCol != math.MinInt {
		res = append(res, cols)
	}
	return res
}
func traversal(root *TreeNode, row, col int) {
	if root == nil {
		return
	}
	list = append(list, data{row, col, root.Val})
	traversal(root.Left, row+1, col-1)
	traversal(root.Right, row+1, col+1)
}

// method two, save each col as a sub list, only sort the sub list by row
// build all sub list -> O(N), sort, consider every number of element is M in the sub list, O(N/M * MlogM) = O(NlogM)
type pair struct {
	Row int
	Val int
}

func verticalOrder2(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	minCol, maxCol := 0, 0
	data := make(map[int][]pair)
	traversalWithMap(root, 0, 0, data, &minCol, &maxCol)
	for i := minCol; i <= maxCol; i++ {
		values := data[i]
		sort.SliceStable(values, func(i, j int) bool {
			return values[i].Row < values[j].Row
		})
		var cur []int
		for _, v := range values {
			cur = append(cur, v.Val)
		}
		res = append(res, cur)
	}
	return res
}
func traversalWithMap(root *TreeNode, row, col int, data map[int][]pair, minCol, maxCol *int) {
	if root == nil {
		return
	}
	data[col] = append(data[col], pair{row, root.Val})
	*minCol = int(math.Min(float64(*minCol), float64(col)))
	*maxCol = int(math.Max(float64(*maxCol), float64(col)))
	traversalWithMap(root.Left, row+1, col-1, data, minCol, maxCol)
	traversalWithMap(root.Right, row+1, col+1, data, minCol, maxCol)
}
