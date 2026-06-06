// 135. 分发糖果
/*
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
你需要按照以下要求，给这些孩子分发糖果：
每个孩子至少分配到 1 个糖果。
相邻两个孩子中，评分更高的那个会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。

示例 1：
输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。

示例 2：
输入：ratings = [1,2,2]
输出：4
解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。

提示：
n == ratings.length
1 <= n <= 2 * 10^4
0 <= ratings[i] <= 2 * 10^4
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 思路：从左向右贪心，再从右向左贪心，取两次贪心的最大值。
// 具体做法：左 → 右，当 ratings[i] > ratings[i - 1]，第 i 个得到的糖果比第 i - 1 个多 1，否则得到 1 个糖果。
// 右 → 左，当 ratings[i] > ratings[i + 1]，第 i 个得到的糖果比第 i + 1 个多 1，否则得到 1 个糖果。
func candy(ratings []int) int {
	var ans int
	n := len(ratings)
	if n == 1 {
		return 1
	}
	// 左 → 右
	left := make([]int, n)
	left[0] = 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	// 右 → 左
	right := make([]int, n)
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}
	}

	for i := 0; i < n; i++ {
		ans += max(left[i], right[i])
	}
	return ans
}

// Test Case 1: [1, 0, 2]  Expected Output: 5
// Test Case 2: [1, 2, 2]  Expected Output: 4
func main() {
	ratings := pkg.CreateSlice[int]()
	fmt.Println(candy(ratings))
}
