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
	//dp := make([]int, 10001)
	//for i := 0; i <= 100; i++ {
	//	dp[i*i] = 1 //完全平方数只需要其自身即可
	//}
	//dp[0] = 0
	//for i := 1; i <= n; i++ {
	//	if dp[i] == 1 {
	//		continue
	//	}
	//	dp[i] = math.MaxInt
	//	for j := i - 1; j >= 1; j-- {
	//		dp[i] = min(dp[i], dp[j]+dp[i-j])
	//	}
	//}
	//return dp

	//更好的解法，时间复杂度O(n*sqrt(n)),即O(n*n^(1/2))
	dp := make([]int, 10001) //dp[i]表示整数i最少需要多少个完全平方数
	for i := 0; i <= 100; i++ {
		dp[i*i] = 1 //完全平方数只需要其自身即可
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		if dp[i] == 1 {
			continue
		}
		dp[i] = math.MaxInt

		//FIXME：从1开始到sqrt(i)，而不用像第一种做法一样遍历[1,i],减小时间复杂度
		//FIXME：递推公式：dp[i]=min(dp[i-j*j])+1,1<=j<=sqrt(i)
		for j := 1; j*j <= i; j-- {
			dp[i] = min(dp[i], dp[i-j*j])
		}
		dp[i] = dp[i] + 1 //dp[i]+1即为dp[i-j*j]加上j这个完全平方数之后的结果
	}
	return dp[n]
}
func main() {
	var n int
	fmt.Println("Input n:")
	fmt.Scan(&n)
	fmt.Println(numSquares(n))
}
