/*
42. 接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

提示：

n == height.length
1 <= n <= 2 * 10^4
0 <= height[i] <= 10^5
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func trap42r(height []int) int {
	stk := make([]int, 0)
	ans := 0
	for i, v := range height {
		for len(stk) > 0 && v > height[stk[len(stk)-1]] {
			top := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			for len(stk) > 0 && height[top] == height[stk[len(stk)-1]] {
				top = stk[len(stk)-1]
				stk = stk[:len(stk)-1]
			}

			if len(stk) > 0 {
				h := min(v, height[stk[len(stk)-1]]) - height[top]
				w := i - stk[len(stk)-1] - 1
				ans += h * w
				//fmt.Println("i:", i, "top:", top, "h:", h, "w:", w)
				//fmt.Println(ans)
			}
		}
		stk = append(stk, i)
	}

	return ans
}

// 双指针，该方法将柱子上能存的水量逐一相加
func trap42twoPointer(height []int) int {
	var left, right = 0, len(height) - 1
	var leftMax, rightMax int //存放扫描过的左右柱子最大值
	var ans int
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax < rightMax {
			ans = ans + (leftMax - height[left]) // 在下标为 left 的柱子上能存的水量
			left++
		} else {
			ans = ans + (rightMax - height[right]) // 在下标为 right 的柱子上能存的水量
			right--
		}
	}
	return ans
}
func main() {
	heights := pkg.CreateIntSlice[int]()
	fmt.Println(trap42r(heights))
	fmt.Println(trap42twoPointer(heights))
}
