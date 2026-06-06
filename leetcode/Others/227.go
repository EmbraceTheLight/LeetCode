/*
227. 基本计算器 II
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
整数除法仅保留整数部分。
你可以假设给定的表达式总是有效的。所有中间结果将在 [-2^31, 2^31 - 1] 的范围内。
注意：不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。

1 <= s.length <= 3 * 10^5
s 由整数和算符 ('+', '-', '*', '/') 组成，中间由一些空格隔开
s 表示一个 有效表达式
表达式中的所有整数都是非负整数，且在范围 [0, 2^31 - 1] 内
题目数据保证答案是一个 32-bit 整数
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

var mp = make(map[byte]int)

func init() {
	mp['+'] = 1
	mp['-'] = 1
	mp['*'] = 2
	mp['/'] = 2
}

func calc(s string) (ret, length int) {
	var ans = 0
	var l = 0
	var idx = 0
	var lastOps byte = '+'
	var curOps byte = '+'

	for ; idx < len(s); idx++ {
		if s[idx] <= '9' && s[idx] >= '0' {
			var num int
			for idx < len(s) && (s[idx] <= '9' && s[idx] >= '0') {
				num = num*10 + int(s[idx]-'0')
				idx++
			}
			if mp[curOps] > mp[lastOps] {

			}
			lastOps = curOps
		} else {
			if s[idx] == '+' || s[idx] == '-' || s[idx] == '*' || s[idx] == '/' {
				curOps = s[idx]
			} else if s[idx] == ')' {
				break
			} else if s[idx] == ')' {
				calc(s[idx+1:])
			}

		}
	}
	return ans, l
}

func main() {
	fmt.Printf("Input s:")
	reader := bufio.NewReader(os.Stdin)
	result, _, _ := reader.ReadLine()
	fmt.Println(calculate(string(result)))
}
