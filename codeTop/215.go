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
	"math/rand"
)

// 一趟快速选择算法
func partition(nums []int, left, right int) int {
	pivotIdx := rand.Intn(right-left+1) + left
	pivot := nums[pivotIdx]
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
func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	n := len(nums)
	idx := partition(nums, left, right)
	if idx == n-k {
		return nums[idx]
	} else if idx < n-k { // 目标值在右边
		return quickSelect(nums, idx+1, right, k)
	} else { // 目标值在左边
		return quickSelect(nums, left, idx-1, k)
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
