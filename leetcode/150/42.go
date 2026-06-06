// 42. 接雨水
/*
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

func trap(height []int) int {
	n := len(height)
	var ans int
	stk := make([]int, 0) // 单调栈，存储的是下标。如果存的是值，则无法获取两根柱子之间的距离
	for i := 0; i < n; i++ {
		// 当前栈为空，则直接入栈
		if len(stk) == 0 {
			stk = append(stk, i)
			continue
		}

		// 取栈顶元素
		curTop := stk[len(stk)-1]
		if height[i] <= height[curTop] {
			stk = append(stk, i)
		} else {
			for height[i] > height[curTop] {
				stk = stk[:len(stk)-1]
				if len(stk) == 0 {
					break
				}
				if len(stk) > 0 {
					left := stk[len(stk)-1]
					for height[left] == height[curTop] {
						stk = stk[:len(stk)-1]
						if len(stk) == 0 {
							break
						}
						left = stk[len(stk)-1]
					}
					h := min(height[left], height[i]) - height[curTop]
					w := i - left - 1
					ans += h * w
					if len(stk) == 0 {
						break
					}
					curTop = stk[len(stk)-1]
				}
			}

			// 若栈为空，入栈
			if len(stk) == 0 || height[i] <= height[curTop] {
				stk = append(stk, i)
			}
		}

	}
	return ans
}

// 单调栈：更简洁的做法: 不用处理繁琐的边界条件
func trapBetter(height []int) int {
	n := len(height)
	var ans int
	stk := make([]int, 0) // 单调栈，存储的是下标。如果存的是值，则无法获取两根柱子之间的距离
	for i := 0; i < n; i++ {
		// 当前栈不为空，且当前高度大于栈顶柱子高度，计算能接多少雨水
		for len(stk) > 0 && height[i] > height[stk[len(stk)-1]] {
			curTop := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if len(stk) > 0 {
				left := stk[len(stk)-1]
				h := min(height[left], height[i]) - height[curTop]
				w := i - left - 1
				ans += h * w
			}
		}
		stk = append(stk, i)
	}
	return ans
}

// Test Data: [0,1,0,2,1,0,1,3,2,1,2,1] expected output: 6
// Test Data: [4,2,0,3,2,5] expected output: 9
// Test Data: [5,2,1,2,1,5] expected output: 14
func main() {
	height := pkg.CreateSlice[int]()
	fmt.Println(trapBetter(height))

}
