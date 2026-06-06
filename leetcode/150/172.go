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

// trailingZeroesMath 使用数学方法
// * 思路: 要求阶乘中末尾 0 的个数, 可以将它分解为一个不含末尾 0 的数乘以 10^x, 其中 x 为整数, 即多少个 10 相乘
// * 对于一个 10, 除了 1 和它本身, 还有一对因子 2, 5 , 所以 10 的个数就是 2 的个数和 5 的个数的最小值,
// * 而 2 的个数肯定比 5 多, 因为在计算阶乘的过程中，每遇到一个偶数就会贡献一个因子 2, 而每遇到一个 5 的倍数才会贡献一个因子 5
// * 所以就要求 5 的个数, 这样就对于每个 5 可以有至少一个 2 与之相乘, 得到 10, 因此有几个 5, 就可以得到几个 10, 也就是最终阶乘结果的 0 的个数
// * 5 的个数 = n / 5 + n / 25 + n / 125 + ... + n / 5^k, 其中 5^k <= n
// * 这是因为有些数可以仅由一个 5 相乘(如 15 = 3 * 5) 得到, 有些数可以由两个 5 相乘(如 25 = 5 * 5) 得到, 所以需要将 n 除以 5 的各个幂次, 累加起来, 最终得到 5 的个数
// 这道题对数字的拆解比较精妙
func trailingZeroesMath(n int) int {
	count := 0
	for i := 5; i <= n; i = i * 5 {
		count += n / i
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
	fmt.Println(trailingZeroesMath(n))
}
