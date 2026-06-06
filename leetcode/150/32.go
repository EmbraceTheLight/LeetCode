// 35. 搜索插入位置
/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。



示例 1:
输入: nums = [1,3,5,6], target = 5
输出: 2

示例 2:
输入: nums = [1,3,5,6], target = 2
输出: 1

示例 3:
输入: nums = [1,3,5,6], target = 7
输出: 4

提示:
1 <= nums.length <= 10^4
-10^4 <= nums[i] <= 10^4
nums 为 无重复元素 的 升序 排列数组
-10^4 <= target <= 10^4
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = right - 1
		}
	}
	if left >= len(nums) {
		return left
	}
	return mid
}

// Test Case1: nums = [1,3,5,6], target = 5		Output: 2
// Test Case2: nums = [1,3,5,6], target = 2		Output: 1
// Test Case3: nums = [1,3,5,6], target = 7		Output: 4
func main() {
	var target int
	fmt.Println("Input target:")
	fmt.Scan(&target)
	fmt.Println("Input nums:")
	nums := pkg.CreateSlice[int]()
	fmt.Println(searchInsert(nums, target))
}
