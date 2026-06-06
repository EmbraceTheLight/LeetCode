/*
70. 爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

示例 1：
输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶

示例 2：
输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶

提示：
1 <= n <= 45
*/
package main

import "fmt"

// 其实就是斐波那契,跳到第i层只需跳到第i-1层再跳一层，或i-2层再跳2层
func climbStairs(n int) int {
	//可以不要dp数组，只用两个变量记录即可
	//var dp = make([]int, n+1)
	//dp[0] = 0
	//dp[1] = 1
	//for i := 2; i <= n; i++ {
	//	dp[i] = dp[i-1] + dp[i-2]
	//}
	var ans = 0
	var pre2 = 0
	var pre1 = 1
	for i := 1; i <= n; i++ {
		ans = pre2 + pre1
		pre2 = pre1
		pre1 = ans
	}
	//return dp[n]
	return ans
}
func main() {
	var n int
	fmt.Println("Input n:")
	fmt.Scanf("%d\n", &n)
	fmt.Println(climbStairs(n))
}
