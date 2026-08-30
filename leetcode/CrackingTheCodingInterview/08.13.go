// 08.13. 堆箱子
/*
堆箱子。给你一堆n个箱子，箱子宽 wi、深 di、高 hi。箱子不能翻转，将箱子堆起来时，下面箱子的宽度、高度和深度必须大于上面的箱子。
实现一种方法，搭出最高的一堆箱子。箱堆的高度为每个箱子高度的总和。
输入使用数组[wi, di, hi]表示每个箱子。

提示:
箱子的数目不大于3000个。
*/

package main

import (
	"fmt"
	. "lc/pkg"
	"sort"
)

func pileBox(box [][]int) int {
	sort.Slice(box, func(i, j int) bool {
		return box[i][0] < box[j][0]
	})
	var ans int
	dp := make([]int, len(box))
	for i := 0; i < len(dp); i++ {
		dp[i] = box[i][2]
	}
	for i := 0; i < len(box); i++ {
		w1, d1, h1 := box[i][0], box[i][1], box[i][2]
		for j := 0; j < i; j++ {
			w2, d2, h2 := box[j][0], box[j][1], box[j][2]
			if w1 > w2 && d1 > d2 && h1 > h2 {
				dp[i] = max(dp[i], dp[j]+h1)
			}
		}
	}
	for i := 0; i < len(dp); i++ {
		ans = max(ans, dp[i])
	}
	return ans
}

// 示例 1：
//
//	输入：box = [[1, 1, 1], [2, 2, 2], [3, 3, 3]]
//	输出：6
//
// 示例 2：
//
//	输入：box = [[1, 1, 1], [2, 3, 4], [2, 6, 7], [3, 4, 5]]
//	输出：10
func main() {
	box := CreateSlice2D[int]()
	fmt.Println(pileBox(box))
}
