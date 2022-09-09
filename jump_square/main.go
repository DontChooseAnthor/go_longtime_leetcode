package main

import "fmt"

type Tree struct {
	Value int32
	Left  *Tree
	Right *Tree
}

func BuildTree(i int32, t *Tree) *Tree {
	if i >= 0 {
		t.Value = i
		t1 := new(Tree)
		t2 := new(Tree)
		if i-2 >= 0 {
			t.Left = BuildTree(i-2, t1)
		}
		if i-1 >= 0 {
			t.Right = BuildTree(i-1, t2)
		}
	}
	return t
}

var output int32 = 0

func trave(t *Tree) {
	if t.Value > 0 {
		if t.Left != nil {
			trave(t.Left)
		}
		if t.Right != nil {
			trave(t.Right)
		}
		output = output + 0
	} else if t.Value < 0 {
		output = output + 0
	} else {
		output = output + 1
	}
}

func main() {
	var input int32 = 10
	t := new(Tree)
	BuildTree(input, t)
	trave(t)
	fmt.Println(output)
}
