// 66. 加一
/*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
将大整数加 1，并返回结果的数字数组。

提示：
1 <= digits.length <= 100
0 <= digits[i] <= 9
digits 不包含任何前导 0。
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func plusOne(digits []int) []int {
	n := len(digits)
	carry := 0
	if digits[n-1]+1 >= 10 {
		digits[n-1] = (digits[n-1] + 1) % 10
		carry = 1
	} else {
		digits[n-1]++
		return digits
	}

	for i := n - 2; i >= 0; i-- {
		if digits[i]+carry < 10 {
			digits[i] = digits[i] + carry
			carry = 0
			break
		}
		digits[i] = (digits[i] + carry) % 10
		carry = 1
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// 示例 1：
// 输入：digits = [1,2,3]
// 输出：[1,2,4]
// 解释：输入数组表示数字 123。
// 加 1 后得到 123 + 1 = 124。
// 因此，结果应该是 [1,2,4]。
//
// 示例 2：
// 输入：digits = [4,3,2,1]
// 输出：[4,3,2,2]
// 解释：输入数组表示数字 4321。
// 加 1 后得到 4321 + 1 = 4322。
// 因此，结果应该是 [4,3,2,2]。
//
// 示例 3：
// 输入：digits = [9]
// 输出：[1,0]
// 解释：输入数组表示数字 9。
// 加 1 得到了 9 + 1 = 10。
// 因此，结果应该是 [1,0]。
func main() {
	digits := pkg.CreateSlice[int]()
	fmt.Println(plusOne(digits))
}
