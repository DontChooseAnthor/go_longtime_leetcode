package main

import "fmt"

func bubbleSort(foo []int) (bar []int) {
	bar = foo
	n := len(bar)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			var temp = bar[i]
			if bar[i] > bar[j] {
				bar[i] = bar[j]
				bar[j] = temp
			}
		}
	}
	return bar
}

func main() {
	arr := []int{4, 2, 1, -1, 6, 4}
	brr := bubbleSort(arr)
	fmt.Println(brr)
}
