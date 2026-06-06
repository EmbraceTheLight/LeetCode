/*
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]

输入：n = 1
输出：[["Q"]]

1 <= n <= 9
*/
package main

import "fmt"

func check51(x, y, n int, board []int) bool {
	for i := 0; i < n; i++ {
		if board[i] != -1 &&
			(x == i || y == board[i] || //同一行/列
				x+y == i+board[i] || y-x == board[i]-i) { //同一条斜线
			return false
		}
	}
	return true
}
func handleTmp(x, y int, tmp []string) []string {
	s := ""
	for i := 0; i < len(tmp); i++ {
		if i == y {
			s += "Q"
			continue
		}
		s += "."
	}
	tmp[x] = s
	return tmp
}
func dfs51(step, n int, board []int, tmp []string, ans *[][]string) {
	if step == n { //找到了一种解决方案
		str := make([]string, len(tmp))
		copy(str, tmp)
		*ans = append(*ans, str)
		return
	}

	for j := 0; j < n; j++ {
		if check51(step, j, n, board) {
			board[step] = j
			tmp = handleTmp(step, j, tmp)
			dfs51(step+1, n, board, tmp, ans)
			board[step] = -1
		}
	}

}
func solveNQueens(n int) [][]string {
	ans := make([][]string, 0)
	board := make([]int, n) //表示棋盘上的某一行的皇后位置，如：board[0] = 1 代表在(0,1)处有一个皇后
	for i := 0; i < n; i++ {
		board[i] = -1
	}
	tmp := make([]string, n)
	dfs51(0, n, board, tmp, &ans)
	return ans
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solveNQueens(n))
}
