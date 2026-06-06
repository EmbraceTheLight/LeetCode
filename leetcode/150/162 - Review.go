// 162. 寻找峰值
/*
峰值元素是指其值严格大于左右相邻值的元素。
给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
你可以假设 nums[-1] = nums[n] = -∞ 。
你必须实现时间复杂度为 O(log n) 的算法来解决此问题。


示例 1：
输入：nums = [1,2,3,1]
输出：2
解释：3 是峰值元素，你的函数应该返回其索引 2。

示例 2：
输入：nums = [1,2,1,3,5,6,4]
输出：1 或 5
解释：你的函数可以返回索引 1，其峰值元素为 2；
     或者返回索引 5， 其峰值元素为 6。

提示：
1 <= nums.length <= 1000
-2^31 <= nums[i] <= 2^31 - 1
对于所有有效的 i 都有 nums[i] != nums[i + 1]
*/

package main

import (
	"fmt"
	"lc/pkg"
)

func findPeakElementR(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	return dfs162R(nums, 0, len(nums)-1)
}
func dfs162R(nums []int, left, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if mid == 0 {
		if nums[0] > nums[1] {
			return mid
		}
	} else if mid == len(nums)-1 {
		if nums[mid] > nums[mid-1] {
			return mid
		}
	} else {
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}
	}

	resL := dfs162R(nums, left, mid-1)
	if resL != -1 {
		return resL
	}
	resR := dfs162R(nums, mid+1, right)
	if resR != -1 {
		return resR
	}
	return -1
}

// 官方题解, 更易理解
// - 当前位置为 mid, 若 nums[mid] < nums[mid - 1], 说明向左走为上坡, 一定有一个峰值, 因此下个区间取 [left, mid - 1]
// - 同理, 若 nums[mid] < nums[mid + 1], 说明向右走为上坡, 一定有一个峰值, 因此下个区间取 [mid + 1, right]
// - 若 nums[mid] > nums[mid - 1] && nums[mid] > nums[mid + 1], 说明当前位置就是峰值, 直接返回 mid
func findPeakElementBest(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if mid == 0 {
			if nums[0] > nums[1] {
				return 0
			}
		} else if mid == len(nums)-1 {
			if nums[mid] > nums[mid-1] {
				return mid
			}
		} else if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else if nums[mid] < nums[mid-1] {
			right = mid - 1
		}
	}
	return -1
}

// Test Case1:
// nums = [1,2,3,1]
// Output: 2

// Test Case2:
// nums = [1,2,1,3,5,6,4]
// Output: 1 或 5

// Test Case3:
// nums = [1,2]
// Output: 1
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(findPeakElementR(nums))
	fmt.Println(findPeakElementBest(nums))
}
