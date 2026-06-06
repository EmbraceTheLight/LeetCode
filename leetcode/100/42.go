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

type helper struct {
	height int
	idx    int //横坐标
}

func trap(height []int) int {
	var ans int
	if len(height) < 3 {
		return 0
	}
	helperStack := make([]helper, 0)
	a := 0
	for ; height[a] == 0; a++ {
	}
	var lastMin int //记录参与计算的柱子的最高的高度，避免重复计算
	helperStack = append(helperStack, helper{
		height: height[a],
		idx:    a,
	})

	for i := a + 1; i < len(height); i++ {
		if height[i] == 0 {
			lastMin = 0
			continue
		}
		top := helperStack[len(helperStack)-1]

		if top.idx < i-1 { //遇到空的柱子
			if height[i] >= top.height { //当前柱子比上一个柱子高
				for height[i] >= top.height { //计算面积，顺便清除冗余柱子
					ans += (i - top.idx - 1) * (top.height - lastMin)
					lastMin = top.height
					helperStack = helperStack[:len(helperStack)-1]
					if len(helperStack) == 0 {
						break
					}
					top = helperStack[len(helperStack)-1]
				}
				if len(helperStack) > 0 {
					ans += (i - top.idx - 1) * (height[i] - lastMin)
					lastMin = height[i]
				}
				lastMin = 0
			} else { //当前柱子不比上一个柱子高
				if len(helperStack) > 0 {
					ans += (i - top.idx - 1) * (height[i] - lastMin)
					lastMin = height[i]
				}
			}

			if len(helperStack) == 0 {
				lastMin = 0
			}
			helperStack = append(helperStack, helper{
				height: height[i],
				idx:    i,
			})
			continue
		}

		if height[i] > top.height { //当前柱子高度大于栈顶元素高度
			lastMin = top.height
			helperStack = helperStack[:len(helperStack)-1]
			if len(helperStack) > 0 {
				top = helperStack[len(helperStack)-1]
				if len(helperStack) > 0 {
					for top.height == lastMin { //清除冗余柱子
						helperStack = helperStack[:len(helperStack)-1]
						if len(helperStack) == 0 {
							break
						}
						top = helperStack[len(helperStack)-1]
					}
				}
				for height[i] >= top.height { //计算接雨水的面积，直到栈为空或者遇到更高的柱子
					ans += (i - top.idx - 1) * (top.height - lastMin)
					lastMin = top.height
					if len(helperStack) == 0 || top.height > height[i] { //栈为空或者遇到更高的柱子
						break
					}
					helperStack = helperStack[:len(helperStack)-1]
					if len(helperStack) == 0 {
						break
					}
					top = helperStack[len(helperStack)-1]
				}

			}
			if len(helperStack) > 0 {
				ans += (i - top.idx - 1) * (height[i] - lastMin)
				lastMin = height[i]
			}

		} else { //当前柱子高度小于栈顶元素高度
			lastMin = 0
		}
		helperStack = append(helperStack, helper{
			height: height[i],
			idx:    i,
		})
	}
	return ans
}

func trapBest(height []int) int {
	var ans int
	stk := make([]int, 0) //存储柱子对应下标
	for i, v := range height {
		for len(stk) > 0 && v > height[stk[len(stk)-1]] {
			top := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if len(stk) > 0 {
				width := i - stk[len(stk)-1] - 1
				hei := min(height[stk[len(stk)-1]], v) - height[top]
				ans += width * hei
			}
		}
		stk = append(stk, i)
	}
	return ans
}

func main() {
	heights := pkg.CreateSlice[int]()
	fmt.Println(trap(heights))
	fmt.Println(trapBest(heights))
}
