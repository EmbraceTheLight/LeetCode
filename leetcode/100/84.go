/*
84. 柱状图中最大的矩形

给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。

示例：
输入：heights = [2,1,5,6,2,3]
输出：10

输入： heights = [2,4]
输出： 4

1 <= heights.length <=10^5
0 <= heights[i] <= 10^4
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func largestRectangleArea(heights []int) int {
	// 思路：遍历每一根柱子
	// 以当前柱子为高，从该柱子开始向两侧延伸，直到遇到高度小于 该柱子高度 的柱子，就确定了矩形的左右边界，也就确定了矩形的宽度
	// 可以利用单调栈计算出每根柱子的左右边界，然后计算矩形面积。这样做，时间复杂度将优化为为 O(n)
	var ans = 0
	n := len(heights)
	stk := make([]int, 0)

	// ! left[i]表示第i根柱子左侧的第一个比它矮的柱子的下标（当前柱子的左边界）
	// ! right[i]表示第i根柱子右侧的第一个比它矮的柱子的下标（当前柱子的右边界）
	left, right := make([]int, n), make([]int, n)

	// 确定每根柱子的左边界
	for i := 0; i < n; i++ {
		for len(stk) > 0 && heights[i] <= heights[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			left[i] = -1
		} else {
			left[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}

	stk = []int{}

	// 确定每根柱子的右边界
	for i := n - 1; i >= 0; i-- {
		for len(stk) > 0 && heights[i] <= heights[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			right[i] = n
		} else {
			right[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}

	// 根据每根柱子的高度和左右边界计算矩形面积
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}
func main() {
	heights := pkg.CreateSlice[int]()
	fmt.Println(largestRectangleArea(heights))
}
