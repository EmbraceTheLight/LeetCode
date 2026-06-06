// 77. 组合
/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。

示例 1：
输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

示例 2：
输入：n = 1, k = 1
输出：[[1]]

提示：

1 <= n <= 20
1 <= k <= n
*/
package main

import "fmt"

func combine(n int, k int) [][]int {
	var ans [][]int
	var tmp []int
	var dfs func(start, n, k int)
	dfs = func(start, n, k int) {
		if k == 0 {
			arr := make([]int, len(tmp))
			copy(arr, tmp)
			ans = append(ans, arr)
			return
		}
		for i := start; i <= n-k+1; i++ {
			tmp = append(tmp, i)
			dfs(i+1, n, k-1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(1, n, k)
	return ans
}

// Test Case1: n=4, k=2	Output: [[2,4],[3,4],[2,3],[1,2],[1,3],[1,4]]
// Test Case2: n=1, k=1	Output: [[1]]
func main() {
	fmt.Println("Input n,k:")
	var n, k int
	fmt.Scan(&n, &k)
	fmt.Println(combine(n, k))
}
