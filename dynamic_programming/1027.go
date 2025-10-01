/*
1027. 最长等差数列
给你一个整数数组 nums，返回 nums 中最长等差子序列的长度。

回想一下，nums 的子序列是一个列表 nums[i1], nums[i2], ..., nums[ik]，
且 0 <= i1 < i2 < ... < ik <= nums.length - 1。
并且如果 seq[i+1] - seq[i](0 <= i < seq.length - 1) 的值都相同，那么序列 seq 是等差的。

示例 1：
输入：nums = [3,6,9,12]
输出：4
解释：
整个数组是公差为 3 的等差数列。

示例 2：
输入：nums = [9,4,7,2,10]
输出：3
解释：
最长的等差子序列是 [4,7,10]。

示例 3：
输入：nums = [20,1,15,3,10,5,8]
输出：4
解释：
最长的等差子序列是 [20,15,10,5]。

提示：

2 <= nums.length <= 1000
0 <= nums[i] <= 500
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func handler(nums []int, diff int) int {
	dp := make(map[int]int)
	ans := 0
	for i := 0; i < len(nums); i++ {
		dp[nums[i]] = dp[nums[i]-diff] + 1
		if dp[nums[i]] > ans {
			ans = dp[nums[i]]
		}
	}
	return ans
}
func longestArithSeqLength(nums []int) int {
	mp := make(map[int]bool)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			mp[nums[j]-nums[i]] = true
		}
	}
	ans := 0
	for k, _ := range mp {
		ans = max(ans, handler(nums, k))
	}
	return ans
}

// FIXME:更快的解法
func longestArithSeqLengthDP(nums []int) int {
	numsMin, numsMax := nums[0], nums[0] //确定最小、最大值，用于确认等差差值的范围
	for _, v := range nums {
		if v > numsMax {
			numsMax = v
		} else if v < numsMin {
			numsMin = v
		}
	}
	diffMax := numsMax - numsMin //差值范围在-diffMax到diffMax之间
	//遍历差值
	ans := 0
	for diff := -diffMax; diff <= diffMax; diff++ {
		/*  复杂度O(N^2)
		dp := make(map[int]int)
		for _, v := range nums {
			dp[v]=dp[v-diff]+1
			ans=max(dp[v],ans)
		}
		*/
		//复杂度O(N*R),R为nums数据范围
		dp := make([]int, numsMax+1) //dp[i]表示以i结尾的最长等差数列长度.因为nums元素最小为0，所以dp大小取numsMax+1正好可以覆盖所有值
		for i := 0; i < numsMax; i++ {
			dp[i] = -1
		}
		for _, v := range nums {
			prev := v - diff
			if prev >= numsMin && prev <= numsMax && dp[prev] != -1 {
				dp[v] = max(dp[v], dp[prev]+1)
				ans = max(ans, dp[v])
			}
			dp[v] = max(dp[v], 1)
		}
	}
	return ans
}
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(longestArithSeqLength(nums))
}
