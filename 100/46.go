/*
46.全排列
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

func dfs46(nums, tmp []int, ans *[][]int, isVisit []bool) {
	if len(tmp) == len(nums) {
		var ins = make([]int, len(tmp))
		copy(ins, tmp)
		*ans = append(*ans, ins)
		return
	}
	for j := 0; j < len(nums); j++ {
		if isVisit[j] == true {
			continue
		}
		tmp = append(tmp, nums[j])
		isVisit[j] = true
		dfs46(nums, tmp, ans, isVisit)
		isVisit[j] = false
		tmp = tmp[:len(tmp)-1]
	}

}
func permute(nums []int) [][]int {
	var ans [][]int
	var isVisit = make([]bool, len(nums))
	var tmp = make([]int, 0)
	dfs46(nums, tmp, &ans, isVisit)
	return ans
}
func main() {
	nums := pkg.CreateIntSlice()
	fmt.Println(permute(nums))
}
