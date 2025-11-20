// 1. 两数之和
/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]

提示：

2 <= nums.length <= 10^4
-10^9 <= nums[i] <= 10^9
-10^9 <= target <= 10^9
只会存在一个有效答案

进阶：你可以想出一个时间复杂度小于 O(n^2) 的算法吗？
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 新理解：如果 nums 中有多个相同元素（如 3），target 恰好为这两个相同元素之和，不会影响结果。
// 因为 mp 最后存储的是两个相同元素靠后的那个元素的 index，而遍历 nums 时，会首先遇到第一个相同元素，然后再在 mp 中找到另一个相同元素的 index，
func twoSum1(nums []int, target int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := mp[target-nums[i]]; ok && mp[target-nums[i]] != i {
			return []int{i, mp[target-nums[i]]}
		}
	}
	return []int{}
}

// Test Case1: [2,7,11,15], target = 9	Output: [0,1]
// Test Case2: [3,2,4], target = 6	Output: [1,2]
// Test Case3: [3,3], target = 6	Output: [0,1]
// Test Case2: [3,2,3,4], target = 6	Output: [0,2]
func main() {
	var target int
	fmt.Println("Input target:")
	fmt.Scan(&target)

	nums := pkg.CreateSlice[int]()
	fmt.Println(twoSum1(nums, target))
}
