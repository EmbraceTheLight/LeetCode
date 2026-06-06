// 128. 最长连续序列
/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

示例 2：
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9

示例 3：
输入：nums = [1,0,1,2]
输出：3


提示：
0 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func longestConsecutive(nums []int) int {
	n := len(nums)
	ans := 0
	mp := make(map[int]bool)
	visited := make(map[int]bool)
	for i := 0; i < n; i++ {
		mp[nums[i]] = true
	}

	//for i := 0; i < n; i++ {
	//	if visited[nums[i]] {
	//		continue
	//	}
	//	count := 1
	//	visited[nums[i]] = true
	//	bigger := nums[i] + 1
	//	smaller := nums[i] - 1
	//	for mp[bigger] {
	//		visited[bigger] = true
	//		bigger++
	//		count++
	//	}
	//	for mp[smaller] {
	//		visited[smaller] = true
	//		smaller--
	//		count++
	//	}
	//	ans = max(count, ans)
	//}

	for num := range mp {
		if visited[num] {
			continue
		}
		count := 1
		visited[num] = true
		bigger := num + 1
		smaller := num - 1
		for mp[bigger] {
			visited[bigger] = true
			bigger++
			count++
		}
		for mp[smaller] {
			visited[smaller] = true
			smaller--
			count++
		}
		ans = max(count, ans)
	}
	return ans
}

// Test Case1: [100,4,200,1,3,2]	Output: 4
// Test Case2: [0,3,7,2,5,8,4,6,0,1]	Output: 9
// Test Case3: [1,0,1,2]	Output: 3
func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(longestConsecutive(nums))
}
