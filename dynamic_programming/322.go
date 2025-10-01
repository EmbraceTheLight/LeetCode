/*
322. 零钱兑换
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。

示例 1：
输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

示例 2：
输入：coins = [2], amount = 3
输出：-1

示例 3：
输入：coins = [1], amount = 0
输出：0

提示：

1 <= coins.length <= 12
1 <= coins[i] <= 2^31 - 1
0 <= amount <= 10^4
*/
package main

import (
	"fmt"
	"lc/pkg"
	"sort"
)

func coinChange(coins []int, amount int) int {
	n := len(coins)
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[j]
	})
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for j := 0; j < n; j++ {
			if coins[j] > i {
				break
			} else {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for _, c := range coins {
		for j := 1; j <= amount; j++ {
			if c <= j {
				dp[j] = min(dp[j], dp[j-c]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
func main() {
	fmt.Println("Input amount:")
	var amount int
	fmt.Scan(&amount)
	coins := pkg.CreateSlice[int]()
	fmt.Println(coinChange(coins, amount))
}
