// 188. 买卖股票的最佳时机 IV
/*
给你一个整数数组 prices 和一个整数 k ，其中 prices[i] 是某支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。也就是说，你最多可以买 k 次，卖 k 次。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

提示：
1 <= k <= 100
1 <= prices.length <= 1000
0 <= prices[i] <= 1000
*/
package main

import (
	"fmt"
	"lc/pkg"
	"math"
)

func maxProfit188(k int, prices []int) int {
	n := len(prices)
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k+1)
		for j := 0; j <= k; j++ {
			if k != 0 {
				dp[i][k][0], dp[i][k][1] = math.MinInt32, math.MinInt32
			}
		}
	}
	dp[0][0][1] = -prices[0]
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}

// 示例 1：
// 输入：k = 2, prices = [2,4,1]
// 输出：2
// 解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2。
//
// 示例 2：
// 输入：k = 2, prices = [3,2,6,5,0,3]
// 输出：7
// 解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4。
//
//	随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3。
func main() {
	var k int
	var prices []int

	fmt.Println("input k:")
	fmt.Scan(&k)
	fmt.Println("input prices:")
	prices = pkg.CreateSlice[int]()
	fmt.Println(maxProfit188(k, prices))
}
