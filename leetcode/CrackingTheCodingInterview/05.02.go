// 05.02. 二进制数转字符串
/*
二进制数转字符串。给定一个介于 0 和 1 之间的实数（如 0.72），类型为 double，打印它的二进制表达式。如果该数字无法精确地用 32 位以内的二进制表示，则打印“ERROR”。

提示：
32位包括输出中的 "0." 这两位。
题目保证输入用例的小数位数最多只有 6 位
*/
package main

import (
	"fmt"
	"math"
	"strings"
)

func printBin(num float64) string {
	var ans strings.Builder
	exponent := -1
	for num != 0 && ans.Len() <= 32 {
		expoValue := math.Pow(2, float64(exponent))
		if expoValue > num {
			ans.WriteByte('0')
		} else {
			num -= expoValue
			ans.WriteByte('1')
		}
		exponent--
	}
	if ans.Len() > 32 {
		return "ERROR"
	} else {
		return "0." + ans.String()
	}
}

// 示例 1：
// 输入：0.625
// 输出："0.101"
//
// 示例 2：
// 输入：0.1
// 输出："ERROR"
// 提示：0.1 无法被二进制准确表示
func main() {
	var num float64
	fmt.Scan(&num)
	fmt.Println(printBin(num))
}
