/*
78.子集
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

示例 2：
输入：nums = [0]
输出：[[],[0]]

提示：
1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func dfs78(sub, nums []int, ans *[][]int, start int) {
	*ans = append(*ans, append([]int{}, sub...))
	for i := start; i < len(nums); i++ {
		sub = append(sub, nums[i])
		dfs78(sub, nums, ans, i+1)
		sub = sub[:len(sub)-1]
	}
}
func subsets(nums []int) [][]int {
	var ans = make([][]int, 0)
	ans = append(ans, []int{})
	var sub = make([]int, 0)
	for i := 0; i < len(nums); i++ {
		sub = append(sub, nums[i])
		dfs78(sub, nums, &ans, i+1)
		sub = sub[:len(sub)-1]
	}
	return ans
}
func main() {
	nums := pkg.CreateIntSlice()
	fmt.Println(subsets(nums))
}
