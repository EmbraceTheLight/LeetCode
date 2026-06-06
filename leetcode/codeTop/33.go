// 33. 搜索旋转排序数组
/*
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 向左旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 下标 3 上向左旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target，
如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

提示：
1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
nums 中的每个值都 独一无二
题目数据保证 nums 在预先未知的某个下标上进行了旋转
-10^4 <= target <= 10^4
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		// 始终用 mid ± 1：因为 nums[mid] == target 已提前返回
		if nums[l] <= nums[mid] { // 左半段有序
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1 // target 在左半段
			} else {
				l = mid + 1 // target 在右半段
			}
		} else { // 右半段有序
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1 // target 在右半段
			} else {
				r = mid - 1 // target 在左半段
			}
		}
	}
	return -1
}

// 示例 1：
// 输入：nums = [4,5,6,7,0,1,2], target = 0
// 输出：4
//
// 示例 2：
// 输入：nums = [4,5,6,7,0,1,2], target = 3
// 输出：-1
//
// 示例 3：
// 输入：nums = [1], target = 0
// 输出：-1
func main() {
	var target int
	fmt.Println("Input target:")
	fmt.Scan(&target)
	nums := pkg.CreateSlice[int]()
	fmt.Println(search(nums, target))
}
