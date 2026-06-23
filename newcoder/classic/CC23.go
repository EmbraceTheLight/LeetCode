// CC23 最长的连续元素序列长度 同 LeetCode No.128 最长连续序列
/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

提示：
0 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func longestConsecutive(nums []int) int {
	n := len(nums)
	var ans int
	mp := make(map[int]bool)
	for i := 0; i < n; i++ {
		mp[nums[i]] = true
	}
	for num := range mp {
		smaller := num - 1
		if mp[smaller] == true { // 当前数值不是最小的, 继续寻找连续序列中最小的数, 从最小的数开始统计
			continue
		}

		count := 1
		bigger := num + 1
		for mp[bigger] == true {
			bigger = bigger + 1
			count++
		}
		ans = max(ans, count)
	}
	return ans
}

// 示例 1：
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 示例 2：
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9
//
// 示例 3：
// 输入：nums = [1,0,1,2]
// 输出：3
func main() {
	nums := CreateSlice[int]()
	fmt.Println(longestConsecutive(nums))
}
