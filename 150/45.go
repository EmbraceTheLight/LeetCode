// 跳跃游戏 II
/*
给定一个长度为 n 的 0 索引整数数组 nums。初始位置在下标 0。

每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在索引 i 处，你可以跳转到任意 (i + j) 处：
0 <= j <= nums[i] 且 i + j < n
返回到达 n - 1 的最小跳跃次数。测试用例保证可以到达 n - 1。

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
题目保证可以到达 n - 1
*/
package main

import (
	"fmt"
	"lc/pkg"
	"math"
)

// 朴素做法
func jump45(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	dp := make([]int, n) // dp[i] 表示从 0 到 i 位置的最小跳跃次数
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt
		for j := i - 1; j >= 0; j-- {
			if j+nums[j] >= i {
				dp[i] = min(dp[i], dp[j]) + 1
			}
		}
	}
	return dp[n-1]
}

// 优化后的做法
func jump45Better(nums []int) int {
	step := 0
	maxPos := 0 // 最远能跳到的位置下标
	end := 0    // 跳 step 步时最多能到达的位置下标
	n := len(nums)

	// 特殊情况：只有一个元素，不用跳。条约次数为 0
	if n == 1 {
		return 0
	}
	for i := 0; i < n; i++ {
		// maxPos： 在 [i, end] 范围内寻找最远能跳到的位置
		maxPos = max(maxPos, i+nums[i])

		// 到达了最远能跳到的位置，更新 end 位置，即更新最多能到达的位置下标，同时步数 + 1
		if i == end {
			end = maxPos
			step++
			if end >= n-1 { // 到达了最后一个元素，跳出循环
				break
			}
		}
	}
	return step
}

// Test Data1: 5  2 3 1 1 4  	Expected Output: 2
// Test Data2: 5  2 3 0 1 4  	Expected Output: 2
// Test Data3: 7  1 2 0 2 0 1 1	Expected Output: 4
// Test Data4: 15 7 0 9 6 9 6 1 7 9 0 1 2 9 0 3		Expected Output: 2
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(jump45Better(nums))
}
