/*15. 三数之和*/
package main

import (
	"fmt"
	"sort"
)

func threeSum15R(nums []int) [][]int {
	sort.Ints(nums)
	var ans = make([][]int, 0)
	l := len(nums)
	for i := 0; i < l-2; i++ {
		if i >= 1 && nums[i-1] == nums[i] {
			continue
		}
		k := l - 1
		for j := i + 1; j < l-1 && j < k; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[i]+nums[j]+nums[k] > 0 {
				k--
			}
			if j == k {
				break
			}
			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[k-1] == nums[k] {
					k--
				}
			}
		}
	}
	return ans
}
func main() {
	//nums := pkg.CreateIntSlice()
	//fmt.Println(threeSum15R(nums))
	fmt.Println(threeSum15R([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum15R([]int{34, 55, 79, 28, 46, 33, 2, 48, 31, -3, 84, 71, 52, -3, 93, 15, 21, -43, 57, -6, 86, 56, 94, 74, 83, -14, 28, -66, 46, -49, 62, -11, 43, 65, 77, 12, 47, 61, 26, 1, 13, 29, 55, -82, 76, 26, 15, -29, 36, -29, 10, -70, 69, 17, 49}))
}
