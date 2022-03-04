package main

// Nested List Weight Sum

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedIntegerType int

const (
	TYPE_INTEGER NestedIntegerType = 0
	TYPE_LIST    NestedIntegerType = 1
)

type NestedInteger struct {
	Type NestedIntegerType
	Val  int
	List []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool { return n.Type == TYPE_INTEGER }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int { return n.Val }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.Val = value
	n.Type = TYPE_INTEGER
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	n.List = append(n.List, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger { return n.List }

func depthSum(nestedList []*NestedInteger) int {
	return calculate(nestedList, 1)
}

func calculate(nestedList []*NestedInteger, level int) int {
	var total int
	for _, v := range nestedList {
		if v.IsInteger() {
			total += v.GetInteger()
		} else {
			total += calculate(v.GetList(), level+1)
		}
	}
	return total
}
