// 212. 单词搜索 II
/*
给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词 。
单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，
其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。


示例 1：
输入：board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
输出：["eat","oath"]

示例 2：
输入：board = [["a","b"],["c","d"]], words = ["abcb"]
输出：[]

提示：
m == board.length
n == board[i].length
1 <= m, n <= 12
board[i][j] 是一个小写英文字母
1 <= words.length <= 3 * 10^4
1 <= words[i].length <= 10
words[i] 由小写英文字母组成
words 中的所有字符串互不相同
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 普通 dfs, 超时
func findWords(board [][]byte, words []string) []string {
	m, n := len(board), len(board[0])
	wordMp := make(map[byte][][2]int)
	visited := make(map[int]map[int]bool)
	for i := 0; i < m; i++ {
		visited[i] = make(map[int]bool)
		for j := 0; j < n; j++ {
			wordMp[board[i][j]] = append(wordMp[board[i][j]], [2]int{i, j})
		}
	}
	var ans []string
	for i := 0; i < len(words); i++ {
		if findHelper(wordMp, words[i], board, visited) {
			ans = append(ans, words[i])
		}
	}
	return ans
}

func findHelper(wordMp map[byte][][2]int, str string, board [][]byte, visited map[int]map[int]bool) bool {
	ch := str[0]

	var dfs func(x, y int, idx int, str string, board [][]byte) bool
	dfs = func(x, y int, idx int, str string, board [][]byte) bool {
		if idx == len(str) {
			return true
		}
		if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || visited[x][y] == true || board[x][y] != str[idx] {
			return false
		}
		visited[x][y] = true
		defer func() { visited[x][y] = false }()
		return dfs(x-1, y, idx+1, str, board) ||
			dfs(x, y+1, idx+1, str, board) ||
			dfs(x+1, y, idx+1, str, board) ||
			dfs(x, y-1, idx+1, str, board)
	}

	start := wordMp[ch]
	for _, startIdx := range start {
		ok := dfs(startIdx[0], startIdx[1], 0, str, board)
		if ok {
			return true
		}
	}

	return false
}

//type trie struct {
//	children [26]*trie
//	isEnd    bool
//}

// Test Case1:  [['o','a','a','n'],['e','t','a','e'],['i','h','k','r'],['i','f','l','v']]	["oath","pea","eat","rain"]	Output:	["eat","oath"]
// Test Case2:	[['a','b'],['c','d']]	["abcb"]		Output:	[]
// Test Case3:	[['o','a','b','n'],['o','t','a','e'],['a','h','k','r'],['a','f','l','v']]	["oa","oaa"]	Output: ["oa","oaa"]
func main() {
	fmt.Println("Input board:")
	board := pkg.CreateSlice2D[byte]()
	fmt.Println("Input words:")
	words := pkg.CreateSlice[string]()
	fmt.Println(findWords(board, words))
}
