/*
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
m == matrix.length
n == matrix[i].length
1 <= m, n <= 300
matrix[i][j] 为 '0' 或 '1'
*/
package main

import (
	"fmt"
)

func maximalSquare(matrix [][]byte) int {
	var ans int
	var rows = len(matrix)
	var cols = len(matrix[0])
	var dp = make([][]int, rows)
	for i := 0; i < rows; i++ {
		var tmp = make([]int, cols)
		tmp[0] = int(matrix[i][0])
		dp[i] = tmp
		if ans == 0 && matrix[i][0] == 1 {
			ans = 1
		}
	}
	for i := 0; i < cols; i++ {
		dp[0][i] = int(matrix[0][i])
		if ans == 0 && dp[0][i] == 1 {
			ans = 1
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 1 {
				if matrix[i-1][j] == 1 && matrix[i][j-1] == 1 && matrix[i-1][j-1] == 1 {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				} else {
					dp[i][j] = 1
				}
				ans = max(ans, dp[i][j])
			}
		}
	}
	return ans * ans

}
func main() {
	var m int
	var n int
	fmt.Println("Input m,n:")
	fmt.Scan(&m)
	fmt.Scan(&n)
	var in byte
	var matrix = make([][]byte, m)
	for i := 0; i < m; i++ {
		var tmp = make([]byte, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&in)
			tmp[j] = in
		}
		matrix[i] = tmp
	}
	fmt.Println(maximalSquare(matrix))
}
