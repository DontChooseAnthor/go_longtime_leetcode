package main

// func mergeSort(arr []int) []int {
// 	n := len(arr)
// 	if n < 2 {
// 		return arr
// 	}
// 	index := n / 2
// 	left := mergeSort(arr[0:index])
// 	right := mergeSort(arr[index:])
// 	return merge(left, right)
// }

// func merge(left []int, right []int) []int {
// 	tmp := make([]int, 0)
// 	i, j := 0, 0
// }

func main() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
}
