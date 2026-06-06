/*
215.数组中的第K个最大元素
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:
输入: [3,2,1,5,6,4], k = 2
输出: 5

示例 2:
输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4

提示：
1 <= k <= nums.length <= 10^5
-10^4 <= nums[i] <= 10
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

// PrioriQueue215 大根堆
type PrioriQueue215 struct {
	q   []int
	pos int //指示最后一个元素的下标
}

func (pq *PrioriQueue215) Len() int {
	return len(pq.q)
}
func (pq *PrioriQueue215) Insert(val int) {
	//if pq.Len() == 2 {
	//	pq.q[1] = val
	//	return
	//}
	pq.q = append(pq.q, val)
	pq.pos++
	if pq.pos == 1 { //插入的是第一个元素
		return
	} else {
		var i = pq.pos
		for i/2 >= 1 && pq.q[i] > pq.q[i/2] { //元素上浮操作
			pq.q[i], pq.q[i/2] = pq.q[i/2], pq.q[i]
			i = i / 2
		}
	}

}
func (pq *PrioriQueue215) Top() int {
	return pq.q[1]
}
func (pq *PrioriQueue215) Pop() []int {
	pq.q[1], pq.q[pq.pos] = pq.q[pq.pos], pq.q[1] //交换堆顶与堆尾元素
	pq.q = pq.q[:pq.pos]                          //移除堆尾元素
	pq.pos--
	var i = 1
	//对堆顶元素进行下沉操作
	for i*2+1 <= pq.pos {
		if pq.q[i] > pq.q[i*2+1] && pq.q[i] > pq.q[i*2] { //已经移动到了合适的位置，不用再下沉
			break
		} else {
			if pq.q[i*2] > pq.q[i*2+1] {
				pq.q[i], pq.q[i*2] = pq.q[i*2], pq.q[i]
				i = i * 2
			} else {
				pq.q[i], pq.q[i*2+1] = pq.q[i*2+1], pq.q[i]
				i = i*2 + 1
			}
		}
	}

	if i*2 <= pq.pos {
		if pq.q[i] < pq.q[i*2] {
			pq.q[i], pq.q[i*2] = pq.q[i*2], pq.q[i]
			i = i * 2
		}
	}
	return pq.q
}
func NewPrioriQueue() *PrioriQueue215 {
	return &PrioriQueue215{
		q:   make([]int, 1),
		pos: 0,
	}
}

func findKthLargest(nums []int, k int) int {
	pq := NewPrioriQueue()
	for _, v := range nums {
		pq.Insert(v)
	}

	for i := 0; i < k-1; i++ {
		pq.q = pq.Pop()
	}
	return pq.Top()
}

func BucketSort(nums []int, k int) int {
	var array [20001]int
	l := len(nums)
	for i := 0; i < l; i++ {
		array[nums[i]+10000]++ //加10000为了防止下标为负数
	}
	for i := 20000; i >= 0; i-- {
		k -= array[i]
		if k <= 0 {
			return i - 10000
		}
	}
	return 0
}
func main() {
	var K int
	fmt.Println("Input K:")
	fmt.Scan(&K)

	nums := pkg.CreateIntSlice()
	fmt.Println(findKthLargest(nums, K))
	fmt.Println(BucketSort(nums, K))

}
