// 79. 单词搜索
/*
给定一个 m x n 二维字符网格 board 和一个字符串单词 word。
如果 word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母不允许被重复使用。


示例 1：
输入：board = , word = "ABCCED"
输出：true

示例 2：[['A','B','C','E'],['S','F','C','S'],['A','D','E','E']]
输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "SEE"
输出：true

示例 3：
输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCB"
输出：false

提示：

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board 和 word 仅由大小写英文字母组成
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func exist(board [][]byte, word string) bool {
	type pos struct {
		x, y int
	}
	mp := make(map[byte][]pos)
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mp[board[i][j]] = append(mp[board[i][j]], pos{x: i, y: j})
		}
	}

	start := word[0]
	var isExist func(x, y int, idx int) bool
	isExist = func(x, y int, idx int) bool {
		if idx == len(word) {
			return true
		}
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] == '@' || board[x][y] != word[idx] {
			return false
		}

		ch := board[x][y]
		board[x][y] = '@'
		defer func() { board[x][y] = ch }()
		return isExist(x+1, y, idx+1) ||
			isExist(x-1, y, idx+1) ||
			isExist(x, y+1, idx+1) ||
			isExist(x, y-1, idx+1)
	}

	for _, p := range mp[start] {
		if isExist(p.x, p.y, 0) {
			return true
		}
	}
	return false
}

// Test Case1: board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCCED"	Output: true
// Test Case2: board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "SEE"	Output: true
// Test Case3: board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCB"	Output: false
func main() {
	var word string
	fmt.Println("Input word:")
	fmt.Scan(&word)
	fmt.Println("Input board")
	board := pkg.CreateSlice2D[byte]()
	fmt.Println(exist(board, word))
}
