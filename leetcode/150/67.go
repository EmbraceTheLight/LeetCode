// 67. 二进制求和
/*
给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。

提示：
1 <= a.length, b.length <= 10^4
a 和 b 仅由字符 '0' 或 '1' 组成
字符串如果不是 "0" ，就不含前导零
*/
package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	na, nb := len(a), len(b)
	carry := 0
	var ans string
	i, j := na-1, nb-1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		tmpa := int(a[i] - '0')
		tmpb := int(b[j] - '0')
		ans = strconv.Itoa((tmpa+tmpb+carry)%2) + ans
		if tmpa+tmpb+carry >= 2 {
			carry = 1
		} else {
			carry = 0
		}
	}
	for ; i >= 0; i-- {
		tmpa := int(a[i] - '0')
		ans = strconv.Itoa((tmpa+carry)%2) + ans
		if tmpa+carry >= 2 {
			carry = 1
		} else {
			carry = 0
		}
	}
	for ; j >= 0; j-- {
		tmpb := int(a[j] - '0')
		ans = strconv.Itoa((tmpb+carry)%2) + ans
		if tmpb+carry >= 2 {
			carry = 1
		} else {
			carry = 0
		}
	}
	if carry == 1 {
		ans = "1" + ans
	}
	return ans
}

// 示例 1：
// 输入:a = "11", b = "1"
// 输出："100"
//
// 示例 2：
// 输入：a = "1010", b = "1011"
// 输出："10101"
func main() {
	var a, b string
	fmt.Println("Input binary a, binary b:")
	fmt.Scan(&a, &b)
	fmt.Println(addBinary(a, b))
}
