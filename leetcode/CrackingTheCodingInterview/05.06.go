// 05.06. 整数转换
/*
整数转换。编写一个函数，确定需要改变几个位才能将整数 A 转成整数 B。

提示:
A，B范围在[-2147483648, 2147483647]之间
*/
package main

import "fmt"

func convertInteger(A int, B int) int {
	var ans int
	for i := 0; i < 32; i++ {
		if A&(1<<i) != B&(1<<i) {
			ans++
		}
	}
	return ans
}

// 示例 1：
//
//	输入：A = 29 （或者 0b11101）, B = 15（或者 0b01111）
//	输出：2
//
// 示例 2：
//
//	输入：A = 1，B = 2
//	输出：2
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(convertInteger(a, b))
}
