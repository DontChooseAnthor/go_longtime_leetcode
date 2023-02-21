package main

import "fmt"

type Node struct {
	value  rune
	weight int
	lchild *Node
	rchild *Node
}

type Nodes []Node

type Tree struct {
	Root *Node
}

func sortNodes(nodes Nodes) (sorted []Node) {
	for _, v := range nodes {
		if len(sorted) == 0 {
			sorted = append(sorted, v)
		} else {
			i := 0
			// 找到权重该在的位置
			for i < len(sorted) {
				if v.weight < sorted[i].weight {
					break
				}
				i++
			}
			if i == 0 { // 头插
				temp := []Node{v}
				sorted = append(temp, sorted...)
			} else if i == len(sorted) { // 尾插
				sorted = append(sorted, v)
			} else {
				var pre []Node
				pre = append(pre, sorted[:i]...)

				var after []Node
				after = append(after, sorted[i:]...)

				temp := append(pre, v)
				sorted = append(temp, after...)
			}
		}
	}
	return sorted
}

func makeHuffManTree(nodes Nodes) *Tree {
	var root Node
	for {
		if len(nodes) == 1 {
			break
		} else {
			nodes = sortNodes(nodes)
			pre := nodes[:2]
			after := nodes[2:]
			newNode := Node{}
			newNode.weight = pre[0].weight + pre[1].weight
			newNode.lchild = &pre[0]
			newNode.rchild = &pre[1]

			var temp []Node
			temp = append(temp, newNode)
			nodes = append(temp, after...)
		}
	}

	root = nodes[0]

	tree := &Tree{Root: &root}
	return tree
}

func main() {
	a := Node{
		value:  'A',
		weight: 27,
	}

	b := Node{
		value:  'B',
		weight: 8,
	}
	c := Node{
		value:  'C',
		weight: 15,
	}
	d := Node{
		value:  'D',
		weight: 15,
	}
	e := Node{
		value:  'E',
		weight: 30,
	}

	f := Node{
		value:  'F',
		weight: 5,
	}
	nodes := []Node{a, b, c, d, e, f}
	tree := makeHuffManTree(nodes)
	fmt.Println(tree)
}
