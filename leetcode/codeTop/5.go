// 5. 最长回文子串
/*
给你一个字符串 s，找到 s 中最长的 回文 子串。

提示：
1 <= s.length <= 1000
s 仅由数字和英文字母组成
*/
package main

import "fmt"

func longestPalindrome(s string) string {
	var maxLen int
	var ans string
	n := len(s)
	for i := 0; i < n; i++ {
		l, r := i, i
		for l > 0 && s[l-1] == s[l] {
			l--
		}
		for r < n-1 && s[r+1] == s[r] {
			r++
		}
		for l > 0 && r < n-1 && s[l-1] == s[r+1] {
			l--
			r++
		}
		if r-l+1 > maxLen {
			ans = s[l : r+1]
			maxLen = r - l + 1
		}
	}
	return ans
}

// 示例 1：
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
//
// 示例 2：
// 输入：s = "cbbd"
// 输出："bb"
func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(longestPalindrome(str))
}
