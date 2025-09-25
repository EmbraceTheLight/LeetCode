/*
415. 字符串相加
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。
你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。

示例 1：
输入：num1 = "11", num2 = "123"
输出："134"

示例 2：
输入：num1 = "456", num2 = "77"
输出："533"

示例 3：
输入：num1 = "0", num2 = "0"
输出："0"

1 <= num1.length, num2.length <= 10^4
num1 和num2 都只包含数字 0-9
num1 和num2 都不包含任何前导零
*/
package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	var ans string
	var c int
	l1 := len(num1) - 1
	l2 := len(num2) - 1
	for l1 >= 0 && l2 >= 0 {
		a := int(num1[l1] - '0')
		b := int(num2[l2] - '0')
		sum := a + b + c
		if sum >= 10 {
			ans = strconv.Itoa(a+b+c-10) + ans
			c = 1
		} else {
			ans = strconv.Itoa(a+b+c) + ans
			c = 0
		}
		l1--
		l2--
	}

	for l1 >= 0 {
		if c+int(num1[l1]-'0') >= 10 {
			ans = strconv.Itoa(c+int(num1[l1]-'0')-10) + ans
			c = 1
		} else {
			ans = strconv.Itoa(c+int(num1[l1]-'0')) + ans
			c = 0
		}
		l1--
	}
	for l2 >= 0 {
		if c+int(num2[l2]-'0') >= 10 {
			ans = strconv.Itoa(c+int(num2[l2]-'0')-10) + ans
			c = 1
		} else {
			ans = strconv.Itoa(c+int(num2[l2]-'0')) + ans
			c = 0
		}
		l2--
	}
	if c == 1 {
		ans = "1" + ans
	}

	return ans
}
func main() {
	var str1 string
	var str2 string
	fmt.Scanf("%s\r\n", &str1)
	fmt.Scanf("%s\r\n", &str2)
	fmt.Println(addStrings(str1, str2))
}
