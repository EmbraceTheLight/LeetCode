/*
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。

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

import "fmt"

// FIXME:建议使用第一种方法，即中心扩展法，而不是dp。因为方法一便于理解，效率也高
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	maxLen := 1
	ans := s[:1]
	var left, right int
	for i := 0; i < n; i++ {
		left, right = i-1, i+1

		for right < n && s[right] == s[right-1] {
			dp[i] = dp[i] + 1
			right++
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
			ans = s[right-dp[i] : right]
		}
		for left >= 0 && right < n {
			if s[left] != s[right] {
				break
			} else {
				dp[i] = dp[i] + 2
				if dp[i] > maxLen {
					maxLen = dp[i]
					ans = s[right-dp[i]+1 : right+1]
				}
				left, right = left-1, right+1
			}
		}
	}
	return ans
}

func longestPalindromeDP(s string) string {
	n := len(s)
	dp := make([][]bool, n) //dp[i][j]表示字符串s[i][j]是否是回文串
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	var begin, end = 0, 0
	//列遍历，然后再行遍历
	for i := 0; i < n; i++ {
		dp[i][i] = true
		for j := 0; j < i; j++ {
			if i-j == 1 && s[i] == s[j] {
				dp[j][i] = true
			} else if dp[j+1][i-1] == true && s[i] == s[j] {
				dp[j][i] = true
			}
			if dp[j][i] == true && i-j > end-begin {
				begin = j
				end = i
			}
		}

	}

	return s[begin : end+1]
}
func main() {
	var s string
	fmt.Scan(&s)
	//fmt.Println(longestPalindrome(s))
	fmt.Println(longestPalindromeDP(s))
}
