// 239. 滑动窗口最大值
/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
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
	"lc/pkg"
)

func maxSlidingWindow(nums []int, k int) []int {
	var ans []int

	// queue 是一个双端单调队列, 存储下标, nums[下标] 从大到小.
	// 当右指针扫描到的数比该队列最后一个下标的数还大时, 将该队列最后一个下标移除, 直到队列为空或者队尾下标对应的数比当前数大
	// 当左指针向右移动时, 指示窗口最大值的下标 queue[0] 可能位于窗口左侧, 不应继续参与计算, 也要移除
	queue := make([]int, 0)
	n := len(nums)
	left, right := 0, 0
	for right = 0; right < k; right++ {
		if len(queue) == 0 {
			queue = append(queue, right)
			continue
		}
		if nums[right] > nums[queue[len(queue)-1]] {
			for len(queue) > 0 && nums[right] >= nums[queue[len(queue)-1]] {
				queue = queue[:len(queue)-1]
			}
		}
		queue = append(queue, right)
	}
	right--

	for {
		// 左指针右移一位
		ans = append(ans, nums[queue[0]])
		left++
		// 窗口维护的最大元素位置不再在窗口中了, 移除
		if queue[0] < left {
			queue = queue[1:]
		}

		// 右指针右移一位, 完成一轮窗口右移
		right++
		if right == n {
			break
		}
		if len(queue) > 0 && nums[right] > nums[queue[len(queue)-1]] {
			for len(queue) > 0 && nums[right] >= nums[queue[len(queue)-1]] {
				queue = queue[:len(queue)-1]
			}
		}
		queue = append(queue, right)
	}
	return ans
}
func main() {
	var k int
	fmt.Println("Input k:")
	fmt.Scan(&k)
	nums := pkg.CreateSlice[int]()
	fmt.Println(maxSlidingWindow(nums, k))
}
