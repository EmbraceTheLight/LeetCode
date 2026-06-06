// 5. 最长回文子串
/*
给你一个字符串 s，找到 s 中最长的 回文 子串。

提示：
1 <= s.length <= 1000
s 仅由数字和英文字母组成
*/
package main

import "fmt"

func longestPalindromeR(s string) string {
	var tmp int
	ans := ""
	for i := 0; i < len(s); i++ {
		count := 1
		left, right := i-1, i+1
		for left >= 0 && s[left] == s[i] {
			left--
			count++
		}
		for right < len(s) && s[right] == s[i] {
			right++
			count++
		}
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
			count += 2
		}
		if tmp < count {
			tmp = count
			ans = s[left+1 : right]
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
	var s string
	fmt.Scan(&s)
	fmt.Println(longestPalindromeR(s))
}
