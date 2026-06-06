// 31. 下一个排列
/*
整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，
那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。
必须 原地 修改，只允许使用额外常数空间。

提示：
1 <= nums.length <= 100
0 <= nums[i] <= 100
*/
package main

import (
	"fmt"
	"lc/pkg"
	"sort"
)

func nextPermutation(nums []int) {
	n := len(nums)
	p1, p2 := n-1, n-1
	sourceIdx, targetIdx := -1, -1
	for p1 >= 0 {
		p2 = p1
		for p2 >= 0 && nums[p2] >= nums[p1] {
			p2--
		}
		if targetIdx < p2 {
			sourceIdx = p1
			targetIdx = p2
		}
		p1--
	}
	if targetIdx == -1 {
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		return
	}
	nums[sourceIdx], nums[targetIdx] = nums[targetIdx], nums[sourceIdx]
	sort.Slice(nums[targetIdx+1:], func(i, j int) bool {
		return nums[targetIdx+1+i] < nums[targetIdx+1+j]
	})
	return
}

func nextPermutation2(nums []int) {
	/*
		与上面的解法思路一样, 都是寻找右侧较大的数, 与左侧较小的数交换, 且较小数尽量靠右, 较大数尽可能小(仍位于较小数右侧),
		这样可以使得当前排列变大, 且变大幅度尽可能小
		找到这样一对数后, 交换, 并对较小数及其右侧数组进行升序排序, 从而得到下一个排列
		对于降序数组(如 [3, 2, 1]), 直接升序排序整个数组即可

		升序排序可以使用双指针, 因为较小数及其右侧实际上为降序数组, 直接逆序即可
	*/
	n := len(nums)
	sourceIdx, targetIdx := n-2, n-1
	for sourceIdx >= 0 && nums[sourceIdx] >= nums[sourceIdx+1] {
		sourceIdx--
	}
	if sourceIdx >= 0 {
		for targetIdx >= sourceIdx && nums[targetIdx] <= nums[sourceIdx] {
			targetIdx--
		}
		nums[sourceIdx], nums[targetIdx] = nums[targetIdx], nums[sourceIdx]
	}
	reverseIntSlice(nums[sourceIdx+1:])
	return
}

func reverseIntSlice(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 示例 1：
// 输入：nums = [1,2,3]
// 输出：[1,3,2]
//
// 示例 2：
// 输入：nums = [3,2,1]
// 输出：[1,2,3]
//
// 示例 3：
// 输入：nums = [1,1,5]
// 输出：[1,5,1]
func main() {
	nums := pkg.CreateSlice[int]()
	nums2 := make([]int, len(nums), cap(nums))
	copy(nums2, nums)

	nextPermutation(nums)
	fmt.Println(nums)

	nextPermutation2(nums2)
	fmt.Println(nums2)
}
