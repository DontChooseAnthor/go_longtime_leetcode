package main

import "fmt"

func insertSort(foo []int) {
	var n int = len(foo)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && foo[j] < foo[j-1]; j-- {
			foo[j], foo[j-1] = foo[j-1], foo[j]
		}
	}
}

func main() {
	var arr []int = []int{5, 7, 3, 2, 6, 1, 4}
	insertSort(arr)
	fmt.Println(arr)
}
