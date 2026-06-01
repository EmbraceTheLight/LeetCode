// 70. 爬楼梯
/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

提示：
1 <= n <= 45
*/
package main

import "fmt"

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	var ans int         // 到达第 i 层的方法数
	var a, b int = 1, 1 // a: 到达第 i-2 层的方法数   b: 到达第 i-1 层的方法数
	for i := 1; i < n; i++ {
		ans = a + b
		a = b
		b = ans
	}
	return ans
}

// 示例 1：
// 输入：n = 2
// 输出：2
// 解释：有两种方法可以爬到楼顶。
// 1. 1 阶 + 1 阶
// 2. 2 阶
//
// 示例 2：
// 输入：n = 3
// 输出：3
// 解释：有三种方法可以爬到楼顶。
// 1. 1 阶 + 1 阶 + 1 阶
// 2. 1 阶 + 2 阶
// 3. 2 阶 + 1 阶
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(climbStairs(n))
}
