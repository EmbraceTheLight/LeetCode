/*
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。
如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

示例 2：
输入：s = "cbbd"
输出："bb"

提示：
1 <= s.length <= 1000
s 仅由数字和英文字母组成
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func longestPalindrome(s string) string {
	var dp = make([][]bool, 0)
	var maxLen = 1 //记录最长回文子串长度
	var start = 0  //记录最长回文子串长度
	l := len(s)
	for i := 0; i < l; i++ {
		tmp := make([]bool, l)
		dp = append(dp, tmp)
		dp[i][i] = true
	}
	for i := 1; i < l; i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] && (dp[j+1][i-1] == true || j+1 > i-1) {
				if i-j+1 > maxLen {
					start = j
					maxLen = i - j + 1
				}
				dp[j][i] = true
			}
		}
	}

	return s[start : start+maxLen]
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	b, _, _ := reader.ReadLine()
	fmt.Println(longestPalindrome(string(b)))
}
