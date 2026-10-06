// 16.26. 计算器
/*
给定一个包含正整数、加(+)、减(-)、乘(*)、除(/)的算数表达式(括号除外)，计算其结果。
表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。

说明：
你可以假设所给定的表达式都是有效的。
请不要使用内置的库函数 eval。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func calculate(s string) int {
	nums := make([]int, 0)
	ops := make([]byte, 0)
	priority := map[byte]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}
	var ans int
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		} else if s[i] >= '0' && s[i] <= '9' {
			num, idx := getNum(s, i)
			nums = append(nums, num)
			i = idx - 1
		} else {
			ops = append(ops, s[i])
		}
	}
	for len(ops) > 0 {
		op := ops[0]
		num1, num2 := nums[0], nums[1]
		if len(ops) > 1 && priority[op] < priority[ops[1]] {
			op = ops[1]
			ops = append(ops[0:1], ops[2:]...)
			num1, num2 = nums[1], nums[2]
			nums = append(nums[0:1], nums[2:]...)
			nums[1] = calcHandle(num1, num2, op)
		} else {
			ops = ops[1:]
			nums = nums[1:]
			nums[0] = calcHandle(num1, num2, op)
		}
	}
	if len(nums) > 0 {
		ans = nums[0]
	}
	return ans
}
func getNum(s string, start int) (num, idx int) {
	sum := 0
	i := start
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		sum = sum*10 + int(s[i]-'0')
	}
	return sum, i
}

func calcHandle(num1, num2 int, op byte) int {
	switch op {
	case '+':
		return num1 + num2
	case '-':
		return num1 - num2
	case '*':
		return num1 * num2
	case '/':
		return num1 / num2
	default:
		return 0
	}
}

// 示例 1：
// 输入："3+2*2"
// 输出：7
//
// 示例 2：
// 输入：" 3/2 "
// 输出：1
//
// 示例 3：
// 输入：" 3+5 / 2 "
// 输出：5
func main() {
	var s = CreateString()
	fmt.Println(calculate(s))
}
