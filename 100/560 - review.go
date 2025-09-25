package main

import "fmt"

func subarraySumR(nums []int, k int) int {
	//var left, right int
	//var ans, sum int
	//for left = 0; left < len(nums); left++ {
	//	sum = nums[left]
	//	if sum == k {
	//		ans++
	//	}
	//	for right = left + 1; right < len(nums); right++ {
	//		sum += nums[right]
	//		if sum == k {
	//			ans++
	//		}
	//	}
	//}

	//
	var ans, sum int

	var prefixSumMap = make(map[int]int)
	prefixSumMap[0] = 1
	sum = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		ans += prefixSumMap[sum-k]
		prefixSumMap[sum]++
	}
	return ans
}

func main() {
	var k int
	fmt.Println("Input k:")
	fmt.Scan(&k)

	fmt.Println("Input numbers:")
	var nums []int
	var tmp int
	fmt.Scan(&tmp)
	for tmp != -1 {
		nums = append(nums, tmp)
		fmt.Scan(&tmp)
	}
	fmt.Println(subarraySumR(nums, k))
}
