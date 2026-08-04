// 08.01. 三步问题
/*
三步问题。有个小孩正在上楼梯，楼梯有 n 阶台阶，小孩一次可以上 1 阶、2 阶或 3 阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模 1000000007。

提示:
n 范围在[1, 1000000]之间
*/
package main

import "fmt"

func waysToStep(n int) int {
	big := 1000000007
	steps := [3]int{1, 2, 4}
	if n == 1 {
		return steps[0]
	} else if n == 2 {
		return steps[1]
	} else if n == 3 {
		return steps[2]
	}
	for i := 4; i <= n; i++ {
		steps[(i-1)%3] = (steps[0] + steps[1] + steps[2]) % big
	}
	return steps[(n-1)%3]
}

// 示例 1：
// 输入：n = 3
// 输出：4
// 说明：有四种走法
//
// 示例 2：
// 输入：n = 5
// 输出：13
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(waysToStep(n))
}
