/*
516. 最长回文子序列
给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。

示例 1：
输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb" 。

示例 2：
输入：s = "cbbd"
输出：2
解释：一个可能的最长回文子序列为 "bb" 。

提示：
1 <= s.length <= 1000
s 仅由小写英文字母组成
*/
package main

import "fmt"

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n) //dp[i][j]:子串s[i,j]的最长回文子序列
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if s[j] == s[i] {
				dp[j][i] = dp[j+1][i-1] + 2
			} else {
				dp[j][i] = max(dp[j][i-1], dp[j+1][i])
			}

		}
	}
	return dp[0][n-1]
}
func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(longestPalindromeSubseq(s))
}
