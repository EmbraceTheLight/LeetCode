package main

import (
	"fmt"
	"lc/pkg"
)

/*
46. 全排列

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

func dfs46R(nums []int, mp map[int]bool, tmp []int, ans *[][]int) {
	l := len(nums)
	if len(tmp) == l {
		var t = make([]int, l)
		copy(t, tmp)
		*ans = append(*ans, t)
		return
	}

	for i := 0; i < l; i++ {
		if mp[i] == true {
			continue
		}
		tmp = append(tmp, nums[i])
		mp[i] = true
		dfs46R(nums, mp, tmp, ans)
		mp[i] = false
		tmp = tmp[:len(tmp)-1]
	}
}

func permute46R(nums []int) [][]int {
	var mp = make(map[int]bool)
	var ans = make([][]int, 0)
	var tmp = make([]int, 0)
	dfs46R(nums, mp, tmp, &ans)
	return ans
}
func main() {
	sli := pkg.CreateIntSlice[int]()
	fmt.Println(permute46R(sli))
}
