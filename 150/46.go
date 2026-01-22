// 46. 全排列
/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：
输入：nums = [1]
输出：[[1]]


提示：
1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func permute(nums []int) [][]int {
	n := len(nums)
	var ans [][]int
	visited := make(map[int]bool)
	var tmp []int

	var dfs func()
	dfs = func() {
		if len(tmp) == n {
			t := make([]int, n)
			copy(t, tmp)
			ans = append(ans, t)
		}

		for i := 0; i < n; i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			tmp = append(tmp, nums[i])
			dfs()
			visited[i] = false
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs()
	return ans
}

// Test Case1: [1,2,3]	Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// Test Case2: [0,1]	Output: [[0,1],[1,0]]
// Test Case3: [1]		Output: [[1]]
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(permute(nums))
}
