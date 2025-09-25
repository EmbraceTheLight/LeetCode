/*
1. 两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]
提示：

2 <= nums.length <= 104
-10^9 <= nums[i] <= 10^9
-10^9 <= target <= 10^9
只会存在一个有效答案
*/
package main

import "fmt"

//	func twoSum(nums []int, target int) []int {
//		var result []int
//		mp := make(map[int]int) //键--元素值；值--索引
//		for i, j := range nums {
//			mp[j] = i
//		}
//		for i, v := range nums {
//			if _, ok := mp[target-v]; ok == true {
//				if mp[target-v] != i { //索引不同
//					result = append(result, mp[target-v], i)
//					break
//				}
//			}
//		}
//		return result
//	}
func twoSum(nums []int, target int) []int {
	var result []int
	mp := make(map[int]int) //键--元素值；值--索引
	for i, j := range nums {
		mp[j] = i
	}
	for i, v := range mp {
		if _, ok := mp[target-v]; ok == true {
			if mp[target-v] != i { //索引不同
				result = append(result, mp[target-v], i)
				break
			}
		}
	}
	return result
}

func main() {
	ints := make([]int, 0)
	var target int
	fmt.Println("Input target:")
	fmt.Scan(&target)

	fmt.Println("Input Value:")
	var tmp int

	for {
		fmt.Scan(&tmp)
		if tmp == -1 {
			break
		}
		ints = append(ints, tmp)
	}
	fmt.Printf("%+v\n", ints)
	fmt.Println(twoSum(ints, target))
}
