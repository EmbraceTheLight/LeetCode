/*
40.组合总和 II
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用 一次 。
注意：解集不能包含重复的组合。

示例 1:
输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]

示例 2:
输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
]
*/
package main

import (
	"fmt"
	"sort"
	"lc/pkg"
)

func check(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
func dfs40(candidates, tmp []int, start, sum, target int, ans *[][]int) {
	if sum == target {
		*ans = append(*ans, append([]int{}, tmp...))
		return
	}

	for i := start; i < len(candidates); i++ {

		//重要！防止出现重复组合
		if i-1 >= start && candidates[i-1] == candidates[i] {
			continue
		}
		if sum+candidates[i] <= target {
			tmp = append(tmp, candidates[i])
			if sum+candidates[i] > target {
				return
			}
			dfs40(candidates, tmp, i+1, sum+candidates[i], target, ans)
			tmp = tmp[:len(tmp)-1]
		}
	}
}
func combinationSum2(candidates []int, target int) [][]int {
	var ans = make([][]int, 0)
	var tmp = make([]int, 0)
	sort.Ints(candidates)
	dfs40(candidates, tmp, 0, 0, target, &ans)

	return ans
}
func main() {
	fmt.Println("Input target:")
	var target int
	fmt.Scan(&target)
	nums := pkg.CreateIntSlice[int]()
	fmt.Println(combinationSum2(nums, target))

}
