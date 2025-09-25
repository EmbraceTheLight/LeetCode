/*
239. 滑动窗口最大值
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。

示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3

	1 [3  -1  -3] 5  3  6  7       3
	1  3 [-1  -3  5] 3  6  7       5
	1  3  -1 [-3  5  3] 6  7       5
	1  3  -1  -3 [5  3  6] 7       6
	1  3  -1  -3  5 [3  6  7]      7

示例 2：
输入：nums = [1], k = 1
输出：[1]

提示：
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

// 使用单调队列
func maxSlidingWindow(nums []int, k int) []int {
	var q []int //q存储nums下标。这些下标从小到大，对应的值严格单调递减
	var ans []int
	push := func(i int) {
		for len(q) > 0 && nums[i] > nums[q[len(q)-1]] { //若nums[i]比nums[len(q)-1]还要大，说明nums[len(q)-1]不可能是窗口最大值了，应当予以淘汰
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	//初始化单调队列
	for i := 0; i < k; i++ {
		push(i)
	}
	ans = append(ans, nums[q[0]])
	// 进行滑动
	l := len(nums)
	for i := k; i < l; i++ {
		push(i)
		for q[0] <= i-k { //当下标<=i-k时，说明其对应的数组元素在窗口左端点左侧，即已经不在窗口中了，应当予以删除
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}
func main() {
	fmt.Println("Input k:")
	var k int
	fmt.Scan(&k)
	nums := pkg.CreateIntSlice()
	fmt.Println(maxSlidingWindow(nums, k))
}
