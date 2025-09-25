package main

import (
	"fmt"
	"lc/pkg"
)

func maxProfit309R(prices []int) int {
	n := len(prices)
	// dp[i][0]：在第 i 天不持有股票且不处于冷冻期
	// dp[i][1]：在第 i 天持有股票;
	// dp[i][2]：在第 i 天不持有股票，且第二天不能购买股票（处于冷冻期）
	dp := make([][3]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = dp[i-1][1] + prices[i]
	}
	return max(dp[n-1][0], dp[n-1][2])
}

func main() {
	prices := pkg.CreateIntSlice[int]()
	fmt.Println(maxProfit309R(prices))
}
