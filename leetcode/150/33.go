/*
33.搜索旋转排序数组
整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

示例 2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

示例 3：
输入：nums = [1], target = 0
输出：-1
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

//func BinarySearch(nums []int, l, r, target int) int {
//	if l > r {
//		return -1
//	}
//	mid := (l + r) / 2
//	//BinarySearch(nums, 0, len(nums), target)
//	if nums[mid] == target {
//		return mid
//	}
//	lans := BinarySearch(nums, l, mid-1, target)
//	if lans != -1 {
//		return lans
//	}
//	rans := BinarySearch(nums, mid+1, r, target)
//	if rans != -1 {
//		return rans
//	}
//
//	return -1
//}
//
//func search(nums []int, target int) int {
//	var ans = -1
//	length := len(nums)
//	var l, r int
//	l, r = 0, length-1
//	//mid = (l + r) / 2
//	ans = BinarySearch(nums, l, r, target)
//	//for l <= r {
//	//	if nums[mid] == target {
//	//		ans = mid
//	//		break
//	//	} else if nums[mid] < target {
//	//		l = mid + 1
//	//		mid = (l + r) / 2
//	//	} else {
//	//		r = mid - 1
//	//		mid = (l + r) / 2
//	//	}
//	//}
//	return ans
//}

// 翻转后的数组总是一个有序，一个有序或部分有序的
func search(nums []int, target int) int {
	var ans = -1
	var l, r = 0, len(nums)
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			ans = mid
			break
		}
		if nums[l] <= nums[mid] { //l--mid为有序部分
			if nums[l] <= target && target < nums[mid] { //target落在了有序部分中
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { //mid--r为有序部分
			if target > nums[mid] && target <= nums[r] { //落在有序区间
				l = mid + 1
			} else {
				r = mid - 1
			}

		}
	}
	return ans
}
func main() {
	var target int
	fmt.Println("Input Target:")
	fmt.Scanf("%d\n", &target)
	nums := pkg.CreateIntSlice()
	fmt.Println(search(nums, target))

}
