package main

import (
	"fmt"
	"math/rand"
	"time"
)

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	rand.Seed(time.Now().UnixNano())
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right + rand.Intn(2)) / 2
	root := &TreeNode{Data: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}

func main() {
	a := []int{-1, 1, 3}
	b := sortedArrayToBST(a)
	fmt.Println(*b)
}
