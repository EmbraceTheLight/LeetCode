// 123. 买卖股票的最佳时机 III
/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

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

// 方法一 枚举所有状态. 不易写对, 不易调试
func maxProfit123(prices []int) int {
	n := len(prices)

	// dp[i][j][k]: i => 第 i 天, j => 交易次数, k => 是否持有 k=0 不持有. k=1 持有
	// 只有完成一次买入/卖出才算一次交易
	dp := make([][3][2]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 2; k++ {
				dp[i][j][k] = math.MinInt32
			}
		}
	}
	for i := 0; i < n; i++ {
		dp[i][0][0] = 0
	}
	dp[0][1][1] = -prices[0]
	for i := 1; i < n; i++ {
		for j := 1; j < 3; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return max(0, dp[n-1][1][0], dp[n-1][2][0])
}

// 方法二
// * 思路: 使用两个一维数组 dp1 和 dp2, dp1[i] 表示在 [0, i] 内交易一次的最大利润, dp2[j] 表示在 [j, n) 内交易一次的最大利润
// * 当遍历到第 i 天时, dp1[i] + dp2[i] 即以 i 为界限, 将两侧交易一次可以获得的最大利润相加
// * 答案即为这些和中最大的那个
// * 思路简单易行, 但是使用范围局限性较大
func maxProfit123_2(prices []int) int {
	n := len(prices)
	dp1 := make([]int, n) // dp1[i]: 在 [0, i] 内交易一次的最大利润
	dp2 := make([]int, n) // dp2[i]: 在 [i, n) 内交易一次的最大利润
	//dp1[0] = -prices[0]
	tmpMin := prices[0]
	for i := 1; i < n; i++ {
		dp1[i] = max(dp1[i-1], prices[i]-tmpMin)
		tmpMin = min(tmpMin, prices[i])
	}
	tmpMax := prices[n-1]
	for i := n - 1; i > 0; i-- {
		dp2[i] = max(dp2[i-1], tmpMax-prices[i])
		tmpMax = max(tmpMax, prices[i])
	}

	ans := dp1[0]
	for i := 0; i < n; i++ {
		ans = max(ans, dp2[i]+dp1[i])
	}
	return ans
}

// 示例 1:
// 输入：prices = [3,3,5,0,0,3,1,4]
// 输出：6
// 解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
//
// 随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
//
// 示例 2：
// 输入：prices = [1,2,3,4,5]
// 输出：4
// 解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//
//	注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
//	因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
//
// 示例 3：
// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这个情况下, 没有交易完成, 所以最大利润为 0。
//
// 示例 4：
// 输入：prices = [1]
// 输出：0
func main() {
	prices := pkg.CreateSlice[int]()
	fmt.Println(maxProfit123(prices))
	fmt.Println(maxProfit123_2(prices))
}
