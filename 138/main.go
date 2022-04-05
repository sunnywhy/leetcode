package main

// Copy List with Random Pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	dataMap := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		newNode := &Node{Val: cur.Val}
		dataMap[cur] = newNode
		cur = cur.Next
	}

	for original, copied := range dataMap {
		copied.Next = dataMap[original.Next]
		copied.Random = dataMap[original.Random]
	}
	return dataMap[head]
}
