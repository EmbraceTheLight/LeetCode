// 53. 最大子数组和
/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组是数组中的一个连续部分。

示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6。

示例 2：
输入：nums = [1]
输出：1

示例 3：
输入：nums = [5,4,-1,7,8]
输出：23

提示：
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func maxSubArrayR(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := dp[0] // dp[i]: 以 nums[i] 结尾的子数组最大和
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		ans = max(ans, dp[i])
	}
	return ans
}

// Test Case1: [-2,1,-3,4,-1,2,1,-5,4]	Output: 6
// Test Case2: [1]	Output: 1
// Test Case3: [5,4,-1,7,8]	Output: 23
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(maxSubArrayR(nums))
}
