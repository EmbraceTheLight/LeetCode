// 43. 字符串相乘
/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。

提示：
1 <= num1.length, num2.length <= 200
num1 和 num2 只能由数字组成。
num1 和 num2 都不包含任何前导零，除了数字0本身。
*/
package main

import "fmt"

func multiply(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	var tmp = make([][]byte, n2)
	for i := 0; i < n2; i++ {
		tmp[i] = make([]byte, n1+n2) // 多一位, 用于保存必要的进位, 如 2 * 99 ==> len(tmp[0]) == 3
	}
	carry := 0
	i := n2 - 1
	idx := 0
	for ; i >= 0; i-- {
		idx2 := n1 + i
		for j := n1 - 1; j >= 0; j-- {
			mul := int(num1[j]-'0')*int(num2[i]-'0') + carry
			carry = mul / 10
			tmp[idx][idx2] = byte(mul%10 + '0')
			idx2--
		}
		if carry > 0 {
			tmp[idx][idx2] = byte(carry + '0')
			carry = 0
		}
		idx++
	}

	ans := make([]byte, n1+n2)
	idx = n1 + n2 - 1
	for col := n1 + n2 - 1; col >= 0; col-- {
		sum := 0
		for row := 0; row < n2; row++ {
			if tmp[row][col] == 0 {
				continue
			}
			sum += int(tmp[row][col] - '0')
		}
		sum += carry
		carry = sum / 10
		ans[idx] = byte(sum%10 + '0')
		idx--
	}
	if carry > 0 {
		ans[0] = byte(carry + '0')
	}
	ans = eatLeadingZero(ans)
	return string(ans)
}
func eatLeadingZero(b []byte) []byte {
	i := 0
	for ; i < len(b)-1; i++ {
		if b[i] != '0' {
			break
		}
	}
	return b[i:]
}

// 示例 1:
// 输入: num1 = "2", num2 = "3"
// 输出: "6"
//
// 示例 2:
// 输入: num1 = "123", num2 = "456"
// 输出: "56088"
func main() {
	var num1, num2 string
	fmt.Scan(&num1, &num2)
	fmt.Println(multiply(num1, num2))
}
