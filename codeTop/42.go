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

// 示例 1：
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
// 示例 2：
// 输入：height = [4,2,0,3,2,5]
// 输出：9
func main() {
	height := pkg.CreateSlice[int]()
	fmt.Println(trap(height))
}
