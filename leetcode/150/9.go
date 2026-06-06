// 9. 回文数
/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。

提示：
-2^31 <= x <= 2^31 - 1

进阶：你能不将整数转为字符串来解决这个问题吗？
*/
package main

import (
	"fmt"
	"strconv"
)

func isPalindrome9(x int) bool {
	str := strconv.Itoa(x)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

// * 不将数字转换为字符串
// * 思路: 反转 x 的后半部分数字, 比较后半部分与前半部分, 如果二者相等, 则 x 是回文数
// * 判断是否反转了后半部分数字: 原始数字小于或等于反转后的数字时，就意味着我们已经处理了一半位数的数字了
// * 注意边界判断
func isPalindrome9Better(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 { // 负数或个位数为 0 的数, 不可能是回文数, 特殊地, 0 是回文数
		return false
	}
	revertNum := 0

	// x > revertNum, 表示 x 的位数还没有反转完一半, 继续反转
	for x > revertNum {
		revertNum = revertNum*10 + x%10
		x = x / 10
	}

	// x 有偶数位数(如 12344321), 判断 x == revertNum (1234 == 1234),
	// 否则(如 123454321), 舍去最中间的一个数 (即 revertNum/10 = 12345 / 10 = 1234) 然后再判断
	return x == revertNum || x == revertNum/10
}

// 示例 1：
// 输入：x = 121
// 输出：true
//
// 示例 2：
// 输入：x = -121
// 输出：false
// 解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//
// 示例 3：
// 输入：x = 10
// 输出：false
// 解释：从右向左读, 为 01 。因此它不是一个回文数。
func main() {
	var x int
	fmt.Println(0 % 10)
	fmt.Scan(&x)
	fmt.Println(isPalindrome9(x))
	fmt.Println(isPalindrome9Better(x))
}
