// 15. 三数之和
/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。

示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。

示例 2：
输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0。

示例 3：
输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0。


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
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		// 去重
		for i > 0 && i < n-2 && nums[i] == nums[i-1] {
			i++
		}
		// 剪枝优化：
		// nums[i] + nums[i+1] + nums[i+2] > 0，
		// 说明 nums[i] 与 其后面最小的两个数nums[i+1]和nums[i+2]之和都大于 0，
		// 所以 i 后面的数都不可能组成三元组，直接终止循环
		if i < n-2 && nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}

		// nums[i] + nums[n-1] + nums[n-2] < 0，
		// 说明 nums[i] 与 其后面最大的两个数nums[n-1]和nums[n-2]之和都小于 0，
		// 说明 nums[i] 太小了，所以 i 前面的数都不可能组成三元组，继续增加 i 的值
		if i < n-2 && nums[i]+nums[n-1]+nums[n-2] < 0 {
			continue
		}

		// j 左指针，k 右指针
		j, k := i+1, n-1

		for ; j < n-1; j++ {
			// 去重
			for j > i+1 && j < n-1 && nums[j] == nums[j-1] {
				j++
			}

			// 给定 i, j 后，当 k 对应的值与 i, j 对应的值相加后大于 0 时，说明 i, j, k 三者之和大于 0，此时 k 应该减小，即右指针左移
			for j < k && nums[i]+nums[j]+nums[k] > 0 {
				k--
			}

			// k <= j，此时再增加 j 的值，就没有什么意义了
			if j >= k {
				break
			}

			if j < k && nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for k = k - 1; k > j && nums[k] == nums[k+1]; k-- {
				}
			}
		}

	}

	return ans
}

// Test Case 1: [-1,0,1,2,-1,-4]	Output: [[-1,-1,2],[-1,0,1]]
// Test Case 2: [0,1,1]	Output: []
// Test Case 3: [0,0,0]	Output: [[0,0,0]]
// Test Case 4: [0,0,0,0]	Output: [[0,0,0]]
// Test Case 5: [-100,-70,-60,110,120,130,160]	Output: [[-100,-60,160],[-70,-60,130]]-100,130,140],[-70,110,150],[-70,120,140],[-60,110,140],[-60,120,130]]
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(threeSum(nums))
}
