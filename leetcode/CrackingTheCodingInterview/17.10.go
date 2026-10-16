// 17.10. 主要元素
/*
数组中占比超过一半的元素称之为主要元素。给你一个 整数 数组，找出其中的主要元素。若没有，返回 -1。
请设计时间复杂度为 O(N) 、空间复杂度为 O(1) 的解决方案。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func majorityElement(nums []int) int {
	voteCount := 0
	var candidate int = nums[0]
	voteCount++
	for i := 1; i < len(nums); i++ {
		if voteCount == 0 {
			voteCount++
			candidate = nums[i]
			continue
		}
		if nums[i] == candidate {
			voteCount++
		} else {
			voteCount--
		}
	}
	var ans = -1
	var count = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		}
	}
	if count > len(nums)/2 {
		ans = candidate
	}
	return ans
}

// 示例 1：
// 输入：[1,2,5,9,5,9,5,5,5]
// 输出：5
//
// 示例 2：
// 输入：[3,2]
// 输出：-1
//
// 示例 3：
// 输入：[2,2,1,1,1,2,2]
// 输出：2
func main() {
	nums := CreateSlice[int]()
	fmt.Println(majorityElement(nums))
}
