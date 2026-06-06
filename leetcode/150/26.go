//26. 删除有序数组中的重复项
/*
给你一个非严格递增排列的数组nums ，请你原地删除重复出现的元素，使每个元素只出现一次 ，返回删除后数组的新长度。元素的相对顺序应该保持一致。然后返回 nums 中唯一元素的个数。
考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。

示例 1：
输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。

示例 2：
输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
*/
package main

import "fmt"

//func removeDuplicates(nums []int) int {
//	l := len(nums)
//	var p1 = 0
//	var p2 = p1 + 1
//	var idx = 1 //元素待插入位置
//	var cnt = 0
//
//	for p2 < l {
//		for p2 < l && nums[p1] == nums[p2] {
//			p2++
//			cnt++ //记录重复元素数
//		}
//		if p2 != p1+1 { //有重复元素
//			if p2 < l { //后面还有新元素nums[p2]
//				nums[idx] = nums[p2]
//				idx++
//				p1 = p2
//				p2 = p1 + 1
//				continue
//			} else {
//				return l - cnt
//			}
//		} else {
//			p1++
//			p2 = p1 + 1
//			nums[idx] = nums[p1]
//			idx++
//		}
//	}
//	//fmt.Println(nums)
//	return l - cnt
//}

func removeDuplicates_26(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	var slow = 1
	var fast = 1
	for ; fast < l; fast++ {
		if nums[fast] == nums[fast-1] {
			continue
		} else { //nums[fast]为新元素
			nums[slow] = nums[fast]
			slow++
		}
	}
	fmt.Println(nums)
	return slow
}
func main() {
	var tmp int
	nums := make([]int, 0)
	for {
		fmt.Scan(&tmp)
		if tmp == -1 {
			break
		}
		nums = append(nums, tmp)
	}
	l := removeDuplicates_26(nums)
	fmt.Println("l:", l)
}
