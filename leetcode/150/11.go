// 11. 盛最多水的容器

/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i])。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器。

示例 1：
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例 2：
输入：height = [1,1]
输出：1

提示：
n == height.length
2 <= n <= 10^5
0 <= height[i] <= 10^4
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func maxArea(height []int) int {
	var ans int
	n := len(height)
	left, right := 0, n-1
	maxLeft, maxRight := height[left], height[right]
	ans = (right - left) * min(maxLeft, maxRight)
	for left < right {
		if maxLeft < maxRight {
			for left < right && height[left] <= maxLeft {
				left++
			}
			maxLeft = height[left]
		} else {
			for left < right && height[right] <= maxRight {
				right--
			}
			maxRight = height[right]
		}
		ans = max(ans, (right-left)*min(maxLeft, maxRight))
	}

	return ans
}

// Test Case 1: [1,8,6,2,5,4,8,3,7]			Output: 49
// Test Case 2: [1,1]						Output: 1
func main() {
	height := pkg.CreateSlice[int]()
	fmt.Println(maxArea(height))
}
