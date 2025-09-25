/*
560.和为 K 的子数组
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数。
子数组是数组中元素的连续非空序列。

示例 1：
输入：nums = [1,1,1], k = 2
输出：2

示例 2：
输入：nums = [1,2,3], k = 3
输出：2

提示：
1 <= nums.length <= 2 * 10^4
-1000 <= nums[i] <= 1000
-10^7 <= k <= 10^7
*/
package main

import (
	"fmt"
	"lc/pkg"
)

//func subarraySum(nums []int, k int) int {
//	length := len(nums)
//	sumSli := make([]int, 0) //前缀和
//	var sum int
//	for i := 0; i < length; i++ {
//		sum += nums[i]
//		sumSli = append(sumSli, sum)
//	}
//
//	var ans int
//	// b - a = k ---> a = b - k ---> ans += mp[b - k]
//	var mp = make(map[int]int) //键：前缀和；值：该前缀和已经出现的次数
//	mp[0] = 1                  //初始化加入一个元素0，以处理子数组从下标0开始的特殊情况
//	for i := 0; i < length; i++ {
//		ans += mp[sumSli[i]-k]
//		mp[sumSli[i]]++
//	}
//	return ans
//}

/*改进：不用多一个循环专门计算前缀和*/
func subarraySum(nums []int, k int) int {
	length := len(nums)

	var ans, sum int
	// b - a = k ---> a = b - k ---> ans += mp[b - k]
	var mp = make(map[int]int) //键：前缀和；值：该前缀和已经出现的次数
	mp[0] = 1                  //初始化加入一个元素0，以处理子数组从下标0开始的特殊情况
	for i := 0; i < length; i++ {
		sum += nums[i]
		ans += mp[sum-k]
		mp[sum]++
	}
	return ans
}
func main() {
	fmt.Println("Input K:")
	var k int
	fmt.Scan(&k)
	nums := pkg.CreateIntSlice[int]()
	fmt.Println(subarraySum(nums, k))
}
