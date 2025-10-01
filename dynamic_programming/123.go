/*
123. 买卖股票的最佳时机 III TODO:[hard]
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:
输入：prices = [3,3,5,0,0,3,1,4]
输出：6
解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。

	随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天（股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。

示例 2：
输入：prices = [1,2,3,4,5]
输出：4
解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。

	注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
	因这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。

示例 3：
输入：prices = [7,6,4,3,1]
输出：0
解释：在这个情况下, 没有交易完成, 所以最大利润为 0。

示例 4：
输入：prices = [1]
输出：0

提示：
1 <= prices.length <= 10^5
0 <= prices[i] <= 10^5
*/
package main

import (
	"fmt"
	"lc/pkg"
	"math"
)

// 方法1
func maxProfit123dp1(prices []int) int {
	n := len(prices)
	//FIXME：buy[i]表示第i次购买时的最大利润，sell[i]表示第i次卖出时的最大利润。sell[i]由buy[i]递推得到；而buy[i+1]由sell[i]递推得到
	buy, sell := make([]int, 3), make([]int, 3)
	for i := 0; i < 3; i++ { //将buy数组均初始化为最小值，表明还没有递推到这个数组
		buy[i] = math.MinInt32
	}
	sell[0] = 0
	//FIXME：有一点需要注意：这道题允许同一天买入·卖出股票，不过利润是0，不影响题目解答 一天可能买卖多次。
	for i := 0; i < n; i++ {
		for j := 1; j < 3; j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}
	return sell[2]
}

// 方法2
//
//	func maxProfit123dp2(prices []int) int {
//		n := len(prices)
//		dp[i][j][k]：i-->第i天，j-->第j次交易，k==0：没有股票；k==1：持有股票
//		dp := make([][3][2]int, n)
//		dp[0][0][1] = -prices[0]
//		for i := 0; i < n; i++ {
//			dp[i][]
//		}
//	}
func main() {
	prices := pkg.CreateSlice[int]()
	fmt.Println(maxProfit123dp1(prices))
}
