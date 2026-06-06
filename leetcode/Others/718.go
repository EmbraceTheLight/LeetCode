/*
718. 最长重复子数组
给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。

示例 1：
输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
输出：3
解释：长度最长的公共子数组是 [3,2,1] 。

示例 2：
输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
输出：5

提示：
1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 100
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, 0) //dp[i][j]：以nums1[i]和nums2[j]结尾的最长公共子数组的长度
	var ans int
	for i := 0; i < len(nums1)+1; i++ {
		var tmp = make([]int, len(nums2)+1)
		dp = append(dp, tmp)
	}
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			ans = max(ans, dp[i][j])

		}
	}
	return ans
}
func main() {
	nums1 := pkg.CreateIntSlice()
	nums2 := pkg.CreateIntSlice()
	fmt.Println(findLength(nums1, nums2))
}
