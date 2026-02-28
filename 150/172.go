// 172. 阶乘后的零
/*
给定一个整数 n ，返回 n! 结果中尾随零的数量。
提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1

提示：
0 <= n <= 10^4

进阶：你可以设计并实现对数时间复杂度的算法来解决此问题吗
*/
package main

import "fmt"

func trailingZeroes(n int) int {
	const mod = 1e7
	if n == 0 {
		return 0
	}
	// 不能直接相乘得到阶乘结果, 否则会溢出
	count := 0
	var res int64 = 1
	for i := int64(n); i >= 1; i-- {
		res = res * i
		for res%10 == 0 {
			res = res / 10
			count++
		}
		res = res % mod // 防止溢出
	}
	return count
}

// 示例 1：
// 输入：n = 3
// 输出：0
// 解释：3! = 6 ，不含尾随 0
//
// 示例 2：
// 输入：n = 5
// 输出：1
// 解释：5! = 120 ，有一个尾随 0
//
// 示例 3：
// 输入：n = 0
// 输出：0
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(trailingZeroes(n))
}
