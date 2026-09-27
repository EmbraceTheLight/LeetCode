// 16.17. 连续数列
/*
给定一个整数数组，找出总和最大的连续数列，并返回总和。

进阶：
如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/

package main

import (
	"fmt"
	. "lc/pkg"
	"math"
)

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n) // dp[i] 表示以 nums[i] 结尾的最大连续子数组和
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}
	var ans int = math.MinInt
	for i := 0; i < n; i++ {
		ans = max(ans, dp[i])
	}
	return ans
}

// 示例：
// 输入： [-2,1,-3,4,-1,2,1,-5,4]
// 输出： 6
// 解释： 连续子数组 [4,-1,2,1] 的和最大，为 6。
func main() {
	nums := CreateSlice[int]()
	fmt.Println(maxSubArray(nums))
}
