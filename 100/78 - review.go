package main

import (
	"fmt"
	"lc/100/pkg"
)

/*
78. 子集
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的
子集（幂集）。

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
func dfs78R(nums []int, tmp []int, ans *[][]int, step, start int) {
	if len(tmp) == step {
		*ans = append(*ans, tmp)
		tmp = tmp[:len(tmp)-1]
		return
	}

	l := len(nums)
	for i := start; i < l; i++ {
		if l-i < step-len(tmp) { //剪枝：当剩余可添加元素个数小于需要的个数时，停止添加
			break
		}
		tmp = append(tmp, nums[i])
		dfs78R(nums, tmp, ans, step, i+1)
		tmp = tmp[:len(tmp)-1]
	}
}
func subsets78R(nums []int) [][]int {
	var ans = make([][]int, 0)
	ans = append(ans, []int{})
	for i := 1; i <= len(nums); i++ {
		dfs78R(nums, []int{}, &ans, i, 0)
	}
	return ans
}

func main() {
	sli := pkg.CreateIntSlice()
	fmt.Println(subsets78R(sli))
}
