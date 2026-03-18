// 72. 编辑距离
/*
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。
你可以对一个单词进行如下三种操作：
插入一个字符
删除一个字符
替换一个字符



提示：
0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成
*/
package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}
	dp[0][0] = 0
	for i := 1; i <= len1; i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	for i := 1; i <= len2; i++ {
		dp[0][i] = dp[0][i-1] + 1
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // 当前两个字母相等, 无需操作, 这里不能计算为 min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len1][len2]
}

// 示例 1：
// 输入：word1 = "horse", word2 = "ros"
// 输出：3
// 解释：
// horse -> rorse (将 'h' 替换为 'r')
// rorse -> rose (删除 'r')
// rose -> ros (删除 'e')
//
// 示例 2：
// 输入：word1 = "intention", word2 = "execution"
// 输出：5
// 解释：
// intention -> inention (删除 't')
// inention -> enention (将 'i' 替换为 'e')
// enention -> exention (将 'n' 替换为 'x')
// exention -> exection (将 'n' 替换为 'c')
// exection -> execution (插入 'u')
func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	fmt.Println(minDistance(str1, str2))
}
