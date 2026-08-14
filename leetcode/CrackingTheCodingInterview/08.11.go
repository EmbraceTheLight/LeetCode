// 08.11. 硬币
/*
硬币。给定数量不限的硬币，币值为25分、10分、5分和1分，编写代码计算n分有几种表示法。(结果可能会很大，你需要将结果模上1000000007)
说明：
注意:
你可以假设：
0 <= n (总金额) <= 1000000
*/
package main

import "fmt"

func waysToChange(n int) int {
	if n == 0 {
		return 0
	}
	const bigMod = 1000000007
	dp := make([]int, n+1)
	faceValues := [4]int{1, 5, 10, 25}
	dp[0] = 1 // dp[i] 表示 i分 有几种表示法

	// 外层循环遍历四种面值
	// 不会出现重复面值(如 1 + 5 和 5 + 1)的情况, 因为这里遍历的每种情况, 硬币面值都是不同的
	for _, j := range faceValues {
		for i := j; i <= n; i++ {
			dp[i] = (dp[i] + dp[i-j]) % bigMod
		}
	}

	return dp[n]
}

// 示例 1：
// 输入：n = 5
// 输出：2
// 解释：有两种方式可以凑成总金额:
// 5=5
// 5=1+1+1+1+1
//
// 示例 2：
// 输入：n = 10
// 输出：4
// 解释：有四种方式可以凑成总金额:
// 10=10
// 10=5+5
// 10=5+1+1+1+1+1
// 10=1+1+1+1+1+1+1+1+1+1
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(waysToChange(n))
}
