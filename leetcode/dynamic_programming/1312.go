/*
1312. 让字符串成为回文串的最少插入次数 TODO:[hard]
给你一个字符串 s ，每一次操作你都可以在字符串的任意位置插入任意字符。
请你返回让 s 成为回文串的 最少操作次数 。
「回文串」是正读和反读都相同的字符串。

示例 1：
输入：s = "zzazz"
输出：0
解释：字符串 "zzazz" 已经是回文串了，所以不需要做任何插入操作。

示例 2：
输入：s = "mbadm"
输出：2
解释：字符串可变为 "mbdadbm" 或者 "mdbabdm" 。

示例 3：
输入：s = "leetcode"
输出：5
解释：插入 5 个字符后字符串变为 "leetcodocteel" 。

提示：

1 <= s.length <= 500
s 中所有字符都是小写字母。
*/
package main

import (
	"fmt"
	"strings"
)

// 思路：先求出字符串s和其逆序串revs的最长公共子序列的长度，这就是s的最长回文子序列长度。表示这些字符已经是回文的了，只需要让其他字符也回文就行了
// 将其他字符回文的方法就是再加一个相同字符。因此，有多少需要加的字符，这个数字就是答案。即字符串总长度减去LCS总长度
func minInsertions(s string) int {
	n := len(s)
	sb := strings.Builder{}
	for i := n - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	var revs = sb.String()

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == revs[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return n - dp[n][n]
}
func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(minInsertions(s))
}
