/*
221. 最大正方形

在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

示例 1：
输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：4

示例 2：
输入：matrix = [["0","1"],["1","0"]]
输出：1

示例 3：
输入：matrix = [["0"]]
输出：0
*/
package main

import (
	"fmt"
	pkg "lc/100/pkg"
)

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m) //dp[i][j]表示到达matrix[i][j]时的最大正方形个数
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	var ans = 0
	for i := 0; i < m; i++ {
		dp[i][0] = int(matrix[i][0] - '0')
		if ans == 0 && dp[i][0] == 1 {
			ans = 1
		}

	}
	for i := 0; i < n; i++ {
		dp[0][i] = int(matrix[0][i] - '0')
		if ans == 0 && dp[0][i] == 1 {
			ans = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				ans = max(ans, dp[i][j])
			}

		}
	}
	return ans * ans
}
func main() {
	sli := pkg.CreateByteSlice2()
	fmt.Println(maximalSquare(sli))
}
