package main

import "fmt"

/*
*
@desc 升序数组 删除重复元素，返回删除后的长度，元素顺序保持不变
@solution 将数组转化为链表，双指针进行遍历删除操作
*/
type Node struct {
	prev *Node
	next *Node
	data int
}

func NewListNode(data int) *Node {
	return &Node{
		data: data,
	}
}

func (nodes *Node) Push(data int) error {
	currentNode := nodes
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = NewListNode(data)
	currentNode.next.prev = currentNode
	return nil
}

func (nodes *Node) Read() {
	res := make([]int, 0)
	currentNode := nodes
	res = append(res, currentNode.data)
	defer func() {
		fmt.Println("currentNode", res)
	}()
deal:
	if currentNode.next != nil {
		currentNode = currentNode.next
		res = append(res, currentNode.data)
		goto deal
	}
}

func DealLoop(node *Node) {
	if node.next == nil {
		return
	} else {
		if node.data == node.next.data {
			node.next = node.next.next
			DealLoop(node)
		} else {
			DealLoop(node.next)
		}
	}

}

func main() {
	arr := []int{1, 2, 2, 3, 3, 3, 4, 5, 6, 6, 7}
	nodeList := NewListNode(arr[0])
	for _, v := range arr {
		nodeList.Push(v)
	}
	DealLoop(nodeList)
	nodeList.Read()
}
