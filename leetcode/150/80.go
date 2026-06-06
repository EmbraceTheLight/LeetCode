/*
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成
示例 1：
输入：nums = [1,1,1,2,2,3]
输出：5, nums = [1,1,2,2,3]
解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3。 不需要考虑数组中超出新长度后面的元素。

示例 2：
输入：nums = [0,0,1,1,1,1,2,3,3]
输出：7, nums = [0,0,1,1,2,3,3]
解释：函数应返回新长度 length = 7, 并且原数组的前七个元素被修改为 0, 0, 1, 1, 2, 3, 3。不需要考虑数组中超出新长度后面的元素。

提示：
1 <= nums.length <= 3 * 10^4
-10^4 <= nums[i] <= 10^4
nums 已按升序排列
*/
package main

import "fmt"

//	func removeDuplicates_80(nums []int) int {
//		var insertIdx = 0
//		l := len(nums)
//		if l <= 2 {
//			return l
//		}
//		var p1 int
//		var p2 = p1 + 1
//		for p1 < l {
//			for p2 < l && nums[p2] == nums[p1] {
//				p2++
//			}
//
//			if p2 >= p1+2 { //p2==p1+2说明正好有两个相同的元素,即nums[p1]和nums[p1+1]
//				nums[insertIdx] = nums[p1]
//				insertIdx++
//				nums[insertIdx] = nums[p1]
//				insertIdx++
//
//			} else {
//				nums[insertIdx] = nums[p1]
//				insertIdx++
//			}
//
//			p1 = p2
//			p2 = p2 + 1
//		}
//		return insertIdx
//	}

// 双指针，
func removeDuplicates_Best(nums []int) int {
	slow, fast := 2, 2 //双指针从2开始，因为前两个元素一定包含在新数组中
	l := len(nums)
	if l < 2 {
		return l
	}

	for fast < l {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
func main() {
	nums := make([]int, 0)
	fmt.Println("Input Value:")

	var tmp int
	for {
		fmt.Scan(&tmp)
		if tmp == -1 {
			break
		}
		nums = append(nums, tmp)
	}
	fmt.Scan(&nums)
	a := removeDuplicates_Best(nums)
	//a := removeDuplicates_80(nums)
	fmt.Println("return Value:", a)
	for i := 0; i < a; i++ {
		fmt.Printf("%d ", nums[i])

	}
}
