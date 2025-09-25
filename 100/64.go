/*
64.最小路径和
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。

示例 1：
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7

示例 2：
输入：grid = [[1,2,3],[4,5,6]]
输出：12

m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 200
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func minPathSum(grid [][]int) int {
	//var ans int
	var dp = make([][]int, 0) //dp[i][j]:走到第i行第j列对应坐标的最小路径和
	rows := len(grid)
	cols := len(grid[0])
	for i := 0; i < rows; i++ {
		var tmp = make([]int, cols)
		dp = append(dp, tmp)
	}
	dp[0][0] = grid[0][0]
	for j := 1; j < cols; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]

	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[rows-1][cols-1]
}
func main() {
	grid := pkg.CreateIntSlice2()
	fmt.Println(minPathSum(grid))
}
