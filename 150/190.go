// 190. 颠倒二进制位
/*
颠倒给定的 32 位有符号整数的二进制位。

提示：
0 <= n <= 2^31 - 2
n 为偶数
*/
package main

import (
	"fmt"
	"math"
)

func reverseBits(n int) int {

}

func binaryToDecimal(s string) int {
	var num int
	n := len(s)
	count := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '1' {
			num += int(math.Pow(2, float64(count)))
		}
		count++
	}
	return num
}

func decimalToBinary(n int) string {
	var s string

	return s
}

// 示例 1：
// 输入：n = 43261596
// 输出：964176192
// 解释：
// 整数	二进制
// 43261596	00000010100101000001111010011100
// 964176192	00111001011110000010100101000000
//
// 示例 2：
// 输入：n = 2147483644
// 输出：1073741822
// 解释：
// 整数	二进制
// 2147483644	01111111111111111111111111111100
// 1073741822	00111111111111111111111111111110
func main() {
	var num int
	fmt.Println("Input num:")
	fmt.Scan(&num)
}
