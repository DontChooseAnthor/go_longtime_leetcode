package main

import "fmt"

var result [][]int

/*
*回溯算法核心：
*nums:原始列表
*pathNums:路径数字
*used:是否访问过
 */
func backtrack(nums, pathNums []int, used []bool) {
	// 结束条件：走完了，路径上的数字长度等于原始列表长度
	if len(nums) == len(pathNums) {
		tmp := make([]int, len(nums))
		// 切片底层公用数据
		copy(tmp, pathNums)
		// 把本次结果追加到最终结果上
		result = append(result, tmp)
		return
	}

	// 遍历原始数组
	for i := 0; i < len(nums); i++ {
		// 检查是否访问过
		if !used[i] {
			used[i] = true
			// 做选择：将这个数字加入路径的尾部，这里用数组模拟链表
			pathNums = append(pathNums, nums[i])
			backtrack(nums, pathNums, used)
			// 撤销刚才的选择，恢复操作
			pathNums = pathNums[:len(pathNums)-1]
			// 修改标记为未使用
			used[i] = false
		}
	}
}

func permute(nums []int) [][]int {
	var pathNums []int
	var used = make([]bool, len(nums))
	// 清空全局数组
	result = [][]int{}
	backtrack(nums, pathNums, used)
	return result
}

func main() {
	arr := []int{0, 1, 2}
	a := permute(arr)
	fmt.Println(a)
}
