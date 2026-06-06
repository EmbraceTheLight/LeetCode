// 221. 最大正方形
/*
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 300
matrix[i][j] 为 '0' 或 '1'
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func maximalSquareR(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows+1)
	dp[0] = make([]int, cols+1)
	for i := 1; i <= rows; i++ {
		dp[i] = make([]int, cols+1)
		for j := 1; j <= cols; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1
			}
		}
	}
	var ans int
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans * ans
}

// 示例 1：
// 输入：matrix = [['1','0','1','0','0'],['1','0','1','1','1'],['1','1','1','1','1'],['1','0','0','1','0']]
// 输出：4
//
// 示例 2：
// 输入：matrix = [['0','1'],['1','0']]
// 输出：1
//
// 示例 3：
// 输入：matrix = [['0']]
// 输出：0
func main() {
	matrix := pkg.CreateSlice2D[byte]()
	fmt.Println(maximalSquareR(matrix))
}
