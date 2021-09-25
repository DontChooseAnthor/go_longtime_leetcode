package main

import "fmt"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func (tree *Tree) Add(data int) {
	var queue []*TreeNode
	newNode := &TreeNode{Data: data}
	if tree.root == nil {
		tree.root = newNode
	} else {
		queue = append(queue, tree.root)
		for len(queue) != 0 {
			cur := queue[0]
			queue = append(queue[:0], queue[0+1:]...)
			// 往右树添加
			if data > cur.Data {
				if cur.Right == nil {
					cur.Right = newNode
				} else {
					queue = append(queue, cur.Right)
				}
			} else {
				// 往左树添加
				if cur.Left == nil {
					cur.Left = newNode
				} else {
					queue = append(queue, cur.Left)
				}
			}

		}
	}
	fmt.Println(queue)
}

// 前序遍历
func (tree *Tree) preTraverse(node *TreeNode) {
	if node == nil {
		return
	} else {
		fmt.Println(node.Data, "")
		tree.preTraverse(node.Left)
		tree.preTraverse(node.Right)
	}
}

// 中序遍历
func (tree *Tree) inoTraverse(node *TreeNode) {
	if node == nil {
		return
	} else {
		tree.inoTraverse(node.Left)
		fmt.Println(node.Data, "")
		tree.inoTraverse(node.Right)
	}
}

// 后序遍历
func (tree *Tree) postTraverse(node *TreeNode) {
	if node == nil {
		return
	} else {
		tree.postTraverse(node.Left)
		tree.postTraverse(node.Right)
		fmt.Println(node.Data, "")
	}
}

func main() {
	tree := &Tree{}
	tree.Add(10)
	tree.Add(15)
	tree.Add(34)
	tree.Add(122)
	tree.Add(40)
	tree.Add(50)

	fmt.Println("tree:", &tree)
	tree.preTraverse(tree.root)
	fmt.Printf("\n")
	tree.inoTraverse(tree.root)
	fmt.Printf("\n")
	tree.postTraverse(tree.root)

}
