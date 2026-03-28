// 416. 分割等和子集
/*
给你一个 只包含正整数 的 非空 数组 nums。
请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

提示：
1 <= nums.length <= 200
1 <= nums[i] <= 100
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// * 01 背包问题
// * 即求解 nums 是否存在一组数字, 恰好能塞入容量为 sum(nums) / 2 （sum(nums) 为偶数）的背包中
func canPartition(nums []int) bool {
	n := len(nums)

	sum := nums[0]
	for i := 1; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([][]bool, n+1) // dp[i][j]	选择 nums 前 i 个数字, 能否恰好塞满容量为 j 的背包
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, target+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = true // 不选数字, 容量为 0 的背包一定能塞满
	}

	for i := 0; i < n; i++ {
		for j := 1; j <= target; j++ {
			if j >= nums[i] { // 如果当前 nums[i] 小于要插入的背包的容量 j, 则有两个选择: 塞入背包(dp[i][j-nums[i]]) 或不塞入背包(dp[i-1][j])
				dp[i+1][j] = dp[i][j] || dp[i][j-nums[i]]
			} else { // 如果当前 nums[i] 大于要插入的背包的容量 j, 则只能选择不塞入背包(dp[i-1][j])
				dp[i+1][j] = dp[i][j]
			}
		}
	}
	return dp[n][target]
}

// 示例 1：
// 输入：nums = [1,5,11,5]
// 输出：true
// 解释：数组可以分割成 [1, 5, 5] 和 [11]。
//
// 示例 2：
// 输入：nums = [1,2,3,5]
// 输出：false
// 解释：数组不能分割成两个元素和相等的子集。
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(canPartition(nums))
}
