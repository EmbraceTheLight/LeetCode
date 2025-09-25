/*
518. 零钱兑换 II
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
假设每一种面额的硬币有无限个。
题目数据保证结果符合 32 位带符号整数。

示例 1：
输入：amount = 5, coins = [1, 2, 5]
输出：4
解释：有四种方式可以凑成总金额：
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1

示例 2：
输入：amount = 3, coins = [2]
输出：0
解释：只用面额 2 的硬币不能凑成总金额 3。

示例 3：
输入：amount = 10, coins = [10]
输出：1

提示：

1 <= coins.length <= 300
1 <= coins[i] <= 5000
coins 中的所有值 互不相同
0 <= amount <= 5000
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// FIXME：完全背包
func change(amount int, coins []int) int {

	//FIXME：这道题重点/难点在于找到的结果不重复，即找的是组合数，而不是排列数
	dp := make([]int, amount+1) //dp[i]；到达数值i的方法数
	dp[0] = 1
	//for i := 1; i <= amount; i++ {
	//	for j := i; j > i/2; j-- {
	//
	//	}
	//}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ { //FIXME：内层循环表示：当选择硬币coins[i]时，方案数则取决于凑成j-coins[i]的方案数，将方案累加即可
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

// 二维dp
func changedp2(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1) //二维数组，dp[i][j]表示当选择前i中面值的硬币时，组成数值j的方案数
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
	}

	//初始化dp数组
	//组成面额0的方案数，只有一种。即不选硬币
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= amount; i++ {
		dp[0][i] = 0
	}

	//注意：循环次序不能反。如果先枚举的是金额，然后再枚举的硬币，则问题就变成了在某个金额下，由硬币组成的方案数，就成了排列数
	//如果先枚举硬币，则内层循环就用不了之前面值的硬币了，就没有了重复的问题
	for i := 1; i <= n; i++ { //枚举硬币
		for j := 1; j <= amount; j++ { //枚举金额
			if j < coins[i-1] { //金额小于当前硬币面值，则直接使用【不使用当前面值硬币】的方法数
				dp[i][j] = dp[i-1][j]
			} else { //否则，将【不使用当前面值硬币】的方法数 与 【使用当前面值硬币】的方法数 相加即可
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[n][amount]
}

func main() {
	fmt.Println("Input amount:")
	var amount int

	fmt.Scan(&amount)
	coins := pkg.CreateIntSlice[int]()
	fmt.Println(change(amount, coins))
}
