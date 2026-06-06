package main

import (
	"fmt"
	"lc/pkg"
	"sort"
)

func lengthOfLISR(nums []int) int {
	increaseSli := make([]int, 0)
	increaseSli = append(increaseSli, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] > increaseSli[len(increaseSli)-1] {
			increaseSli = append(increaseSli, nums[i])
		} else {
			idx := sort.SearchInts(increaseSli, nums[i])
			increaseSli[idx] = nums[i]
		}
	}
	return len(increaseSli)
}

func main() {
	nums := pkg.CreateSlice[int]()
	fmt.Println(lengthOfLISR(nums))
}
