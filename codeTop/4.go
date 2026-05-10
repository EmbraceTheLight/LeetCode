// 4. 寻找两个正序数组的中位数
/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
算法的时间复杂度应该为 O(log (m+n)) 。

提示：
nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-10^6 <= nums1[i], nums2[i] <= 10^6
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	sum := n1 + n2
	if sum%2 == 1 {
		return float64(findMinKthElement(nums1, nums2, sum/2+1))
	} else {
		return (float64(findMinKthElement(nums1, nums2, sum/2)) + float64(findMinKthElement(nums1, nums2, sum/2+1))) / 2
	}

}
func findMinKthElement(nums1, nums2 []int, k int) int {
	idx1, idx2 := 0, 0
	for {
		// nums1 已经遍历完, 从 nums2 中取第 k 小的数
		if idx1 == len(nums1) {
			return nums2[idx2+k-1]
		}
		if idx2 == len(nums2) {
			return nums1[idx1+k-1]
		}
		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}
		nIdx1 := min(idx1+k/2, len(nums1)) - 1
		nIdx2 := min(idx2+k/2, len(nums2)) - 1
		if nums1[nIdx1] <= nums2[nIdx2] {
			k = k - (nIdx1 - idx1 + 1)
			idx1 = nIdx1 + 1
		} else {
			k = k - (nIdx2 - idx2 + 1)
			idx2 = nIdx2 + 1
		}
	}
}

// 示例 1：
// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2
//
// 示例 2：
// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
func main() {
	nums1 := pkg.CreateSlice[int]()
	nums2 := pkg.CreateSlice[int]()
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
