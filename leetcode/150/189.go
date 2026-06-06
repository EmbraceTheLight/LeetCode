/*
189. 轮转数组
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

示例 2:
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]

提示：

1 <= nums.length <= 10^5
-2^31 <= nums[i] <= 2^31 - 1
0 <= k <= 10^5

进阶：
尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 第一种简单写法，开一个新数组存仿轮转之后的数组数据信息
// ! 另外的原地操作比较复杂，需要数学证明。这里就不实现了。
func rotate189(nums []int, k int) {
	newSli := make([]int, len(nums))

	n := len(nums)
	for i := 0; i < n; i++ {
		targetIdx := (i + k) % n
		newSli[targetIdx] = nums[i]
	}

	for i, v := range newSli {
		nums[i] = v
	}
}

func main() {
	var k int
	fmt.Scan(&k)
	nums := pkg.CreateSlice[int]()
	rotate189(nums, k)
	fmt.Println(nums)
}
