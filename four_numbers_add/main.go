package main

import (
	"fmt"
	"sort"
)

func fourSumAdd(nums []int, target int) (quadruplets [][]int) {
	// 1.排序 ,求长度
	sort.Ints(nums)
	n := len(nums)
	// 2.只循环排序后的四数之和 小于 targer的部分
	for i := 0; i < n-3 && nums[i]+nums[i+2]+nums[i+3] <= target; i++ {
		// if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3	]+nums[n-2]+nums[n-1] < target {
		// 	continue
		// }
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			// if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
			// 	continue
			// }
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}

func main() {
	nums := [6]int{1, 0, -1, 0, -2, 2}
	var target int = 0
	list := fourSumAdd(nums[:], target)
	fmt.Println(list)
}
