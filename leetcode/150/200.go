// 200. 岛屿数量
/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1：
输入：grid = [
  ['1','1','1','1','0'],
  ['1','1','0','1','0'],
  ['1','1','0','0','0'],
  ['0','0','0','0','0']
]
输出：1

示例 2：
输入：grid = [
  ['1','1','0','0','0'],
  ['1','1','0','0','0'],
  ['0','0','1','0','0'],
  ['0','0','0','1','1']
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
	"lc/pkg"
)

func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	var ans int
	var dfs200 func(grid [][]byte, x, y int)
	dfs200 = func(grid [][]byte, x, y int) {
		if (x < 0 || x > rows-1) || (y < 0 || y > cols-1) {
			return
		}
		if grid[x][y] == '1' {
			grid[x][y] = '2'
			dfs200(grid, x, y+1)
			dfs200(grid, x+1, y)
			dfs200(grid, x, y-1)
			dfs200(grid, x-1, y)
		}
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				ans++
				dfs200(grid, row, col)
			}
		}
	}
	return ans
}

// Test Case1: [['1','1','1','1','0'],['1','1','0','1','0'],['1','1','0','0','0'],['0','0','0','0','0']]
// Output: 1

// Test Case2: [['1','1','0','0','0'],['1','1','0','0','0'],['0','0','1','0','0'],['0','0','0','1','1']]
// Output: 3
func main() {
	grid := pkg.CreateSlice2D[byte]()
	fmt.Println(numIslands(grid))
}
