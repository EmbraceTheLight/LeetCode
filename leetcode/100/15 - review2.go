package main

import (
	"fmt"
	"sort"
)

func threeSumRR(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var ans [][]int
	var right int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] && i < len(nums)-2 {
			continue
		}

		right = len(nums) - 1
		for left := i + 1; left < len(nums)-1; left++ {
			if left > i+1 && nums[left-1] == nums[left] {
				continue
			}

			for left < right && nums[i]+nums[left]+nums[right] > 0 {
				right--
			}
			if left == right {
				break
			}

			if nums[i]+nums[left]+nums[right] == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})

				// 接下来nums[left]一定比现在的nums[left]大，因此需要nums[right]比现在的小，这个循环就用于对right去重
				for right < len(nums)-1 && nums[right] == nums[right+1] {
					right--
				}
			}

		}
	}
	return ans
}

func main() {
	var nums []int
	var tmp int
	fmt.Scan(&tmp)
	for tmp != -100 {
		nums = append(nums, tmp)
		fmt.Scan(&tmp)
	}
	fmt.Println(threeSumRR(nums))
}
