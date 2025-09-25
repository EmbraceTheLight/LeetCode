/*
224. 基本计算器
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。

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
	"bufio"
	"fmt"
	"os"
)

// 计算在圆括号中的结果,第一个参数返回结果，第二个参数返回走过的字符串的距离，方便定位字符串索引
func calcuHelper(s string) (ans, l int) {
	length := len(s)
	var lastOperator byte //记录最近的一个运算符
	//var second int
	var HasFirst bool //是否存在第一个操作数

	for i := 0; i < length; i++ {
		var tmpSum int //记录某个字符数的数字形式表示

		if s[i] == ' ' {
			continue
		} else if s[i] == ')' {
			l = i //记录目前字符串索引的位置
			break
		} else if s[i] == '(' {
			sumAns, sumLen := calcuHelper(s[i+1:])
			if HasFirst == false { //还不存在第一个操作数
				if lastOperator == '-' { //处理开头为负号加括号的形式，如-(2+3)+6，此时‘-’为一元运算符
					ans = -sumAns
				} else {
					ans = sumAns
				}
				HasFirst = true
			} else { //已经存在第一个操作数
				if lastOperator == '-' {
					ans -= sumAns
				} else {
					ans += sumAns
				}
			}
			i += sumLen + 1
		} else if s[i] >= '0' && s[i] <= '9' {
			for ; i < length && (s[i] >= '0' && s[i] <= '9'); i++ {
				tmpSum = tmpSum*10 + int(s[i]-'0')
			}
			i--
			if HasFirst == false { //-为一元运算符
				if lastOperator == '-' {
					ans = -tmpSum
				} else {
					ans = tmpSum
				}
				HasFirst = true
			} else if HasFirst == true {
				if lastOperator == '-' {
					ans -= tmpSum
				} else {
					ans += tmpSum
				}
			}
		} else if s[i] == '+' || s[i] == '-' {
			lastOperator = s[i]
		}
	}
	return ans, l
}
func calculate(s string) int {
	var ans int
	ans, _ = calcuHelper(s)
	return ans
}
func main() {
	fmt.Printf("Input s:")
	reader := bufio.NewReader(os.Stdin)
	result, _, _ := reader.ReadLine()
	fmt.Println(calculate(string(result)))
}
