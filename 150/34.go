// 34. 在排序数组中查找元素的第一个和最后一个位置
/*
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。


提示：
0 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
nums 是一个非递减数组
-10^9 <= target <= 10^9
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 思路: 先找到一个 target, 之后向两侧作二分, 查找左右边界
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = findLeftBoundary(nums[left:mid+1], left, target)
			right = findRightBoundary(nums[mid:right+1], mid, target)
			return []int{left, right}
		}
	}
	return []int{-1, -1}

}
func findLeftBoundary(nums []int, start int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2

		// 找到左边界
		if mid > 0 && nums[mid-1] < nums[mid] && nums[mid] == target {
			return start + mid
		}

		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] == target {
			r = mid - 1
		}
	}
	return start
}
func findRightBoundary(nums []int, start int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if mid == r {
			return start + r
		}

		// 找到右边界
		if nums[mid+1] > nums[mid] && nums[mid] == target {
			return start + mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] == target {
			l = mid + 1
		}
	}
	return start
}

// 思路: 找到第一个 target 位置, 再找到第一个 target + 1 应出现的位置, 即为所求
func searchRangeBetter(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := search(nums, target)
	right := search(nums, target+1)

	// 未找到 target
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	return []int{left, right - 1}
}

// Test Case1
// nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]
//
// Test Case2
// nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]
//
// Test Case3
// nums = [], target = 0
// Output: [-1,-1]
//
// Test Case4
// nums = [1,2,3], target = 2
// Output: [1,1]
func main() {
	fmt.Println("Input target: ")
	var target int
	fmt.Scan(&target)
	fmt.Println("Input nums: ")
	nums := pkg.CreateSlice[int]()
	fmt.Println(searchRange(nums, target))
	fmt.Println(searchRangeBetter(nums, target))
}
