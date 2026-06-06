// 6. Z 字形变换
/*
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);

示例 1：
输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"

示例 2：
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"

解释：
P     I    N
A   L S  I G
Y A   H R
P     I

示例 3：
输入：s = "A", numRows = 1
输出："A"
*/
package main

import (
	"fmt"
	"lc/pkg"
	"strings"
)

// * 自己推导出来的，值得鼓励！
// 这道题可以找规律实现。（N 字形可能更容易理解，Z 字形容易引起歧义）
// 规律：
// 记 idx 为当前行的起始位置在原字符串中的索引，则：
// 当遍历到第一行或最后一行时，斜向没有字符，下一个字符索引 nextIdx 为 当前字符位置 + 2*numRows - 2
// 当遍历到中间行时，斜向有字符，因此要设置两个临时下标 nextIdx1（表示斜向字符的下标）和 nextIdx2（表示竖向字符的下标）：
// nextIdx1 初始化为行初始位置 idx + (numRows-line)*2，nextIdx2 初始化为 idx + 2*numRows - 2。之后在迭代中，它们均取 2*numRows - 2 作为偏移量。
// 依次迭代，直到遍历完所有字符。
func convert(s string, numRows int) string {
	n := len(s)
	if numRows == 1 {
		return s
	}
	var sb strings.Builder
	idx := 0
	for line := 1; line <= numRows && idx < n; line++ {
		if line == 1 || line == numRows {
			sb.WriteByte(s[idx])
			nextIdx := idx + 2*numRows - 2
			for nextIdx < n {
				sb.WriteByte(s[nextIdx])
				nextIdx = nextIdx + 2*numRows - 2
			}
		} else {
			sb.WriteByte(s[idx])
			nextIdx1 := idx + (numRows-line)*2
			nextIdx2 := idx + 2*numRows - 2
			for nextIdx1 < n && nextIdx2 < n {
				sb.WriteByte(s[nextIdx1])
				sb.WriteByte(s[nextIdx2])
				nextIdx1 = nextIdx1 + 2*numRows - 2
				nextIdx2 = nextIdx2 + 2*numRows - 2
			}
			if nextIdx1 < n {
				sb.WriteByte(s[nextIdx1])
			}
			if nextIdx2 < n {
				sb.WriteByte(s[nextIdx2])
			}
		}
		idx = line
	}
	return sb.String()
}

// Test Case 1: 3 PAYPALISHIRING  Output: PAHNAPLSIIGYIR
// Test Case 2: 4, PAYPALISHIRING  Output: PINALSIGYAHRPI
// Test Case 3: 1, A  Output: A
func main() {
	fmt.Println("Input numRows:")
	var numRows int
	fmt.Scan(&numRows)
	s := pkg.CreateString()
	fmt.Println(convert(s, numRows))
}
