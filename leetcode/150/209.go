/*
209. 长度最小的子数组
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其总和大于等于 target 的长度最小的 连续
子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

示例 1：
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。

示例 2：
输入：target = 4, nums = [1,4,4]
输出：1

示例 3：
输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0

提示：

1 <= target <= 10^9
1 <= nums.length <= 10^5
1 <= nums[i] <= 10^5

进阶：
如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法

O(n log(n)) ---->前缀和+二分
*/
package main

import (
	"fmt"
	"math"
	"lc/100/pkg"
)

func minSubArrayLen(target int, nums []int) int {
	var ans int = math.MaxInt32
	var sum int
	sum += nums[0]
	var i, j int
	for j < len(nums) && sum < target {
		j++
		if j < len(nums) {
			sum += nums[j]
		}
	}
	for j < len(nums) {
		for i <= j && sum >= target {
			sum = sum - nums[i]
			i++
		}
		ans = min(ans, j-i+2)
		for j < len(nums) && sum < target {
			j++
			if j < len(nums) {
				sum += nums[j]
			}
		}
	}
	if ans == math.MaxInt32 {
		return 0
	} else {
		return ans
	}
}
func main() {
	fmt.Println("Input target:")
	var target int
	fmt.Scan(&target)
	nums := pkg.CreateIntSlice()
	fmt.Println(minSubArrayLen(target, nums))
}
