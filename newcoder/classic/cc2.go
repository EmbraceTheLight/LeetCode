// CC2 后缀表达式求值
/*
计算逆波兰式（后缀表达式）的值 运算符仅包含"+","-","*"和""，被操作数是整数 保证表达式合法，除法时向下取整。
数据范围：表达式的长度满足： 进阶：空间复杂度 O(n)，时间复杂度 O(n)
例如：
["20", "10", "+", "30", "*"] -> ((20 + 10) * 30) -> 900
["40", "130", "50", "/", "+"] -> (40 + (130 / 50)) -> 42

示例1
输入: ["20","10","+","30","*"]
输出: 900

示例2
输入: ["20"]
输出: 20
*/
package main

import (
	"fmt"
	"lc/pkg"
	"strconv"
)

func evalRPN(tokens []string) int {
	// write code here
	numsStk := make([]int, 0)
	for _, token := range tokens {
		num, ok := strconv.Atoi(token)
		if ok == nil {
			numsStk = append(numsStk, num)
		} else {
			num1, num2 := numsStk[len(numsStk)-1], numsStk[len(numsStk)-2]
			numsStk = numsStk[:len(numsStk)-2]
			switch token {
			case "+":
				numsStk = append(numsStk, num1+num2)
			case "-":
				numsStk = append(numsStk, num2-num1)
			case "*":
				numsStk = append(numsStk, num1*num2)
			case "/":
				numsStk = append(numsStk, num2/num1)
			}
		}
	}
	return numsStk[0]
}
func main() {
	tokens := pkg.CreateSlice[string]()
	fmt.Println(evalRPN(tokens))
}
