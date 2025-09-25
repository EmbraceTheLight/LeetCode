package main

import (
	"fmt"
	"lc/pkg"
)

func robR(nums []int) int {
	dp := make([]int, len(nums)+1) // dp[i] 偷窃完第 i 个房间能获得的最大金额
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[len(nums)]
}

func main() {
	nums := pkg.CreateIntSlice[int]()
	fmt.Println(robR(nums))
}
