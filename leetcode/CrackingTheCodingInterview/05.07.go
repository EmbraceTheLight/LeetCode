// 05.07. 配对交换
/*
配对交换。编写程序，交换某个整数的奇数位和偶数位，尽量使用较少的指令（也就是说，位 0 与位 1 交换，位 2 与位 3 交换，以此类推）。

提示:
num 的范围在[0, 2^0 - 1]之间，不会发生整数溢出。
*/
package main

import "fmt"

func exchangeBits(num int) int {
	var ans int
	odd, even := 0b01010101010101010101010101010101, 0b10101010101010101010101010101010 // odd 提取 num 奇数部分, even 提取 num 偶数部分
	ans = ((num & odd) << 1) | ((num & even) >> 1)
	return ans
}

// 示例 1：
// 输入：num = 2（或者 0b10）
// 输出：1 (或者 0b01)
// 示例 2：
// 输入：num = 3
// 输出：3
func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(exchangeBits(num))
}
