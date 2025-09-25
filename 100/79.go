/*
79. 单词搜索
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
示例1：
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true

示例2：
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
输出：true

示例3：
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
输出：false

m == board.length
n = board[i].length
1 <= m, n <= 6
1 <= word.length <= 15
board 和 word 仅由大小写英文字母组成
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var directions = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func check79(m, n, x, y int, visit map[int]map[int]struct{}) bool {
	var isok = false
	if x >= 0 && x < m && y >= 0 && y < n {
		if _, ok := visit[x][y]; !ok {
			isok = true
		}
	}
	return isok
}
func find(m, n, x, y, count int, board [][]byte, visit map[int]map[int]struct{}, word string) bool {

	if board[x][y] != word[count] { // 剪枝
		return false
	}
	if count == len(word)-1 {
		return true
	}
	if _, ok := visit[x]; !ok {
		visit[x] = make(map[int]struct{})
	}
	visit[x][y] = struct{}{}
	defer delete(visit[x], y)
	for i := 0; i < 4; i++ {
		nx, ny := x+directions[i][0], y+directions[i][1]
		if check79(m, n, nx, ny, visit) {
			ok := find(m, n, nx, ny, count+1, board, visit, word)
			if ok {
				return true
			}
		}
	}

	return false
}
func exist(board [][]byte, word string) bool {

	visit := make(map[int]map[int]struct{})
	x, y := 0, 0
	row, col := len(board), len(board[0])
	for x = 0; x < row; x++ {
		for y = 0; y < col; y++ {
			if check79(row, col, x, y, visit) {
				ok := find(row, col, x, y, 0, board, visit, word)
				if ok {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	var m, n int
	fmt.Println("Input row、cow of the board:")
	fmt.Scan(&m, &n)
	var str string
	fmt.Println("Input target word:")
	fmt.Scan(&str)
	var board = make([][]byte, 0)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < m; i++ {
		fmt.Println("Input row", i+1, "of the board:")
		input, _ := reader.ReadString('\n')
		board = append(board, []byte(strings.ReplaceAll(input, " ", ""))[:n])
	}
	fmt.Println(exist(board, str))
}
