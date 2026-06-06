/*
跳跃游戏 II
给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
0 <= j <= nums[i]
i + j < n
返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

示例 1:
输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。

	从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

示例 2:
输入: nums = [2,3,0,1,4]
输出: 2

提示:
1 <= nums.length <= 10^4
0 <= nums[i] <= 1000
题目保证可以到达 nums[n-1]
*/
package main

import (
	"fmt"
	"math"
)

// 动态规划方法，好理解，时间复杂度O(n^2)
func jumpDP(nums []int) int {
	n := len(nums)
	dp := make([]int, n) //从起点到nums[i]所需最小步数
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j < n {
				dp[i+j] = min(dp[i+j], dp[i]+1)
			}
		}
	}

	return dp[n-1]
}

// 贪心,时间复杂度O(n)
func jump(nums []int) int {
	var ans int
	var right int     //维护能到达的最远处下标
	var maxLength int //维护在[i,right]中能到达的最远下标

	for i := 0; i < len(nums)-1; i++ { //不访问最后一个元素。因为一定确保可以到达最后一个元素，如果访问最后一个元素，在边界正好为最后一个位置的情况下，我们会增加一次「不必要的跳跃次数」
		maxLength = max(maxLength, i+nums[i])
		if i == right {
			right = maxLength
			ans++
		}
	}
	return ans
}
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println(jump(nums))
}
