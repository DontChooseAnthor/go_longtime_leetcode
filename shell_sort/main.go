package main

import "fmt"

// 希尔排序
// 固定一个步长，数组根据这个步长划分成n个二个元素的数组，每个数组进行插排
// 不断缩小步长，重复以上过程
func shellSortStep(arr []int, start int, gap int) []int {
	for i := start + gap; i < len(arr); i += gap {
		backup := arr[i]
		j := i - gap
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = backup
	}
	return arr
}
func shellSort(arr []int) []int {
	gap := len(arr) / 2
	for gap > 0 {
		for i := 0; i < gap; i++ {
			shellSortStep(arr, i, gap)
		}
		gap /= 2
	}
	return arr
}

func main() {
	var arr []int = []int{5, 3, 8, 9, 1, 6, 2, 7, 4, 0}
	fmt.Println(shellSort(arr))
}
