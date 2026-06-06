// 41. 缺失的第一个正数
/*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

提示：
1 <= nums.length <= 10^5
-2^31 <= nums[i] <= 2^31 - 1
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func firstMissingPositive(nums []int) int {
	var ans = 1
	mp := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			mp[nums[i]] = true
		}
	}
	for ; mp[ans] == true; ans++ {
	}
	return ans
}
func firstMissingPositive2(nums []int) int {
	var ans = len(nums) + 1 // 初始默认 nums 中包含 1, 2, ... , len(nums), 此时最小正数是 len(nums) + 1
	// 原地哈希. nums[i] = i + 1.
	n := len(nums)
	// 原地哈希. nums[i] = i + 1.
	// nums[i] >= n 或 nums[i] <= 0, 直接跳过
	for i := 0; i < n; i++ {
		if nums[i] >= n || nums[i] <= 0 || nums[i] == i+1 || nums[i] == nums[nums[i]-1] {
			continue
		}
		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		i--
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			ans = i + 1
			break
		}
	}
	return ans
}

// 示例 1：
// 输入：nums = [1,2,0]
// 输出：3
// 解释：范围 [1,2] 中的数字都在数组中。
//
// 示例 2：
// 输入：nums = [3,4,-1,1]
// 输出：2
// 解释：1 在数组中，但 2 没有。
//
// 示例 3：
// 输入：nums = [7,8,9,11,12]
// 输出：1
// 解释：最小的正数 1 没有出现。
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(firstMissingPositive(nums))
	fmt.Println(firstMissingPositive2(nums))
}
