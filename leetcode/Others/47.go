/*
47.全排列 II
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

示例 1：
输入：nums = [1,1,2]
输出：
[[1,1,2],

	[1,2,1],
	[2,1,1]]

示例 2：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

提示：
1 <= nums.length <= 8
-10 <= nums[i] <= 10
*/
package main

import (
	"WorkSpace/src/LeetCode/100/pkg"
	"fmt"
	"sort"
)

var bools []bool

func dfs47(tmp, nums []int, ans *[][]int) {
	if len(tmp) == len(nums) {
		*ans = append(*ans, append([]int{}, tmp...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && bools[i-1] == false { //去重，如果前一个元素未被使用过，则说明该重复元素已经被前一个元素使用过了，跳过该元素
			continue
		}
		if bools[i] == true {
			continue
		}
		tmp = append(tmp, nums[i])
		bools[i] = true
		dfs47(tmp, nums, ans)
		bools[i] = false
		tmp = tmp[:len(tmp)-1]

	}
}
func permuteUnique(nums []int) [][]int {
	var ans = make([][]int, 0)
	bools = make([]bool, len(nums))
	sort.Ints(nums)
	dfs47([]int{}, nums, &ans)
	return ans
}
func main() {
	nums := pkg.CreateIntSlice()
	fmt.Println(permuteUnique(nums))
}
