// 08.03. 魔术索引
/*
魔术索引。 在数组A[0...n-1]中，有所谓的魔术索引，满足条件A[i] = i。
给定一个有序整数数组，编写一种方法找出魔术索引，若有的话，在数组A中找出一个魔术索引，如果没有，则返回-1。
若有多个魔术索引，返回索引值最小的一个。

说明：
nums长度在[1, 1000000]之间
此题为原书中的 Follow-up，即数组中可能包含重复元素的版本
*/
package main

import (
	"fmt"
	. "lc/pkg"
	"math"
)

func findMagicIndex(nums []int) int {
	var ans int = -1
	res := binarySearch0803(nums, 0, len(nums)-1)
	if res != math.MaxInt {
		return res
	}
	return ans
}

func binarySearch0803(nums []int, left, right int) int {
	if left > right {
		return math.MaxInt
	}
	var ans int
	mid := (left + right) / 2
	ans = binarySearch0803(nums, left, mid-1)
	if nums[mid] == mid {
		return mid
	} else if ans != math.MaxInt {
		return ans
	}
	return binarySearch0803(nums, mid+1, right)
}

// 示例 1：
// 输入：nums = [0, 2, 3, 4, 5]
// 输出：0
//
// 说明：0下标的元素为0
// 示例 2：
// 输入：nums = [1, 1, 1]
// 输出：1
func main() {
	nums := CreateSlice[int]()
	fmt.Println(findMagicIndex(nums))
}
