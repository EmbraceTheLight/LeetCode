// 69. x 的平方根
/*
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。
由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。


提示：
0 <= x <= 2^31 - 1
*/
package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	var left int64 = 0
	var right int64 = math.MaxInt32
	for left < right {
		mid := (left + right) / 2
		if mid*mid < int64(x) {
			if (mid+1)*(mid+1) > int64(x) {
				return int(mid)
			}
			left = mid + 1
		} else if mid*mid == int64(x) {
			return int(mid)
		} else {
			right = mid - 1
		}
	}

	return int(left)
}

// 示例 1：
// 输入：x = 4
// 输出：2
//
// 示例 2：
// 输入：x = 8
// 输出：2
// 解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(mySqrt(x))
}
