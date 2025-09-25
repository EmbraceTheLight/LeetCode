package main

import (
	"fmt"
	"lc/pkg"
)

func dfs200R(curRow, curCol int, grid [][]byte) {
	if grid[curRow][curCol] != '1' {
		return
	}
	rows := len(grid)
	cols := len(grid[0])
	grid[curRow][curCol] = '2'
	if curRow+1 < rows {
		dfs200R(curRow+1, curCol, grid)
	}
	if curRow-1 >= 0 {
		dfs200R(curRow-1, curCol, grid)
	}
	if curCol+1 < cols {
		dfs200R(curRow, curCol+1, grid)
	}
	if curCol-1 >= 0 {
		dfs200R(curRow, curCol-1, grid)
	}
}
func numIslandsR(grid [][]byte) int {
	var ans = 0
	rows := len(grid)
	cols := len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				dfs200R(i, j, grid)
				ans++
			}
		}
	}
	return ans
}
func main() {
	fmt.Println(numIslandsR(pkg.CreateByteSlice2()))
}
