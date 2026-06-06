package main

import (
	"fmt"
	"lc/pkg"
)

func maxSlidingWindowR(nums []int, k int) []int {
	var ans []int
	var q []int // 单调队列，存放元素下标。队头元素是当前滑动窗口最大值
	left, right := 0, 0

	for right = 0; right < k; right++ {
		for len(q) > 0 && nums[right] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1] // 在后面遇到了更大的元素，为维护单调队列，将队尾元素出队
		}
		q = append(q, right)
	}
	ans = append(ans, nums[q[0]])

	for left = 1; right < len(nums); left, right = left+1, right+1 {
		if q[0] < left { // 最大元素下标已经不在滑动窗口中，将其出队
			q = q[1:] // 将队头元素出队
		}

		for len(q) > 0 && nums[right] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1] // 在后面遇到了更大的元素，为维护单调队列，将队尾元素出队
		}
		q = append(q, right)

		ans = append(ans, nums[q[0]])
	}
	return ans
}
func main() {
	var k int
	fmt.Scan(&k)
	nums := pkg.CreateSlice[int]()
	fmt.Println(maxSlidingWindowR(nums, k))
}
