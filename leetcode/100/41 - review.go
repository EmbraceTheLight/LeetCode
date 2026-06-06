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
-2^31 <= nums[i] <= 2631 - 1
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 将大小符合(0,len(nums))的数字进行原地排列，大小为n的数字排列在nums[n-1]处并与原位置数字交换，重复检查原位置数字是否符合条件
func firstMissingPositive41r(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(firstMissingPositive41r(nums))
}
