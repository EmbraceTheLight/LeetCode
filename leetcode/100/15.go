/*
15. 三数之和
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。
请你返回所有和为 0 且不重复的三元组。
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
解释：唯一可能的三元组和不为 0 。

示例 3：
输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。

提示：
3 <= nums.length <= 3000
-10^5 <= nums[i] <= 10^5
*/
package main

import (
	"fmt"
	"sort"
)

// func retMinus(a, b, c int) int { //返回c-b-a,其中，c>b>a
//
//		if a > b {
//			if b > c { //a>b>c
//				return a - b - c
//			} else { //c>b
//				if a < c { //c>a>b
//					return c - a - b
//				} else { //a>c>b
//					return a - c - b
//				}
//			}
//		} else { //b>a
//			if a < c {
//				if b < c { //c>b>a
//					return c - b - a
//				} else { //b>c>a
//					return b - c - a
//				}
//			} else { //b>a>c
//				return b - a - c
//			}
//		}
//	}
//
//	func threeSum(nums []int) [][]int {
//		result := make([][]int, 0)
//		mp := make(map[int]int)
//		check := make(map[int]bool) //检查重复元组
//		l := len(nums)
//		if l < 3 {
//			return result
//		}
//		for i, v := range nums {
//			mp[v] = i
//		}
//		for i := 0; i < l-2; i++ { //固定一个数不变，然后执行两数之和逻辑
//			target := 0 - nums[i]
//			for j := i + 1; j < l-1; j++ { //两数之和
//				tmp := make([]int, 0)
//				if _, ok := mp[target-nums[j]]; ok { //确保有需要的元素且i!=j!=k
//					if check[retMinus(nums[i], target-nums[j], nums[j])] == false {
//						tmp = append(tmp, nums[i], target-nums[j], nums[j])
//						result = append(result, tmp)
//						check[retMinus(nums[i], target-nums[j], nums[j])] = true
//					}
//				}
//			}
//		}
//		return result
//	}

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	l := len(nums)
	if l < 3 {
		return result
	}
	sort.Ints(nums)
	var first = 0
	var second = first + 1
	var third = l - 1

	var tf, ts, tt int //存放数组当前值

	for first < l-2 {
		tf = nums[first]
		ts = nums[second]
		tt = nums[third]
		if nums[first] > 0 { //第一个元素都大于0了，以后的数也都必大于0可以结束循环了
			break
		}

		//second、third双指针
		for second < third {
			if second < third && nums[first]+nums[second]+nums[third] == 0 { //找到相应元素了
				ts = nums[second]
				tt = nums[third]
				result = append(result, []int{nums[first], nums[second], nums[third]})
				for second < third && ts == nums[second] {
					second++
				}
				ts = nums[second]
				for second < third && tt == nums[third] {
					third--
				}
				tt = nums[third]
				continue
			}
			if nums[first]+nums[second]+nums[third] > 0 {
				for second < third && tt == nums[third] {
					third--
				}
				tt = nums[third]
			} else { //nums[first]+nums[second]+nums[third] < 0,需要将second右移
				second++
				//third = l - 1
			}

		}
		for first < l-2 && tf == nums[first] {
			first++
		}
		second = first + 1
		third = l - 1

	}
	return result
}
func main() {
	ints := make([]int, 0)
	fmt.Println("Input Value:")
	var tmp int

	for {
		fmt.Scan(&tmp)
		if tmp == -100 {
			break
		}
		ints = append(ints, tmp)
	}
	fmt.Printf("%+v\n", ints)
	fmt.Println(threeSum(ints))
}
