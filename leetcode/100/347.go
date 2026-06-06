// 347. 前 K 个高频元素
/*
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

提示：
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
k 的取值范围是 [1, 数组中不相同的元素的个数]
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的

进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
*/
package main

import (
	"container/heap"
	"fmt"
	"lc/pkg"
)

type struct347 struct {
	num   int
	count int
}
type bigRoot []struct347

func (br *bigRoot) Len() int {
	return len(*br)
}
func (br *bigRoot) Swap(i, j int) {
	(*br)[i], (*br)[j] = (*br)[j], (*br)[i]
}
func (br *bigRoot) Less(i, j int) bool {
	return (*br)[i].count > (*br)[j].count
}
func (br *bigRoot) Push(x any) {
	*br = append(*br, x.(struct347))
}
func (br *bigRoot) Pop() any {
	ret := (*br)[len(*br)-1]
	*br = (*br)[:len(*br)-1]
	return ret
}

func topKFrequent(nums []int, k int) []int {
	br := bigRoot{}
	heap.Init(&br)
	mp1 := make(map[int]*struct347)
	for i := 0; i < len(nums); i++ {
		if _, ok := mp1[nums[i]]; !ok {
			mp1[nums[i]] = &struct347{nums[i], 1}
		} else {
			mp1[nums[i]].count++
		}
	}

	for _, v := range mp1 {
		heap.Push(&br, *v)
	}
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(&br).(struct347).num)
	}
	return ans
}

// 示例 1：
// 输入：nums = [1,1,1,2,2,3], k = 2
// 输出：[1,2]
//
// 示例 2：
// 输入：nums = [1], k = 1
// 输出：[1]
//
// 示例 3：
// 输入：nums = [1,2,1,2,1,2,3,1,3,2], k = 2
// 输出：[1,2]
func main() {
	var k int
	fmt.Println("Input k:")
	fmt.Scan(&k)
	nums := pkg.CreateSlice[int]()
	fmt.Println(topKFrequent(nums, k))
}
