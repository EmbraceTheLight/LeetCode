/*
740. 删除并获得点数
给你一个整数数组 nums ，你可以对它进行一些操作。

每次操作中，选择任意一个 nums[i] ，删除它并获得 nums[i] 的点数。之后，你必须删除 所有 等于 nums[i] - 1 和 nums[i] + 1 的元素。
开始你拥有 0 个点数。返回你能通过这些操作获得的最大点数。

示例 1：

输入：nums = [3,4,2]
输出：6
解释：
删除 4 获得 4 个点数，因此 3 也被删除。
之后，删除 2 获得 2 个点数。总共获得 6 个点数。

示例 2：
输入：nums = [2,2,3,3,3,4]
输出：9
解释：
删除 3 获得 3 个点数，接着要删除两个 2 和 4 。
之后，再次删除 3 获得 3 个点数，再次删除 3 获得 3 个点数。
总共获得 9 个点数。

提示：
1 <= nums.length <= 2 * 10^4
1 <= nums[i] <= 10^4
*/
package main

import (
	"fmt"
	"sort"
)

func deleteAndEarn(nums []int) int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	mp := make(map[int]int) //mp记录nums中每个数字出现的次数
	n := len(nums)
	for i := 0; i < n; i++ {
		mp[nums[i]]++
	}

	if len(mp) == 1 {
		return nums[0] * mp[nums[0]]
	}

	dp := make([]int, len(mp)) //len(mp)为不重复的数字个数
	cur, pre := 0, 0           //记录nums当前数字、上一个不同数字出现的下标

	dp[0] = nums[0] * mp[nums[0]]
	cur = mp[nums[0]]
	if nums[cur]-nums[pre] > 1 {
		dp[1] = dp[0] + nums[cur]*mp[nums[cur]]
	} else {
		dp[1] = max(dp[0], nums[cur]*mp[nums[cur]])
	}
	pre = cur
	for i := 2; i < len(mp); i++ {
		cur = pre + mp[nums[pre]]
		if nums[cur]-nums[pre] > 1 {
			dp[i] = dp[i-1] + nums[cur]*mp[nums[cur]]
		} else {
			dp[i] = max(dp[i-1], dp[i-2]+nums[cur]*mp[nums[cur]])
		}
		pre = cur
	}
	return dp[len(mp)-1]
}

// 思路同 198. 打家劫舍。
// 具体思路：将sums数组看作打家劫舍中的一间间房子，具体的值就是房子中的现金值。
// 当偷取一间房子的时候，就不能再去它相邻的房间偷取
func deleteAndEarnRob(nums []int) int {
	n := len(nums)
	sums := make([]int, 10001) //sum存nums相同数字之和,切片长度要比数字范围大1
	for i := 0; i < n; i++ {
		sums[nums[i]] += nums[i]
	}
	return rob1(sums)
}
func rob1(sli []int) int {
	n := len(sli)
	var first, second int
	first = sli[0]
	second = max(sli[0], sli[1])
	ans := 0
	for i := 2; i < n; i++ {
		ans = max(first+sli[i], second)
		first = second
		second = ans
	}
	return ans
}
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println(deleteAndEarn(nums))
}
