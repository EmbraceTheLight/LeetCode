/*
377. 组合总和 Ⅳ
给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
题目数据保证答案符合 32 位整数范围。

示例 1：
输入：nums = [1,2,3], target = 4
输出：7
解释：
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
请注意，顺序不同的序列被视作不同的组合。

示例 2：
输入：nums = [9], target = 3
输出：0

提示：
1 <= nums.length <= 200
1 <= nums[i] <= 1000
nums 中的所有元素 互不相同
1 <= target <= 1000
*/
package main

import (
	"fmt"
	"lc/pkg"
	"sort"
)

func combinationSum4(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, target+1) //dp[i]表示nums中能组成总和为 i 的元素总和个数
	dp[0] = 1
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	//排列数
	for i := 1; i <= target; i++ {
		for j := 0; j < n && nums[j] <= i; j++ {
			dp[i] += dp[i-nums[j]]
		}
	}

	//组合数
	//for i := 0; i < n; i++ {
	//	for j := nums[i]; j <= target; j++ {
	//		dp[j] += dp[j-nums[i]]
	//	}
	//}
	return dp[target]
}
func main() {
	target := 0
	fmt.Scan(&target)
	nums := pkg.CreateSlice[int]()
	fmt.Println(combinationSum4(nums, target))
}
