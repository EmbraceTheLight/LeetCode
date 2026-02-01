// 918. 环形子数组的最大和
/*
给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。
环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i] 的前一个元素是 nums[(i - 1 + n) % n] 。
子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，不存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。

示例 1：
输入：nums = [1,-2,3,-2]
输出：3
解释：从子数组 [3] 得到最大和 3

示例 2：
输入：nums = [5,-3,5]
输出：10
解释：从子数组 [5,5] 得到最大和 5 + 5 = 10

示例 3：
输入：nums = [3,-2,2,-3]
输出：3
解释：从子数组 [3] 和 [3,-2,2] 都可以得到最大和 3

提示：
n == nums.length
1 <= n <= 3 * 10^4
-3 * 10^4 <= nums[i] <= 3 * 10^4
*/
package main

import (
	"fmt"
	"lc/pkg"
)

//func maxSubarraySumCircular(nums []int) int {
//	n := len(nums)
//	nums = append(nums, nums...)
//	dp := make([]int, len(nums))
//	dp[0] = nums[0]
//	ans := dp[0]
//	for i := 1; i < n; i++ {
//		dp[i] = max(nums[i], dp[i-1]+nums[i])
//		ans = max(ans, dp[i])
//	}
//	for i := n; i < len(nums); i++ {
//		if nums[i] >= 0 {
//			dp[i] = dp[i-1]
//		} else {
//			dp[i-1] = dp[i-1] - nums[i%n]
//			ans = max(ans, dp[i-1])
//		}
//	}
//	return ans
//}

// * 采用反向思维, 求最小的子数组和, 再用数组总和减去这个最小子数组和, 即可得到最大环形子数组和
func maxSubarraySumCircular(nums []int) int {
	ans := nums[0]
	preMax, maxRes := nums[0], nums[0] // preMax: 以前一个元素结尾的最大子数组和, maxRes: 最大子数组和
	preMin, minRes := nums[0], nums[0] // preMin: 以前一个元素结尾的最小子数组和, maxRes: 最小子数组和
	sum := nums[0]                     // 数组数字总和
	for i := 1; i < len(nums); i++ {
		preMax = max(preMax+nums[i], nums[i])
		maxRes = max(maxRes, preMax)
		preMin = min(preMin+nums[i], nums[i])
		minRes = min(minRes, preMin)
		sum += nums[i]
	}

	// 特殊情况: 所有元素均为负数, 此时取最大的负数
	if maxRes < 0 {
		ans = maxRes
	} else {
		// 这里取 max(maxRes, sum-minRes) 是因为有可能数组中全为正数, 此时就不能用总数 - 最小子数组和了
		// 数组中有正数和负数时, ans = 总和 sum - 最小子数组和 minRes
		ans = max(maxRes, sum-minRes)
	}

	return ans
}

// Test Case1: [1,-2,3,-2]	Output: 3
// Test Case2: [5,-3,5]		Output: 10
// Test Case3: [3,-2,2,-3]	Output: 3
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(maxSubarraySumCircular(nums))
}
