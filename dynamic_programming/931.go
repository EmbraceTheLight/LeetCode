/*
931.下降路径最小和
给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。

下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。
具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。

示例 1：
输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
输出：13
解释：如图所示，为和最小的两条下降路径

示例 2：
输入：matrix = [[-19,57],[-40,-5]]
输出：-59
解释：如图所示，为和最小的下降路径

提示：
n == matrix.length == matrix[i].length
1 <= n <= 100
-100 <= matrix[i][j] <= 100
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func minFallingPathSum(matrix [][]int) int {
	m := len(matrix)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		dp[0][i] = matrix[0][i]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = min(dp[i-1][0], dp[i-1][1]) + matrix[i][0]
		for j := 1; j < m-1; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
		}
		dp[i][m-1] = min(dp[i-1][m-2], dp[i-1][m-1]) + matrix[i][m-1]
	}
	ans := dp[m-1][0]
	for i := 1; i < m; i++ {
		ans = min(ans, dp[m-1][i])
	}
	return ans
}
func main() {
	sli := pkg.CreateIntSlice2[int]()
	fmt.Println(minFallingPathSum(sli))
}
