// 16.04. 井字游戏
/*
设计一个算法，判断玩家是否赢了井字游戏。输入是一个 N x N 的数组棋盘，由字符" "，"X"和"O"组成，其中字符" "代表一个空位。
以下是井字游戏的规则：
玩家轮流将字符放入空位（" "）中。
第一个玩家总是放字符"O"，且第二个玩家总是放字符"X"。
"X"和"O"只允许放置在空位中，不允许对已放有字符的位置进行填充。
当有N个相同（且非空）的字符填充任何行、列或对角线时，游戏结束，对应该字符的玩家获胜。
当所有位置非空时，也算为游戏结束。
如果游戏结束，玩家不允许再放置字符。
如果游戏存在获胜者，就返回该游戏的获胜者使用的字符（"X"或"O"）；如果游戏以平局结束，则返回 "Draw"；如果仍会有行动（游戏未结束），则返回 "Pending"。

提示：
1 <= board.length == board[i].length <= 100
输入一定遵循井字棋规则
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func tictactoe(board []string) string {
	n := len(board)
	if checkMainDiagonal(board) == true {
		return string(board[0][0])
	}
	if checkSubDiagonal(board) == true {
		return string(board[0][n-1])
	}
	var hasEmpty bool
	// 行
	for i := 0; i < n; i++ {
		res, empty := checkRow(board, i)
		if hasEmpty == false && empty == true {
			hasEmpty = true
		}
		if res == true {
			return string(board[i][0])
		}

		res, empty = checkColumn(board, i)
		if hasEmpty == false && empty == true {
			hasEmpty = true
		}
		if res == true {
			return string(board[0][i])
		}
	}
	if hasEmpty == true {
		return "Pending"
	}
	return "Draw"
}

// 检查主对角线(左上 -> 右下)
func checkMainDiagonal(board []string) bool {
	char := board[0][0]
	if char == ' ' {
		return false
	}
	for i := 1; i < len(board); i++ {
		if board[i][i] != char {
			return false
		}
	}
	return true
}

// 检查副对角线
func checkSubDiagonal(board []string) bool {
	n := len(board)
	char := board[0][n-1]
	if char == ' ' {
		return false
	}
	for i := n - 2; i >= 0; i-- {
		if board[n-i-1][i] != char {
			return false
		}
	}
	return true
}

// 检查行
func checkRow(board []string, row int) (checkRes, hasSpace bool) {
	char := board[row][0]
	checkRes = true
	if char == ' ' {
		checkRes = false
	}
	for i := 1; i < len(board); i++ {
		if board[row][i] != char {
			checkRes = false
		}
		if board[row][i] == ' ' {
			hasSpace = true
		}
	}
	return
}

// 检查列
func checkColumn(board []string, col int) (checkRes, hasSpace bool) {
	char := board[0][col]
	checkRes = true
	if char == ' ' {
		checkRes = false
	}
	for i := 1; i < len(board); i++ {
		if board[i][col] != char {
			checkRes = false
		}
		if board[i][col] == ' ' {
			hasSpace = true
		}
	}
	return
}

// 示例 1：
// 输入： board = ["O X"," XO","X O"]
// 输出： "X"
//
// 示例 2：
// 输入： board = ["OOX","XXO","OXO"]
// 输出： "Draw"
// 解释： 没有玩家获胜且不存在空位
//
// 示例 3：
// 输入： board = ["OOX","XXO","OX "]
// 输出： "Pending"
// 解释： 没有玩家获胜且仍存在空位
func main() {
	fmt.Println("Input board:")
	board := CreateSlice[string]()
	fmt.Println(tictactoe(board))
}
