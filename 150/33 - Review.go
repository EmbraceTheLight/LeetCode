// 33. 搜索旋转排序数组
/*
整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 向左旋转，
使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
例如， [0,1,2,4,5,6,7] 下标 3 上向左旋转后可能变为 [4,5,6,7,0,1,2] 。
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。


示例 1：

输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

示例 2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

示例 3：
输入：nums = [1], target = 0
输出：-1

提示：
1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
nums 中的每个值都 独一无二
题目数据保证 nums 在预先未知的某个下标上进行了旋转
-10^4 <= target <= 10^4
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func searchR(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	peakIdx := findPeak(nums)
	resLeft := searchInsert33R(nums[:peakIdx+1], target)
	if resLeft != -1 {
		return resLeft
	}
	resRight := searchInsert33R(nums[peakIdx+1:], target)
	if resRight != -1 {
		return peakIdx + 1 + resRight
	}
	return -1
}
func searchInsert33R(nums []int, target int) int {
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

	return -1
}

// findPeak 寻找峰值, 即旋转后的 nums[n - 1] 所在位置
func findPeak(nums []int) int {
	left, right := 0, len(nums)

	for left <= right {
		mid := (left + right) / 2
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

		leftLen := mid - left + 1 // 左半部分 nums[left,mid] 数组长度
		//rightLen := right - mid - 1 // 右半部分 nums[mid + 1, right] 数组长度

		// 由于题目保证数组中的数各不相等, 因此可以根据左半部分/右半部分的起点和终点的大小差确定峰值是否在该部分数组中
		// - 左半部分终点 - 起点 < 左半部分长度 - 1 --> 峰值在左半部分
		// - 右半部分终点 - 起点 < 右半部分长度 - 1 --> 峰值在右半部分
		// 以上条件能且只能同时满足一个
		if nums[mid]-nums[left] < leftLen-1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// Test Case1:
// nums = [4,5,6,7,0,1,2], target = 0
// Output: 4

// Test Case2:
// nums = [4,5,6,7,0,1,2], target = 3
// Output: -1

// Test Case3：
// nums = [1], target = 0
// Output: -1
func main() {
	fmt.Println("Input target:")
	var target int
	fmt.Scan(&target)
	fmt.Println("Input nums:")
	nums := pkg.CreateSlice[int]()
	fmt.Println(searchR(nums, target))
}
