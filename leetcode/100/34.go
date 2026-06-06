/*
34. 在排序数组中查找元素的第一个和最后一个位置
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

示例 3：
输入：nums = [], target = 0
输出：[-1,-1]
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func findLeft(nums []int, l, r, target int) int {
	var ret = r
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			ret = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ret
}
func findRight(nums []int, l, r, target int) int {
	var ret = l
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			ret = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ret
}
func searchRange(nums []int, target int) []int {
	var ans = make([]int, 0)
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			ans = append(ans, findLeft(nums, l, mid, target))
			ans = append(ans, findRight(nums, mid, r, target))
			//left := mid
			//right := mid
			//for left >= l && nums[left] == target {
			//	left--
			//}
			//for right <= r && nums[right] == target {
			//	right++
			//}
			//ans = append(ans, left+1)
			//ans = append(ans, right-1)
			break
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if len(ans) == 0 {
		ans = append(ans, -1)
		ans = append(ans, -1)
	}
	return ans
}
func main() {
	fmt.Println("Input target:")
	var target int
	fmt.Scan(&target)
	nums := pkg.CreateIntSlice()
	fmt.Println(searchRange(nums, target))

}
