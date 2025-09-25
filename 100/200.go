/*
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1：
输入：grid = [

	["1","1","1","1","0"],
	["1","1","0","1","0"],
	["1","1","0","0","0"],
	["0","0","0","0","0"]

]
输出：1

示例 2：
输入：grid = [

	["1","1","0","0","0"],
	["1","1","0","0","0"],
	["0","0","1","0","0"],
	["0","0","0","1","1"]

]
输出：3

提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'
*/
package main

import (
	"fmt"
)

// 注意：这里的grid存储的是数字0和数字1，而并非字符0和字符1
func dfs200(i, j int, grid [][]byte, isVisit *[][]bool) {
	if (*isVisit)[i][j] || grid[i][j] == '0' {
		return
	}
	(*isVisit)[i][j] = true
	if i > 0 {
		dfs200(i-1, j, grid, isVisit)
	}
	if j > 0 {
		dfs200(i, j-1, grid, isVisit)
	}
	if i+1 < len(grid) {
		dfs200(i+1, j, grid, isVisit)
	}
	if j+1 < len(grid[0]) {
		dfs200(i, j+1, grid, isVisit)
	}
}
func numIslands(grid [][]byte) int {
	var ans int
	rows := len(grid)
	cols := len(grid[0])
	isVisit := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		var tmp = make([]bool, cols)
		isVisit[i] = tmp
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && isVisit[i][j] == false {
				dfs200(i, j, grid, &isVisit)
				ans++
			}
		}
	}
	return ans
}
func main() {
	var m int
	var grid [][]byte
	fmt.Println("Input rows:")
	fmt.Scanln(&m)
	var t byte
	for i := 0; i < m; i++ {
		var tmp = make([]byte, 0)
		fmt.Scan(&t)
		for t != 2 {
			tmp = append(tmp, t)
			fmt.Scan(&t)
		}
		grid = append(grid, tmp)
	}
	fmt.Println(numIslands(grid))
}
