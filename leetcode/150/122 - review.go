package main

import (
	"fmt"
	"lc/pkg"
)

func maxProfit122R1(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n) // dp[i][0] 第 i 天不持有股票的最大收益，dp[i][1] 第 i 天持有股票的最大收益
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[n-1][0]
}
func main() {
	prices := pkg.CreateSlice[int]()
	fmt.Println(maxProfit122R1(prices))
}
