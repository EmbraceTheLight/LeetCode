/*
31. 下一个排列
整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。
更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。
必须 原地 修改，只允许使用额外常数空间。

示例 1：
输入：nums = [1,2,3]
输出：[1,3,2]

示例 2：
输入：nums = [3,2,1]
输出：[1,2,3]

示例 3：
输入：nums = [1,1,5]
输出：[1,5,1]

提示：

1 <= nums.length <= 100
0 <= nums[i] <= 100
*/

package main

import (
	"fmt"
	"sort"
	"lc/100/pkg"
)

func nextPermutation(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	i := l - 1
	for ; i >= 1; i-- {
		if nums[i-1] < nums[i] {
			i--
			break
		}
	}
	if i == 0 && nums[0] >= nums[1] { //已经是最大序列,重排为字典序最小的排列（即，其元素按升序排列）
		for j := 0; j < l/2; j++ {
			nums[j], nums[l-j-1] = nums[l-j-1], nums[j]
		}
	} else { //否则下沉nums[i]到比它大的第一个元素之前
		exc := i               //exc记录要被交换的元素的位置
		var minest int = i + 1 //记录比exc大的最小元素下标
		for i < l {
			if nums[i] > nums[exc] {
				if nums[i] < nums[minest] {
					minest = i
				}
			}
			i++
		}

		nums[exc], nums[minest] = nums[minest], nums[exc]
		sort.Ints(nums[exc+1:])
		//for i < l-1 {
		//	nums[i], nums[i+1] = nums[i+1], nums[i]
		//	i++
		//}
	}
}

func main() {
	nums := pkg.CreateIntSlice()
	nextPermutation(nums)
	fmt.Println(nums)
}
