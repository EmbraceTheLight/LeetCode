// 16.24. 数对和
/*
设计一个算法，找出数组中两数之和为指定值的所有整数对。一个数只能属于一个数对。

提示：
nums.length <= 100000
-10^5 <= nums[i], target <= 10^5
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func pairSums(nums []int, target int) [][]int {
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	var ans [][]int
	for _, v := range nums {
		if mp[v] == 0 {
			continue
		}
		if v == target-v && mp[v] == 1 {
			continue
		}
		if mp[target-v] > 0 {
			mp[target-v]--
			mp[v]--
			ans = append(ans, []int{v, target - v})
		}
	}
	return ans
}

// 示例 1：
// 输入：nums = [5,6,5], target = 11
// 输出：[[5,6]]
//
// 示例 2：
// 输入：nums = [5,6,5,6], target = 11
// 输出：[[5,6],[5,6]]
func main() {
	var target int
	fmt.Println("Input target:")
	fmt.Scan(&target)
	nums := CreateSlice[int]()
	fmt.Println(pairSums(nums, target))
}
