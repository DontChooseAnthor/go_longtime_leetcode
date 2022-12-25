package main

import "fmt"

type NodeList struct {
	Value int
	Next  *NodeList
}

func hasCycle(head *NodeList) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	quick := head
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next

		if quick == slow {
			return true
		}
	}
	return false
}

func main() {
	node5 := NodeList{5, nil}
	node4 := NodeList{4, &node5}
	node3 := NodeList{3, &node4}
	node2 := NodeList{2, &node3}
	node1 := NodeList{1, &node2}
	node5.Next = &node3
	fmt.Println(hasCycle(&node1))
}
