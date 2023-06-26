package main

import (
	"errors"
	"fmt"
	"sync"
)

/**
树上有n个节点以0 - n-1排序
实现TreeAncestor类：
	- TreeAncestor(int n, int[] parent) 对树和父组件中的节点数初始化对象
		- TreeAncestor(7, [-1, 0,0 ,1,1, 2,2])
		- 树结构：0-(12) 1-(34) 2-(56)
	- getKthAncestor(int node, int k)返回节点node的第k个祖先节点，如果不存在返回-1
*/

type Node struct {
	pre   *Node
	value int
	LNext *Node
	RNext *Node
}

type TreeNodeDescripe struct {
	desc   int
	capArr []int
}

var wg sync.WaitGroup

// build value:level map for AddNode func
func TreeAncestor(parent []int) (map[int]int, error) {
	if len(parent) <= 0 {
		return nil, errors.New("parent not legal")
	}

	mapTreeDesc := make(map[int]int)
	for i, v := range parent {
		mapTreeDesc[i] = v
	}
	return mapTreeDesc, nil
}

// make root node: value = 0
func InitNodeList() *Node {
	return &Node{
		pre:   nil,
		value: 0,
		LNext: nil,
		RNext: nil,
	}
}

func (n *Node) AddNode(value int, level int) error {
	if level < 0 {
		return errors.New("err: get root level")
	}
	// level = node.next...
	currentLevel := 0
	currentNode := n

	if currentLevel < level {
		
	}

	wg.Done()
}

func main() {
	parent := []int{-1, 0, 0, 1, 1, 1, 1}
	n := len(parent)
	treeList, err := TreeAncestor(parent)
	if err != nil {
		fmt.Println("error: get 0 treeList")
		return
	}
	nodeList := InitNodeList()
	wg.Add(n)
	for key, _ := range treeList {
		go nodeList.AddNode(key, treeList[key])
	}
	wg.Wait()
}
