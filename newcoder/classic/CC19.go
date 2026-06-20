// CC19 LeetCode 132. 分割回文串 II
/*
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文串。
返回符合要求的 最少分割次数 。

提示：
1 <= s.length <= 2000
s 仅由小写英文字母组成
*/
package main

import (
	"fmt"
	"math"
)

func minCut(s string) int {
	n := len(s)
	dp := make([]int, n+1)

	// palindrome[i][j] s[i...j] 是否是回文串.
	// i > j 时, palindrome[i][j] == true. 这表示空字符串为回文串
	palindrome := make([][]bool, n)
	for i := 0; i < n; i++ {
		palindrome[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			palindrome[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			palindrome[i][j] = palindrome[i+1][j-1] && s[i] == s[j]
		}
	}

	for i := 1; i <= n; i++ {
		if palindrome[0][i-1] == true {
			dp[i] = 0
			continue
		}

		dp[i] = math.MaxInt
		for j := 1; j < i; j++ {
			if palindrome[j][i-1] == true {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n]
}

// 示例 1：
// 输入：s = "aab"
// 输出：1
// 解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
//
// 示例 2：
// 输入：s = "a"
// 输出：0
//
// 示例 3：
// 输入：s = "ab"
// 输出：1
func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(minCut(s))
}
