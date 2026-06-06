/*
279. 完全平方数
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

示例 1：
输入：n = 12
输出：3
解释：12 = 4 + 4 + 4

示例 2：
输入：n = 13
输出：2
解释：13 = 4 + 9

提示：
1 <= n <= 10^4
*/
package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n+1) //dp[i]：数字i的完全平方数数量
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt
		//利用之前已经求过的平方数来进行遍历
		for j := 0; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j])
		}
		dp[i] += 1 //i=(i-j*j)+j*j
	}
	return dp[n]
}
func main() {
	var n int
	fmt.Println("Input n:")
	fmt.Scan(&n)
	fmt.Println(numSquares(n))
}
