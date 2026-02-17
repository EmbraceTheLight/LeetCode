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
	"strconv"
)

func reverseBits(n int) int {
	reverseStr := func(s string) string {
		i, j := 0, len(s)-1
		b := []byte(s)
		for ; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
		return string(b)
	}
	return binaryToDecimal(reverseStr(decimalToBinary(n)))
}

// 二进制转十进制
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

// 十进制转二进制. 注意补充位数至 32 位.
// 由于本题指定 n 只能为非负数, 因此补充过程中不需要注意符号位是 1 (负数)还是 0 (正数)
func decimalToBinary(n int) string {
	var s string
	for n > 0 {
		s = strconv.Itoa(n%2) + s
		n = n / 2
	}
	tmp := ""
	for i := 0; i < 31-len(s); i++ {
		tmp += "0"
	}
	//if n < 0 {
	//	tmp = "1" + tmp
	//} else {
	//	tmp = "0" + tmp
	//}
	return tmp + s
}

// 使用位运算, 无需进制转换
func reverseBitsBetter(n int) int {
	var ret int
	for i := 0; i < 32; i++ {
		curBit := n & 1
		ret = ret | (curBit << (31 - i))
		n >>= 1
	}
	return ret
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
	fmt.Println(reverseBits(num))
	fmt.Println(reverseBitsBetter(num))
}
