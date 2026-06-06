// 15. 三数之和
/*
给你一个整数数组 nums，
判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

提示：
3 <= nums.length <= 3000
-10^5 <= nums[i] <= 10^5
*/
package main

import (
	"fmt"
	"lc/pkg"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var ans [][]int
	var first, second, third int
	for i := 0; i < n-2; i++ {
		first = nums[i]
		j, k := i+1, n-1
		for j < k {
			second, third = nums[j], nums[k]
			if first+second+third == 0 {
				ans = append(ans, []int{first, second, third})
				for j < k && nums[j] == second {
					j++
				}
				for j < k && nums[k] == third {
					k--
				}
			} else if first+second+third < 0 {
				j++
			} else {
				k--
			}
		}
		for i < n && nums[i] == first {
			i++
		}
		i--
	}
	return ans
}

// 示例 1：
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 解释：
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0。
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0。
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0。
// 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
// 注意，输出的顺序和三元组的顺序并不重要。
//
// 示例 2：
// 输入：nums = [0,1,1]
// 输出：[]
// 解释：唯一可能的三元组和不为 0。
//
// 示例 3：
// 输入：nums = [0,0,0]
// 输出：[[0,0,0]]
// 解释：唯一可能的三元组和为 0。
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(threeSum(nums))
}
