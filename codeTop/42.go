// 42. 接雨水
/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

提示：
n == height.lengt
1 <= n <= 2 * 10^4
0 <= height[i] <= 10^5
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 单调栈
func trap(height []int) int {
	stk := make([]int, 0)
	var ans int
	idx := 0
	for ; idx < len(height)-1; idx++ {
		if height[idx] > height[idx+1] {
			stk = append(stk, idx)
			idx++
			break
		}
	}
	for ; idx < len(height); idx++ {
		if len(stk) == 0 {
			stk = append(stk, idx)
			continue
		}

		top := stk[len(stk)-1]
		if height[idx] <= height[top] {
			stk = append(stk, idx)
		} else {
			for len(stk) > 1 && height[top] < height[idx] {
				w := idx - stk[len(stk)-2] - 1
				h := min(height[stk[len(stk)-2]], height[idx]) - height[top]
				ans += w * h
				stk = stk[:len(stk)-1]
				top = stk[len(stk)-1]
			}
			if height[top] < height[idx] {
				stk = stk[:len(stk)-1]
			}
			stk = append(stk, idx)
		}
	}
	return ans
}

// 双指针
func trap2(height []int) int {
	var ans int
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0

	for left < right {
		// 当前左侧柱子高度小于右侧柱子高度
		if height[left] < height[right] {
			// 当前左侧柱子比已记录的左侧最大柱子要高, 更新左侧最高柱子高度
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				// 不用管右侧柱子. 因为此时 leftMax 小于 rightMax, leftMax 决定了能接多少水
				// 证明: 因为初始时 height[left] == leftMax, 因此, 当 left 要移动时, 一定是当前 height[right] 大于 height[left], 即 height[right] > leftMax
				// 因此, rightMax ＞ height[left] ＞= leftMax
				// 因此不用考虑此时的 rightMax
				ans += leftMax - height[left]
			}
			left++

			// 当前右侧柱子高度小于左侧柱子高度
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
	}
	return ans
}

// 示例 1：
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
// 示例 2：
// 输入：height = [4,2,0,3,2,5]
// 输出：9
//
// 示例 3：
// 输入：[5,5,1,7,1,1,5,2,7,6]
// 输出：23
func main() {
	height := pkg.CreateSlice[int]()
	fmt.Println(trap(height))
	fmt.Println(trap2(height))
}
