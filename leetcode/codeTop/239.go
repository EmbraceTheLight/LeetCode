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

// 之前使用大根堆维护窗口最大值:
// 遇到比大根堆最大值大的元素, 就将大根堆清空, 并插入该元素; 窗口移动时, 如果移出窗口的元素恰好是大根堆最大值, 将大根堆最大值弹出
// 实际上该方法不成立
// 反例：[100, 5, 4, 6], k = 3
// 当前候选最大值：[100,5,4]
// 新元素：6
//
// 6 虽然小于 100, 无法淘汰 100,
// 但 6 比 5、4 更大, 且更晚过期,
// 因此 5、4 以后永远不可能成为窗口最大值，应该被淘汰。
// 如果按照大根堆解法, 在 100 出堆后, 6 不会入堆, 实际会取到 5, 然而本来应该取到数字 6 的
// 因此答案变为了 [100, 5]. 正确答案应为 [100, 6]
//
// 这说明：
// 新元素不一定只在“大于当前最大值”时才有意义
// 它还有可能淘汰一部分较小元素
//
// 因此简单的大根堆思路不成立：
// 堆只能快速获取最大值, 无法高效淘汰“小于等于新元素”的部分元素。
// 所以本题需要使用“单调递减队列”维护候选最大值。
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
