package main

import "fmt"

func selectSort(foo []int) {
	var n int = len(foo)
	for i := 0; i < n-1; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			if foo[k] > foo[j] {
				k = j
			}
		}
		if k != i {
			foo[k], foo[i] = foo[i], foo[k]
		}
	}
}

func main() {
	var arr []int = []int{0, 2, 5, 1, 4, 3}
	selectSort(arr)
	fmt.Println(arr)
}
