// N 皇后 II
/*
n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。

示例 1：
输入：n = 4
输出：2
解释：如上图所示，4 皇后问题存在两个不同的解法。

示例 2：
输入：n = 1
输出：1

提示：

1 <= n <= 9
*/
package main

import "fmt"

func totalNQueens(n int) int {
	type pair struct {
		x, y int
	}
	var ans int
	queens := make([]pair, 0)

	check := func(x, y int, queens []pair) bool {
		for _, q := range queens {
			if q.x == x || q.y == y || x+y == q.x+q.y || x-y == q.x-q.y {
				return false
			}
		}
		return true
	}

	var dfs func(start int, count int)
	dfs = func(start int, count int) {
		if count == n {
			ans++
			return
		}
		for i := 0; i < n; i++ {
			if check(start, i, queens) == true {
				queens = append(queens, pair{start, i})
				dfs(start+1, count+1)
				queens = queens[:len(queens)-1]
			}
		}
	}

	dfs(0, 0)
	return ans
}

// Test Case1: n = 4	Output: 2
// Test Case2: n = 1	Output: 1
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(totalNQueens(n))
}
