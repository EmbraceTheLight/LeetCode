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

// 使用位运算+分治, O(1) 时间复杂度
// * 思想: 将二进制位折半, 左半边颠倒, 右半边颠倒, 再将右半边拼接到左半边前即可, 左右两边翻转还可以细分
// * 以 8 位二进制数举例: 12345678, 它们分别代表第几位, 即 1 表示第一位, 2 表示第二位, 以此类推
// * 先两位两位翻转, 翻转后: 21436587
// * 再四位四位翻转, 翻转后: 43218765
// * 最后翻转全部位数, 最后结果位: 87654321
/*
官方题解:
对于递归的最底层，我们需要交换所有奇偶位：
取出所有奇数位和偶数位；
将奇数位移到偶数位上，偶数位移到奇数位上。

类似地，对于倒数第二层，每两位分一组，
按组号取出所有奇数组和偶数组，然后将奇数组移到偶数组上，偶数组移到奇数组上。以此类推。
*/
func reverseBitsBest(n int) int {
	// 定义掩码辅助位翻转
	const (
		MASK1 = 0x55555555 // 01010101010101010101010101010101
		MASK2 = 0x33333333 // 00110011001100110011001100110011
		MASK3 = 0x0f0f0f0f // 00001111000011110000111100001111
		MASK4 = 0x00ff00ff // 00000000111111110000000011111111
		MASK5 = 0x0000ffff // 00000000000000001111111111111111
	)
	ret := n
	ret = ret>>1&MASK1 | ret&MASK1<<1
	ret = ret>>2&MASK2 | ret&MASK2<<2
	ret = ret>>4&MASK3 | ret&MASK3<<4
	ret = ret>>8&MASK4 | ret&MASK4<<8
	ret = ret>>16&MASK5 | ret&MASK5<<16
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
	fmt.Println(reverseBitsBest(num))
}
