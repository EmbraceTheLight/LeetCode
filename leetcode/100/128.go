/*
128. 最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

示例 2：
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9


提示：
0 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9
*/

package main

import (
	"fmt"
)

//	func longestConsecutive(nums []int) int {
//		if len(nums) == 0 {
//			return 0
//		}
//		mp := make(map[int]int)
//		var result = 1
//		for _, i := range nums {
//			_, ok := mp[i-1]
//			if ok {
//				mp[i-1] = i
//			}
//			_, ok = mp[i+1]
//			if ok {
//				mp[i] = i + 1
//				continue
//			}
//			mp[i] = int(1e10)
//		}
//		visit := make(map[int]int)
//		sort.Ints(nums)
//		for _, i := range nums {
//			var cnt = 1
//			if visit[mp[i]] == 1 {
//				continue
//			}
//			for mp[i] < int(1e10) {
//				visit[i] = 1 //标记为已访问
//				i = mp[i]
//				cnt++
//			}
//			result = max(result, cnt)
//		}
//		return result
//	}
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 0
	mp := make(map[int]bool)
	for _, i := range nums {
		mp[i] = true
	}

	for i := range mp {
		var length int
		if mp[i-1] == false { //查看前面还有没有比自己小1的元素，若无，则自己是某个连续序列的第一个元素，可以看出前一个for的重要性：表示存在该元素
			length++
			c := i
			for mp[c+1] {
				c++
				length++
			}
		}

		if result < length {
			result = length
		}
	}
	return result
}
func main() {
	ints := make([]int, 0)
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
	a := longestConsecutive(ints)
	fmt.Println(a)
}
