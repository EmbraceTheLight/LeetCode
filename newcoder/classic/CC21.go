// CC21 包围区域 同 LeetCode No.130
/*
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' 组成，捕获 所有 被围绕的区域：

连接：一个单元格与水平或垂直方向上相邻的单元格连接。
区域：连接所有 'O' 的单元格来形成一个区域。
围绕：如果一个区域中的所有 'O' 单元格都不在棋盘的边缘，则该区域被包围。这样的区域 完全 被 'X' 单元格包围。
通过 原地 将输入矩阵中的所有 'O' 替换为 'X' 来 捕获被围绕的区域。你不需要返回任何值。

提示：
m == board.length
n == board[i].length
1 <= m, n <= 200
board[i][j] 为 'X' 或 'O'
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func solve(board [][]byte) {
	rows, cols := len(board), len(board[0])
	var dfs func(x, y int)

	// 反向思维: 统计有多少 'O' 是连通边界的, 将对应格子打上标记
	// 所有的不被包围的 O 都直接或间接与边界上的 O 相连
	// 剩下的仍为 'O' 的格子就是被包围的
	dfs = func(x, y int) {
		if x < 0 || x >= rows || y < 0 || y >= cols || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'A'
		dfs(x-1, y)
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
		return
	}

	// 从四个边界向中间收缩
	for i := 0; i < rows; i++ {
		dfs(i, 0)
		dfs(i, cols-1)
	}
	for i := 0; i < cols; i++ {
		dfs(0, i)
		dfs(rows-1, i)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}

// 示例 1：
// 输入：board = [['X','X','X','X'],['X','O','O','X'],['X','X','O','X'],['X','O','X','X']]
// 输出：[['X','X','X','X'],['X','X','X','X'],['X','X','X','X'],['X','O','X','X']]
//
// 示例 2：
// 输入：board = [['X']]
// 输出：[['X']]
func main() {
	board := CreateSlice2D[byte]()
	solve(board)
	for _, row := range board {
		for _, v := range row {
			fmt.Print(string(v))
		}
		fmt.Println()
	}
}
