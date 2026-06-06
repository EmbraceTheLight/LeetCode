// 219. 存在重复元素 II
/*
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。
如果存在，返回 true ；否则，返回 false 。

示例 1：
输入：nums = [1,2,3,1], k = 3
输出：true

示例 2：
输入：nums = [1,0,1,1], k = 1
输出：true

示例 3：
输入：nums = [1,2,3,1,2,3], k = 2
输出：false


提示：

1 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
0 <= k <= 10^5
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 本题也可用 滑动窗口解决
func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)
	mp := make(map[int]int) //key: value  value: index
	for i := 0; i < n; i++ {
		if _, ok := mp[nums[i]]; ok {
			if i-mp[nums[i]] <= k {
				return true
			}
		}
		mp[nums[i]] = i
	}
	return false
}

// Test Case1: 3, [1,2,3,1]		Output: true
// Test Case2: 1, [1,0,1,1]		Output: true
// Test Case3: 2, [1,2,3,1,2,3]	Output: false
func main() {
	var k int
	fmt.Println("Input k:")
	fmt.Scan(&k)
	nums := pkg.CreateSlice[int]()
	fmt.Println(containsNearbyDuplicate(nums, k))
}
