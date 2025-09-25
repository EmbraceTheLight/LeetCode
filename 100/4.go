/*
4.寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

示例 2：
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
*/
package main

import (
	"fmt"
	"lc/pkg"
)

//	func merge4(nums1 []int, nums2 []int) []int {
//		var ret = make([]int, len(nums1)+len(nums2))
//		l1 := len(nums1)
//		l2 := len(nums2)
//		var i, j, idx = 0, 0, 0
//		for i < l1 && j < l2 {
//			if nums1[i] < nums2[j] {
//				ret[idx] = nums1[i]
//				i++
//			} else {
//				ret[idx] = nums2[j]
//				j++
//			}
//
//			idx++
//		}
//		for i < l1 {
//			ret[idx] = nums1[i]
//			i++
//			idx++
//		}
//		for j < l2 {
//			ret[idx] = nums2[j]
//			j++
//			idx++
//		}
//		return ret
//	}
//
//	func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//		var sli = merge4(nums1, nums2)
//		var ans float64
//		if len(sli)%2 == 1 {
//			ans = float64(sli[len(sli)/2])
//		} else {
//			ans = (float64(sli[len(sli)/2-1]) + float64(sli[len(sli)/2])) / 2
//		}
//		return ans
//	}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 { //总长度为奇数的情况，取最中间的数即可
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else { //总长为偶数，取中间两数的平均值
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

// getKthElement 取第K大的元素，即下标K-1
func getKthElement(nums1, nums2 []int, k int) int {
	idx1, idx2 := 0, 0
	for {
		if idx1 == len(nums1) {
			return nums2[idx1+k-1]
		}
		if idx2 == len(nums2) {
			return nums1[idx2+k-1]
		}
		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}

		mid := k / 2 //将k二分，可以一次丢弃nums1或nums2的k/2个元素

		nIdx1 := min(idx1+mid, len(nums1)) - 1 //注意-1，否则就是一次性找k/2+1个元素
		nIdx2 := min(idx2+mid, len(nums2)) - 1
		//nums[nIdx1]>nums2[nIdx2],说明nums2的[0,nIdx2]部分可以全部丢弃，
		//因为即便是nums2[nIdx2]最多也只比(它自己的)k/2-1+(nums1的[idx,idx+k/2-1])k/2-1=k-2个数大，
		//nums2[nIdx2]自身最多就是第k-1大的元素。不符合要求，可以不用再找这一部分了
		//下同
		if nums1[nIdx1] > nums2[nIdx2] {
			k = k - (nIdx2 - idx2 + 1)
			idx2 = nIdx2 + 1 //将前idx个元素丢弃，开始新一轮遍历
		} else {
			k = k - (nIdx1 - idx1 + 1)
			idx1 = nIdx1 + 1
		}
	}
}
func main() {
	nums1 := pkg.CreateIntSlice[int]()
	nums2 := pkg.CreateIntSlice[int]()
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
