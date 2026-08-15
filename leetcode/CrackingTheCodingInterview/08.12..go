// 08.12. 八皇后
/*
设计一种算法，打印 N 皇后在 N × N 棋盘上的各种摆法，其中每个皇后都不同行、不同列，也不在对角线上。
这里的“对角线”指的是所有的对角线，不只是平分整个棋盘的那两条对角线。
*/
package main

import "fmt"

func solveNQueens(n int) [][]string {
	var ans [][]string
	var pos [][]byte
	for i := 0; i < n; i++ {
		pos = append(pos, make([]byte, n))
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			pos[i][j] = '.'
		}
	}
	dfs0812(0, n, pos, &ans)
	return ans
}
func dfs0812(step int, n int, pos [][]byte, ans *[][]string) {
	if step == n {
		tmp := make([]string, n)
		for i := 0; i < n; i++ {
			tmp[i] = string(pos[i])
		}
		*ans = append(*ans, tmp)
		return
	}
	for i := 0; i < n; i++ {
		if isValid(step, i, n, pos) {
			pos[step][i] = 'Q'
			dfs0812(step+1, n, pos, ans)
			pos[step][i] = '.'
		}
	}
}
func isValid(x, y, n int, pos [][]byte) bool {
	// 检查同行是否已有皇后
	for i := 0; i < n; i++ {
		if i == y {
			continue
		}
		if pos[x][i] == 'Q' {
			return false
		}
	}

	// 检查同列是否已有皇后
	for i := 0; i < n; i++ {
		if i == x {
			continue
		}
		if pos[i][y] == 'Q' {
			return false
		}
	}

	// 检查两对角线是否已有皇后
	// 左上 → 右下
	for i, j := x+1, y+1; i < n && j < n; i, j = i+1, j+1 {
		if pos[i][j] == 'Q' {
			return false
		}
	}
	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if pos[i][j] == 'Q' {
			return false
		}
	}

	// 左下 → 右上
	for i, j := x+1, y-1; i < n && j >= 0; i, j = i+1, j-1 {
		if pos[i][j] == 'Q' {
			return false
		}
	}
	for i, j := x-1, y+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if pos[i][j] == 'Q' {
			return false
		}
	}
	return true
}

// 示例：
// 输入：4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// 解释：4 皇后问题存在如下两个不同的解法。
// [
// [".Q..",  // 解法 1
//
//	"...Q",
//	"Q...",
//	"..Q."],
//
// ["..Q.",  // 解法 2
//
//	"Q...",
//	"...Q",
//	".Q.."]
//
// ]
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solveNQueens(n))
}
