package main

import (
	"fmt"
	"lc/pkg"
)

func canJump55R1(nums []int) bool {
	n := len(nums)
	// 特殊情况判断
	// 只有一个元素
	if n == 1 {
		return true
	}
	dp := make([]int, n) // dp[i] 表示在位置 i 处能跳到的最远距离
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		if i <= dp[i-1] {
			dp[i] = max(dp[i-1], i+nums[i])
			if dp[i] >= n-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

// Test Data1: 5  2 3 1 1 4
// Test Data2: 5  3 2 1 0 4
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(canJump55R1(nums))
}
