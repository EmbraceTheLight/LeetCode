/*
7.整数反转

给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。

示例 1：
输入：x = 123
输出：321

示例 2：
输入：x = -123
输出：-321

示例 3：
输入：x = 120
输出：21

示例 4：
输入：x = 0
输出：0

提示：
-2^31 <= x <= 2^31 - 1
*/
package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var ans int32
	var isNegative = x < 0
	x = int(math.Abs(float64(x)))
	for x != 0 {
		var tmp = int32(x % 10)
		//ta := ans
		//for i := 0; i < 9; i++ {
		//	last := ans //记录ans做运算之前的状态
		//	ans += ta
		//	if last > ans {
		//		return 0
		//	}
		//}
		//ans += tmp
		if ans > 214748364 || (ans == 214748364 && tmp > 7) {
			return 0
		}

		ans = ans*10 + tmp
		x /= 10
	}
	if isNegative == true {
		ans = -ans
	}
	return int(ans)
}
func main() {
	fmt.Println("Input x:")
	var x int
	fmt.Scan(&x)
	fmt.Println(reverse(x))
}
