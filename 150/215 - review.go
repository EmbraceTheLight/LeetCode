// 215. 数组中的第K个最大元素
/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

提示：
1 <= k <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
*/

package main

import (
	"container/heap"
	"fmt"
	"lc/pkg"
)

type ints []int

func (q *ints) Len() int {
	return len(*q)
}
func (q *ints) Less(i, j int) bool {
	return (*q)[i] > (*q)[j]
}
func (q *ints) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}
func (q *ints) Push(x any) {
	item := x.(int)
	*q = append(*q, item)
}
func (q *ints) Pop() any {
	n := len(*q)
	item := (*q)[n-1]
	*q = (*q)[0 : n-1]
	return item
}
func findKthLargestR(nums []int, k int) int {
	var q ints
	heap.Init(&q)
	for _, num := range nums {
		heap.Push(&q, num)
	}

	for i := 1; i < k; i++ {
		heap.Pop(&q)
	}
	return heap.Pop(&q).(int)
}

// 示例 1:
// 输入: [3,2,1,5,6,4], k = 2
// 输出: 5
//
// 示例 2:
// 输入: [3,2,3,1,2,4,5,5,6], k = 4
// 输出: 4
func main() {
	var k int
	fmt.Println("Input k:")
	fmt.Scan(&k)
	fmt.Println("Input nums:")
	nums := pkg.CreateSlice[int]()
	fmt.Println(findKthLargestR(nums, k))
}
