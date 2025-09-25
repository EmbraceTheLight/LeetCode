/*
189. 轮转数组

给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

示例 2:
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]

提示：
1 <= nums.length <= 10^5
-2^31 <= nums[i] <= 2^31 - 1
0 <= k <= 10^5
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

//func rotate48(nums []int, k int) {
//	if k == 0 {
//		return
//	}
//
//	//var cnt = 0
//	var length = len(nums)
//	var move = k % length
//
//	nums = append(nums, nums...)
//	fmt.Println(nums)
//	start := length - move
//	//nums = nums[start : start+length]
//	for i := 0; i < length; i++ {
//		nums[i] = nums[start+i]
//	}
//	fmt.Println(nums)
//}

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}

	//var cnt = 0
	var length = len(nums)
	k = k % length
	var count int

	for i := 0; count < length; i++ {
		curIdx := i
		curVal := nums[curIdx]
		for ok := true; ok; ok = curIdx != i {
			newIdx := (curIdx + k) % length
			nValue := nums[newIdx]

			nums[newIdx] = curVal

			curIdx = newIdx
			curVal = nValue
			count++
		}
	}
}

func main() {
	fmt.Println("Input a rotate48 number k:")
	var k int
	fmt.Scan(&k)
	nums := pkg.CreateIntSlice()
	rotate(nums, k)
	fmt.Println(nums)
}
