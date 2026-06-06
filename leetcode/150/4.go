// 4. 寻找两个正序数组的中位数
/*
定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
请你找出并返回这两个正序数组的 中位数 。
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

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	n1, n2 := len(nums1), len(nums2)
//	total := n1 + n2
//	isEven := total%2 == 0
//	leftCount1, leftCount2 := 0, 0
//	l1, r1, l2, r2 := 0, n1-1, 0, n2-1
//	for l1 < r1 && l2 < r2 {
//		if nums1[l1] < nums2[l2] {
//			l1 = l1 + (r1-l1+1)/2
//			leftCount1 = l1
//		} else {
//			l2 = l2 + (r2-l2+1)/2
//			leftCount2 = l2
//		}
//		if !isEven && leftCount1+leftCount2 == total/2 {
//			if leftCount1 < leftCount2 {
//				return float64(nums1[l1])
//			} else {
//				return float64(nums2[l2])
//			}
//		} else if isEven && leftCount1+leftCount2 == (total-1)/2 {
//			return float64(nums1[l1]+nums2[l2]) / 2
//		}
//	}
//
//}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	total := n1 + n2
	midIdx := total / 2
	if total%2 == 0 {
		return (float64(findKthElement(nums1, nums2, midIdx)) + float64(findKthElement(nums1, nums2, midIdx+1))) / 2
	} else {
		return float64(findKthElement(nums1, nums2, midIdx+1))
	}
}

// 寻找 nums1 和 nums2 中第 k 小的元素
func findKthElement(nums1, nums2 []int, k int) int {
	idx1, idx2 := 0, 0
	for {
		if idx1 == len(nums1) {
			return nums2[idx2+k-1]
		}
		if idx2 == len(nums2) {
			return nums1[idx1+k-1]
		}
		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}
		mid := k / 2
		nIdx1 := min(idx1+mid, len(nums1)) - 1 // 最多选择 nums1 k/2 - 1 个数
		nIdx2 := min(idx2+mid, len(nums2)) - 1 // 最多选择 nums2 k/2 - 1 个数

		// nums1[nIdx1] <= nums2[nIdx2], 即 nums1[nIdx1] 最多比 nums1[idx1:nIdx1] 这 k/2 - 1 个数和 nums2[idx:nIdx2] 这 k/2 - 1 个数大,
		// 即最多比 这 k/2 - 1 +  k/2 - 1 = k - 2 个数大
		// 说明目前 nums1[idx1:nIdx1] 都不满足要求, 可以全部跳过
		if nums1[nIdx1] <= nums2[nIdx2] {
			k = k - (nIdx1 - idx1 + 1)
			idx1 = nIdx1 + 1
		} else {
			k = k - (nIdx2 - idx2 + 1)
			idx2 = nIdx2 + 1
		}
	}
}

/*
示例 1：
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3]，中位数 2

示例 2：
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4]，中位数 (2 + 3) / 2 = 2.5
*/
func main() {
	nums1 := pkg.CreateSlice[int]()
	nums2 := pkg.CreateSlice[int]()
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
