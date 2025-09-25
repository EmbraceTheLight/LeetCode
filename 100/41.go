/*
41. 缺失的第一个正数
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

示例 1：
输入：nums = [1,2,0]
输出：3
解释：范围 [1,2] 中的数字都在数组中。

示例 2：
输入：nums = [3,4,-1,1]
输出：2
解释：1 在数组中，但 2 没有。

示例 3：
输入：nums = [7,8,9,11,12]
输出：1
解释：最小的正数 1 没有出现。

提示：
1 <= nums.length <= 10^5
-2^31 <= nums[i] <= 2^31 - 1
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

// 原地哈希
func firstMissingPositive(nums []int) int {
	N := len(nums)
	//只要位于1--N的数，不在这个范围的不考虑
	for i := 0; i < N; i++ {
		for nums[i] > 0 && nums[i] <= N && nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1] //将nums[i]移动到nums[nums[i]-1]的位置，原nums[nums[i]-1]可能位置还要发生变动
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 { //若nums[i]!=i+1,说明缺少该元素
			return i + 1
		}
	}
	return N + 1
}
func main() {
	nums := pkg.CreateIntSlice()
	fmt.Println(firstMissingPositive(nums))
}
