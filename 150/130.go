// 130. 被围绕的区域
/*
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' 组成，捕获 所有 被围绕的区域：
连接：一个单元格与水平或垂直方向上相邻的单元格连接。
区域：连接所有 'O' 的单元格来形成一个区域。
围绕：如果您可以用 'X' 单元格 连接这个区域，并且区域中没有任何单元格位于 board 边缘，则该区域被 'X' 单元格围绕。
通过 原地 将输入矩阵中的所有 'O' 替换为 'X' 来 捕获被围绕的区域。你不需要返回任何值。

示例 1：
输入：board = [['X','X','X','X'],['X','O','O','X'],['X','X','O','X'],['X','O','X','X']]
输出：[['X','X','X','X'],['X','X','X','X'],['X','X','X','X'],['X','O','X','X']]

示例 2：
输入：board = [['X']]
输出：[['X']]

提示：
m == board.length
n == board[i].length
1 <= m, n <= 200
board[i][j] 为 'X' 或 'O'
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 暴力解法
func solve(board [][]byte) {
	rows, cols := len(board), len(board[0])
	type rowAndCol struct {
		row, col int
	}
	rc := make([]rowAndCol, 0)
	var checkIsSurrounded func(x, y int) bool
	// checkIsSurrounded 检查 board[row][col] 是否被'X'围绕
	checkIsSurrounded = func(row, col int) bool {
		// 触及边界, 一定不会被围绕
		if row < 0 || row >= rows || col < 0 || col >= cols {
			return false
		}

		// 已经被标记为'X', 则不再检查
		// 或已经被检查过, 则不再检查
		if board[row][col] == 'X' || board[row][col] == '#' {
			return true
		}

		rc = append(rc, rowAndCol{row, col})
		// 标记一个中间状态, 防止重复遍历
		// 若果当前位置被围绕, 则标记为 '#', 用于后续处理
		board[row][col] = '#'
		isSurrounded := checkIsSurrounded(row+1, col) && checkIsSurrounded(row-1, col) && checkIsSurrounded(row, col-1) && checkIsSurrounded(row, col+1)
		// 若果当前位置没有被围绕, 则恢复状态
		if !isSurrounded {
			for _, item := range rc {
				board[item.row][item.col] = 'O'
			}
		}
		return isSurrounded
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				checkIsSurrounded(i, j)
				rc = make([]rowAndCol, 0)
			}
		}
	}

	// 将所有被围绕的'#'替换为'X'
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'X'
			}
		}
	}
}

// Test Case1:	[['X','X','X','X'],['X','O','O','X'],['X','X','O','X'],['X','O','X','X']]	Output:	[['X','X','X','X'],['X','X','X','X'],['X','X','X','X'],['X','O','X','X']]
// Test Case2:	[['X']]	Output:	[['X']]
func main() {
	board := pkg.CreateSlice2D[byte]()
	solve(board)
	for _, row := range board {
		for _, ch := range row {
			fmt.Printf("%c ", ch)
		}
		fmt.Println()
	}
}
