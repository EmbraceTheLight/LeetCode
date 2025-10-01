/*
300. 最长递增子序列
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

示例 1：
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4。

示例 2：
输入：nums = [0,1,0,3,2,3]
输出：4

示例 3：
输入：nums = [7,7,7,7,7,7,7]
输出：1

提示：
1 <= nums.length <= 2500
-10^4 <= nums[i] <= 10^4
*/
package main

import (
	"fmt"
	"lc/pkg"
	"sort"
)

func lengthOfLIS(nums []int) int {
	n := len(nums)
	tmp := make([]int, 0) //维护一个单调递增的数组,数组 tmp[i]表示长度为 i+1 的最长上升子序列的末尾元素的最小值
	tmp = append(tmp, nums[0])
	for i := 1; i < n; i++ {
		ind := sort.SearchInts(tmp, nums[i])
		if ind == len(tmp) {
			tmp = append(tmp, nums[i])
		} else {
			tmp[ind] = nums[i]
		}

	}
	return len(tmp)
}
func lengthOfLISDP(nums []int) int {
	var ans = 1
	n := len(nums)
	dp := make([]int, n) //dp[i]表示以nums[i]结尾的最长递增子序列,其中,nums[i]必须被选择
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := i; j >= 0; j-- {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1) // 最长递增子序列为 以 nums[j] 结尾再追加 nums[i]
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}
func main() {
	sli := pkg.CreateSlice[int]()
	fmt.Println(lengthOfLIS(sli))
}
