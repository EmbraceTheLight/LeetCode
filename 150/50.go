// 50. Pow(x, n)
/*
实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。

提示：
-100.0 < x < 100.0
-2^31 <= n <= 2^31-1
n 是一个整数
要么 x 不为零，要么 n > 0 。
*/
package main

import "fmt"

func myPow(x float64, n int) float64 {
	exp := n
	if exp < 0 {
		exp = -exp
	}
	val := dfs50(x, exp)
	if n < 0 {
		return 1 / val
	}
	return val
}
func dfs50(cur float64, exp int) float64 {
	if exp == 2 || exp == 3 {
		ret := cur
		for i := 0; i < exp-1; i++ {
			ret *= cur
		}
		return ret
	}
	if exp%2 == 0 {
		return dfs50(cur*cur, exp/2)
	} else {
		return dfs50(cur*cur, exp/2) * cur
	}
}

// 示例 1：
// 输入：x = 2.00000, n = 10
// 输出：1024.00000
//
// 示例 2：
// 输入：x = 2.10000, n = 3
// 输出：9.26100
//
// 示例 3：
// 输入：x = 2.00000, n = -2
// 输出：0.25000
// 解释：2-2 = 1/22 = 1/4 = 0.25
func main() {
	var x float64
	var n int
	fmt.Println("Input x, n:")
	fmt.Scan(&x, &n)
	fmt.Println(myPow(x, n))
}
