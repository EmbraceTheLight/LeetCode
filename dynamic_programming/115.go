/*
115. 不同的子序列 [hard]
给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数，结果需要对 10^9 + 7 取模。

示例 1：
输入：s = "rabbbit", t = "rabbit"
输出：3
解释：
如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
rabbbit
rabbbit
rabbbit

示例 2：
输入：s = "babgbag", t = "bag"
输出：5
解释：
如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。
babgbag
babgbag
babgbag
babgbag
babgbag

提示：

1 <= s.length, t.length <= 1000
s 和 t 由英文字母组成
*/
package main

import (
	"fmt"
)

func numDistinct(s string, t string) int {
	const mod = 1000000007
	m, n := len(s), len(t)
	if n > m {
		return 0
	}
	dp := make([][]int, n) //dp[i][j]:从t[0,i]到s[0,j]的不同的子序列的个数
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	if s[0] == t[0] {
		dp[0][0] = 1
	}
	for i := 1; i < m; i++ {
		if t[0] == s[i] {
			dp[0][i] = dp[0][i-1] + 1
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i := 1; i < n; i++ {
		for j := i; j < m; j++ {
			if s[j] == t[i] {
				//当s[j]==t[i]时，将 {满足t[0,i]是s[0,j-1]的子序列的个数(s[j]不与t[i]匹配)}与{满足t[0,i-1]是s[0,j-1]的子序列个数(s[j]与t[i]匹配)} 相加
				//s[j]不与t[i]匹配，即dp[i][j-1],表示的是不考虑当前相等的情况（不把s[j]算进来），有多少匹配的
				//s[j]与t[i]匹配，即dp[i-1][j-1],表示的是考虑当前相等的情况（将s[j]算进来），有多少匹配的
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[n-1][m-1] % mod
}
func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	fmt.Println(numDistinct(s1, s2))
}
