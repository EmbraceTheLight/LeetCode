/*
283. 移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]

示例 2:
输入: nums = [0]
输出: [0]

提示:
1 <= nums.length <= 10^4
-2^31 <= nums[i] <= 2^31 - 1
*/
package main

import "fmt"

func moveZeroes(nums []int) {
	l := len(nums)
	var pn0 = 0 //寻找第一个从0开始的非零位置
	var p0 = 0  //寻找距离数组左侧最近的0的位置
	for ; p0 < l; p0++ {
		if nums[p0] == 0 {
			break
		}
	}
	pn0 = p0
	for ; pn0 < l; pn0++ {
		if nums[pn0] != 0 {
			break
		}
	}

	for pn0 < l {
		if p0 == l-1 {
			break
		}
		nums[p0], nums[pn0] = nums[pn0], nums[p0]
		for ; pn0 != l; pn0++ {
			if nums[pn0] != 0 {
				break
			}
		}
		for ; p0 < l; p0++ {
			if nums[p0] == 0 {
				break
			}
		}
	}
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
	moveZeroes(ints)
	fmt.Println(ints)
}
