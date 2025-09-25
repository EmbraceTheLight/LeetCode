package main

import (
	"fmt"
	"lc/pkg"
)

func firstMissingPositiveR2(nums []int) int {
	n := len(nums)
	var tmp int
	for i := 0; i < n; {
		if nums[i] < n && nums[i] > 0 && nums[i] != i+1 {
			ind := nums[i] - 1
			tmp = nums[ind]
			nums[ind], nums[i] = nums[i], tmp
			if nums[ind] == nums[i] { // 防止两个位置的元素相等从而陷入死循环
				i++
				continue
			}
		} else {
			i++
			continue
		}

	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1

}

func main() {
	nums := pkg.CreateIntSlice[int]()
	fmt.Println(firstMissingPositiveR2(nums))

}
