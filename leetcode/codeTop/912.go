// 912. 排序数组
/*
给你一个整数数组 nums，请你将该数组升序排列。
你必须在 不使用任何内置函数 的情况下解决问题，时间复杂度为 O(nlog(n))，并且空间复杂度尽可能小。

提示：
1 <= nums.length <= 5 * 10^4
-5 * 10^4 <= nums[i] <= 5 * 10^4
*/
package main

import (
	"fmt"
	"lc/pkg"
	"math/rand"
)

func partition912(nums []int, left, right int) int {
	// 引入随机化基准值, 而不是以第一个/最后一个作为基准值, 防止最坏情况导致退化到 O(n^2)
	// 某些情况下, 随机化基准只需要 30ms, 而固定基准值则需要 600ms+, 差异显著
	pivot := nums[left+rand.Intn(right-left+1)]
	l, r := left-1, right+1
	for l < r {
		for {
			l++
			if nums[l] >= pivot {
				break
			}
		}
		for {
			r--
			if nums[r] <= pivot {
				break
			}
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	return r
}

func quickSort912(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition912(nums, left, right)
	quickSort912(nums, left, pivot)
	quickSort912(nums, pivot+1, right)
}
func sortArray(nums []int) []int {
	quickSort912(nums, 0, len(nums)-1)
	return nums
}

// 示例 1：
// 输入：nums = [5,2,3,1]
// 输出：[1,2,3,5]
// 解释：数组排序后，某些数字的位置没有改变（例如，2 和 3），而其他数字的位置发生了改变（例如，1 和 5）。
//
// 示例 2：
// 输入：nums = [5,1,1,2,0,0]
// 输出：[0,0,1,1,2,5]
// 解释：请注意，nums 的值不一定唯一。
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(sortArray(nums))
}
