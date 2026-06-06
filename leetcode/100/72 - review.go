package main

import "fmt"

/*
以下注解取自 leetcode 用户 ⚡ 的评论：

情况1: 字符相等
如果 word1 的第 i 个字符（即 word1[i-1]）等于 word2 的第 j 个字符（即 word2[j-1]），那么这两个字符之间不需要任何编辑操作，因此 dp[i][j] = dp[i-1][j-1]。

情况2: 字符不等
如果 word1[i-1] != word2[j-1]，那么我们需要考虑三种编辑操作：
替换：我们可以将 word1[i-1] 替换为 word2[j-1]，这样两个字符串的最后一个字符就相同了，然后问题就转化为了将 word1 的前 i-1 个字符转换成 word2 的前 j-1 个字符，所以需要的操作数是 dp[i-1][j-1] + 1。
插入：在 word1 中插入 word2[j-1]，这相当于将 word1 的前 i 个字符转换成 word2 的前 j-1 个字符后，再插入一个字符，因此需要的操作数是 dp[i][j-1] + 1。（考虑前i个字符转换为j-1，再额外插入第j个字符）
删除：从 word1 中删除 word1[i-1]，这就把问题转化为将 word1 的前 i-1 个字符转换成 word2 的前 j 个字符，所需的操作数是 dp[i-1][j] + 1。（不用考虑删除的第i个字符，即前i-1个字符转换为前j个字符）

最终，dp[i][j] 应该取这三种情况中的最小值，因为我们要找的是最少编辑操作次数。
*/
func minDistanceR(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1) // dp[i][j]: cong word1[0,i] 到 word2[0,j] 的编辑距离
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}

	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func main() {
	var word1, word2 string
	fmt.Scan(&word1, &word2)
	fmt.Println(minDistanceR(word1, word2))
}
