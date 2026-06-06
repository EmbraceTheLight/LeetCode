// 152. 乘积最大子数组
/*
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续 子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
测试用例的答案是一个 32-位 整数。
请注意，一个只包含一个元素的数组的乘积是这个元素的值。

提示:
1 <= nums.length <= 2 * 10^4
-10 <= nums[i] <= 10
nums 的任何子数组的乘积都 保证 是一个 32 位 整数
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 时间复杂度 O(n^2)
func maxProduct(nums []int) int {
	ans := nums[0]
	dp := make([]int, len(nums)+1) // dp[i] 以 nums[i] 结尾的最大乘积
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		dp[i+1] = max(nums[i], dp[i]*nums[i])
		ans = max(ans, dp[i+1])
	}
	return ans
}

// * 优化: 以 nums[i] 结尾的最大子数组乘积可以分类讨论:
// * 1. nums[i] 单独成子数组
// * 2.1 nums[i] > 0, 则 nums[i] 与 nums[i-1] 结尾的最大子数组相乘即为答案
// * (不易想到)2.2 nums[i] < 0, 则 nums[i] 与 nums[i-1] 结尾的最小子数组相乘即为答案
// * 因此, 除了维护最大子数组乘积外, 还要维护最小子数组乘积
// 时间复杂度 O(n)
func maxProduct2(nums []int) int {
	ans := nums[0]
	dpMax := make([]int, len(nums)) // dpMax[i] 以 nums[i-1] 结尾的最大子数组乘积
	dpMin := make([]int, len(nums)) // dpMin[i] 以 nums[i-1] 结尾的最小子数组乘积
	dpMin[0], dpMax[0] = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			dpMax[i] = max(nums[i], dpMax[i-1]*nums[i])
			dpMin[i] = min(nums[i], dpMin[i-1]*nums[i])
		} else {
			dpMax[i] = max(nums[i], dpMin[i-1]*nums[i])
			dpMin[i] = min(nums[i], dpMax[i-1]*nums[i])
		}
		ans = max(ans, dpMax[i])
	}
	return ans
}

// 示例 1:
// 输入: nums = [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。
//
// 示例 2:
// 输入: nums = [-2,0,-1]
// 输出: 0
// 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(maxProduct(nums))
	fmt.Println(maxProduct2(nums))
}
