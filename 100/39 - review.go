package main

import (
	"fmt"
	"slices"
	"lc/100/pkg"
)

func dfs39R(candidates []int, target, start, sum int, now []int, ans *[][]int) {
	for i := start; i < len(candidates); i++ {
		sum += candidates[i]
		if sum > target {
			break
		}
		now = append(now, candidates[i])
		if sum == target {
			t := make([]int, len(now))
			copy(t, now)
			*ans = append(*ans, t)
			break
		}
		dfs39R(candidates, target, i, sum, now, ans)
		now = now[:len(now)-1]
		sum -= candidates[i]
	}
}

func combinationSum39R(candidates []int, target int) [][]int {
	var ans = make([][]int, 0)
	slices.Sort(candidates)
	dfs39R(candidates, target, 0, 0, []int{}, &ans)
	return ans
}
func main() {
	var target int
	fmt.Println("Input target: ")
	fmt.Scan(&target)
	sli := pkg.CreateIntSlice()
	fmt.Println(combinationSum39R(sli, target))
}
