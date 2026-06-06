package main

import "fmt"

func longestCommonSubsequenceR(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 0; i < len(dp); i++ {
		dp[i][0] = 0
	}

	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = 0
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}
func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	fmt.Println(longestCommonSubsequenceR(str1, str2))
}
