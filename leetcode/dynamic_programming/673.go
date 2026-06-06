/*
673. 最长递增子序列的个数
给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。
注意 这个数列必须是 严格 递增的。

示例 1:

输入: [1,3,5,4,7]
输出: 2
解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。

示例 2:
输入: [2,2,2,2,2]
输出: 5
解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。

提示:

1 <= nums.length <= 2000
-10^6 <= nums[i] <= 10^6
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func findNumberOfLIS(nums []int) int {
	var ans = 0
	n := len(nums)
	dp := make([]int, len(nums))
	cnt := make([]int, n) //cnt[i]：以nums[i]结尾的最长上升子序列个数
	lengthMax := 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		cnt[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				//FIXME：难点
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1 //更新最大长度
					cnt[i] = cnt[j]   //由于以nums[i]结尾的最长上升子序列长度更新，因此其个数也要重置
				} else if dp[j]+1 == dp[i] { //长度恰好相等，则更新序列个数
					cnt[i] += cnt[j]
				}
				/*dp[j]+1长度小于dp[i]的情况，不符合cnt[i]的定义，没有必要关心。故不予讨论*/
			}
		}
		if dp[i] > lengthMax { //最大长度更新，故ans也要更新（覆盖更新）
			lengthMax = dp[i]
			ans = cnt[i]
		} else if dp[i] == lengthMax { //dp[i]最大长度刚好等于之前的最大长度，将ans累加
			ans += cnt[i]
		}
	}
	return ans
}
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(findNumberOfLIS(nums))
}
