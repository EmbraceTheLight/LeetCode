/*
53.最大子数组和
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组
是数组中的一个连续部分。
示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

示例 2：
输入：nums = [1]
输出：1

示例 3：
输入：nums = [5,4,-1,7,8]
输出：23

提示：
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

// dp[i]表示以nums[i]结尾的最大子数组和
func maxSubArray(nums []int) int {
	var ans = nums[0]
	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > dp[i-1]+nums[i] {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		ans = max(ans, dp[i])
	}
	return ans
}
func main() {
	nums := pkg.CreateIntSlice()
	fmt.Println(maxSubArray(nums))
}
