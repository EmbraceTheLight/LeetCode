// 215. 数组中的第K个最大元素
/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
提示：
1 <= k <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 一趟快速选择算法
func partitionLomuto(nums []int, left, right int) int {
	/*
		pivot 在 nums[right]
		该函数返回 pivot 的最终位置
	*/
	pivot := nums[right]
	i := left // i: 一趟快速选择后基准值 pivot 的位置下标
	for j := left; j < right; j++ {
		// 遇到 nums[j] 小于基准的情况, 交换 nums[i], nums[j], i, j 均 + 1
		// 这是因为 nums[i] 可能大于 nums[j]
		if nums[j] < pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func partitionHoare(nums []int, left, right int) int {
	/*
		// Hoare partition（霍尔分区）
		// 特点：
		// 1. pivot 通常选在左边（这里是 nums[left]）
		// 2. 使用双指针 i / j 从两侧向中间逼近
		// 3. 返回值 j 只是“分界点”，不是 pivot 的最终位置
		// 4. 分区结果：
		//    nums[left...j]   <= pivot
		//    nums[j+1...right] >= pivot
		// ⚠️ pivot 不保证在 j 上！
	*/
	pivot := nums[left]
	i := left - 1
	j := right + 1 // j: 一趟快速选择后基准值 pivot 的位置下标
	for i < j {
		for {
			i++
			if nums[i] >= pivot {
				break
			}
		}
		for {
			j--
			if nums[j] <= pivot {
				break
			}
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			break
		}
	}
	//for j := left; j < right; j++ {
	//	// 遇到 nums[j] 小于基准的情况, 交换 nums[i], nums[j], i, j 均 + 1
	//	// 这是因为 nums[i] 可能大于 nums[j]
	//	if nums[j] < pivot {
	//		nums[j], nums[i] = nums[i], nums[j]
	//		i++
	//	}
	//}

	return j
}
func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	n := len(nums)

	// Lomuto partition
	//idx := partitionLomuto(nums, left, right)
	//if idx == n-k {
	//	return nums[idx]
	//} else if idx < n-k { // 目标值在右边
	//	return quickSelect(nums, idx+1, right, k)
	//} else { // 目标值在左边
	//	return quickSelect(nums, left, idx-1, k)
	//}

	// Hoare partition
	idx := partitionHoare(nums, left, right)
	if idx < n-k { // 目标值在右边
		return quickSelect(nums, idx+1, right, k)
	} else { // 目标值在左边
		return quickSelect(nums, left, idx, k)
	}
}
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

// 示例 1:
// 输入: [3,2,1,5,6,4], k = 2
// 输出: 5
//
// 示例 2:
// 输入: [3,2,3,1,2,4,5,5,6], k = 4
// 输出: 4
func main() {
	var k int
	fmt.Println("Input k:")
	fmt.Scan(&k)
	nums := pkg.CreateSlice[int]()
	fmt.Println(findKthLargest(nums, k))
}
