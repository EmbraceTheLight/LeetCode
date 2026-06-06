/*
695.岛屿的最大面积
给你一个大小为 m x n 的二进制矩阵 grid 。
岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
岛屿的面积是岛上值为 1 的单元格的数目。
计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
m == grid.length
n == grid[i].length
1 <= m, n <= 50
grid[i][j] 为 0 或 1
*/
package main

import (
	"fmt"
)

// 每调用一次该函数，就获得一块岛屿的面积
func dfs695(i, j int, grid [][]int) int {
	if (i < 0 || i >= len(grid)) || (j < 0 || j >= len(grid[0])) || grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = 2 //标记为已扫描

	return 1 + //返回自身面积和周围四块的面积
		dfs695(i+1, j, grid) +
		dfs695(i, j+1, grid) +
		dfs695(i-1, j, grid) +
		dfs695(i, j-1, grid)
}

func maxAreaOfIsland(grid [][]int) int {
	var ans int
	rows := len(grid)
	cols := len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				ans = max(ans, dfs695(i, j, grid))
			}
		}
	}
	return ans
}
func main() {
	var m int
	var grid [][]int
	fmt.Println("Input rows:")
	fmt.Scanln(&m)
	var t int
	for i := 0; i < m; i++ {
		var tmp = make([]int, 0)
		fmt.Scan(&t)
		for t != -1 {
			tmp = append(tmp, t)
			fmt.Scan(&t)
		}
		grid = append(grid, tmp)
	}
	fmt.Println(maxAreaOfIsland(grid))
}
