// 08.05. 递归乘法
/*
递归乘法。 写一个递归函数，不使用 * 运算符，
实现两个正整数的相乘。可以使用加号、减号、位移，但要吝啬一些。

提示：
保证乘法范围不会溢出
*/
package main

import "fmt"

func multiply(A int, B int) int {
	ans := 0
	small, big := getSmallAndBig(A, B)
	cache := map[int]int{
		0: big,
	}
	// small 可以拆分成 2^x1 + 2^x2 + ... 2^xn (+1) (small 为奇数则最后再加 1)
	// cache 保存 2^0 * big, 2^1 * big, 2^2 * big, ...
	// 当 small 为 0, 递归终止, 得到答案
	dfs0805(small, big, 0, cache, &ans)
	return ans
}
func getSmallAndBig(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

// small 较小数, big 较大数
// expr: 当前指数(以 2 为底)
// cache: key: 指数, value:  2 ^ key * big
func dfs0805(small, big int, expr int, cache map[int]int, ans *int) {
	if small == 0 {
		return
	}
	if expr > 0 {
		cache[expr] = cache[expr-1] + cache[expr-1]
	}
	if small&1 == 1 {
		*ans += cache[expr]
	}
	small = small >> 1
	expr++
	dfs0805(small, big, expr, cache, ans)
}

/* 空间更优化的解法, 不需要 cache map
func multiply(A int, B int) int {
	ans := 0
	small, big := getSmallAndBig(A, B)

	// small 可以拆分成 2^x1 + 2^x2 + ... 2^xn (+1) (small 为奇数则最后再加 1)
	// 当 small 为 0, 递归终止, 得到答案
	dfs0805(small, big, &ans)
	return ans
}
func getSmallAndBig(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

// small 较小数, big 较大数
// expr: 当前指数(以 2 为底)
func dfs0805(small, big int, ans *int) {
	if small == 0 {
		return
	}

	if small&1 == 1 {
		*ans += big
	}

	dfs0805(small>>1, big+big, ans)
}
*/

// 示例 1：
// 输入：A = 1, B = 10
// 输出：10
//
// 示例 2：
// 输入：A = 3, B = 4
// 输出：12
func main() {
	var a, b int
	fmt.Println("Input a, b:")
	fmt.Scan(&a, &b)
	fmt.Println(multiply(a, b))
}
