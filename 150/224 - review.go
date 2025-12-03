// 224. 基本计算器
/*
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。

示例 1：
输入：s = "1 + 1"
输出：2

示例 2：
输入：s = " 2-1 + 2 "
输出：3

示例 3：
输入：s = "(1+(4+5+2)-3)+(6+8)"
输出：23

提示：
1 <= s.length <= 3 * 10^5
s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
s 表示一个有效的表达式
'+' 不能用作一元运算(例如， "+1" 和 "+(2 + 3)" 无效)
'-' 可以用作一元运算(即 "-1" 和 "-(2 + 3)" 是有效的)
输入中不存在两个连续的操作符
每个数字和运行的计算将适合于一个有符号的 32位 整数
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 本题由于只有加减两种操作, 不用考虑优先级, 因此实现起来较为简单
/* 思路：
使用两个栈：一个数字栈和一个操作符栈。
例如 1+2-(3-4), 会先在数字栈中放入一个 0, 以便于处理 - 作为一元运算符的情况。
然后遍历字符串 s, 遇到数字则入栈, 遇到操作符则出栈两个数字进行计算, 并将结果入栈。
若遇到左括号, 则入栈一个操作符 '(', 并且再次入栈一个 0, 并再向操作符栈中插入 '+' 或 '-' 符号, 取决于左括号后紧跟着的是否是一元运算符 '-'.
对于上式的左括号, 这就相当于在原式当中进行补充：1+2-(0+3-4)
补充运算符的操作具体细节见 insertOp 函数。
*/
func calculate224R(s string) int {
	var nums []int
	var ops []byte

	// 操作数栈预先插入一个 0, 作为栈底
	nums = append(nums, 0)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			continue
		case '+':
			calc(&nums, &ops)
			ops = append(ops, '+')
		case '-':
			calc(&nums, &ops) // 无需考虑优先级, 计算前两个操作数的结果
			ops = append(ops, '-')
		case '(':
			ops = append(ops, '(')
			nums = append(nums, 0)     // 向数字栈中插入 0, 作为预处理
			i = insertOp(s, i+1, &ops) // 确定 0 后应当插入的操作符符号为 + 还是 -
		case ')':
			for ops[len(ops)-1] != '(' {
				calc(&nums, &ops)
			}
			ops = ops[:len(ops)-1] // 弹出 '('
		default:
			num, idx := getNumber(s, i)
			i = idx - 1
			nums = append(nums, num)
		}
	}
	calc(&nums, &ops)
	return nums[len(nums)-1]
}

// getNumber 获取一个数字, 并返回数字和下一个字符在 s 中的索引.
func getNumber(s string, startIndex int) (res, idx int) {
	num := 0
	i := startIndex
	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		tmp := int(s[i] - '0')
		num = num*10 + tmp
	}
	return num, i
}

// calc 计算栈顶两个数字，并将结果入栈. 本题只有 + 和 - 两种运算符，因此不用考虑运算符优先级.
// 若有乘法&除法, 对该函数进行修改即可.
func calc(nums *[]int, ops *[]byte) {
	if len(*ops) == 0 {
		return
	}
	num1, num2 := (*nums)[len(*nums)-2], (*nums)[len(*nums)-1]
	op := (*ops)[len(*ops)-1]
	*nums = (*nums)[:len(*nums)-2]
	*ops = (*ops)[:len(*ops)-1]
	if op == '+' {
		*nums = append(*nums, num1+num2)
	} else if op == '-' {
		*nums = append(*nums, num1-num2)
	}
}

// insertOp 向 ops 栈中插入一个操作符. 该操作用于遇到左括号的情况. 这种情况会先向数字栈中插入 0, 之后调用该函数确定 0 后应当插入的操作符符号为 + 或 -.
// 若遇到的第一个非空操作符为 '-', 则向 ops 栈中插入 '-' 符号, 否则插入 '+'.
// 返回下一个字符在 s 中的索引.
func insertOp(s string, startIdx int, ops *[]byte) (idx int) {
	for s[startIdx] == ' ' {
		startIdx++
	}
	if s[startIdx] == '-' {
		*ops = append(*ops, '-')
	} else {
		*ops = append(*ops, '+')
		startIdx--
	}
	return startIdx
}

// Test Case1: "1 + 1"					Output: 2
// Test Case2: " 2-1 + 2 "				Output: 3
// Test Case3: "(1+(4+5+2)-3)+(6+8)"	Output: 23
// Test Case4: 1-(-2)					Output: 3
func main() {
	s := pkg.CreateString()
	fmt.Println(calculate224R(s))
}
