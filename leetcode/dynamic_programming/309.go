/*
309. 买卖股票的最佳时机含冷冻期
给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格。
设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:
输入: prices = [1,2,3,0,2]
输出: 3
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]

示例 2:
输入: prices = [1]
输出: 0

提示：

1 <= prices.length <= 5000
0 <= prices[i] <= 1000
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func maxProfit309(prices []int) int {
	n := len(prices)
	//FIXME：dp[i][0]表示当天持有股票（不处于冷冻期），
	//FIXME：dp[i][1]表示当天不持有股票且不处于冷冻期，即明天能买
	//FIXME：dp[i][2]表示当天不持有股票且明天不能买-即当天买了股票
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i]) //要么上一天持有股票，今天继续持有，要么上一天没有股票，今天购买
		dp[i][1] = max(dp[i-1][1], dp[i-1][2])           //昨天未持有股票且前天没卖出，或者昨天未持有股票但是前天卖出了(即昨天处于冷冻期)
		dp[i][2] = dp[i-1][0] + prices[i]                //昨天卖出了持有的股票
	}
	return max(dp[n-1][1], dp[n-1][2])
}
func main() {
	sli := pkg.CreateSlice[int]()
	fmt.Println(maxProfit309(sli))
}
