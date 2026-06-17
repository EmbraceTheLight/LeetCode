// CC16 分糖果
/*
描述
有N个小朋友站在一排，每个小朋友都有一个评分
你现在要按以下的规则给孩子们分糖果:
每个小朋友至少要分得一颗糖果
分数高的小朋友要他比旁边得分低的小朋友分得的糖果多
你最少要分发多少颗糖果？
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func candy(ratings []int) int {
	n := len(ratings)
	if n == 1 {
		return 1
	}
	left, right := make([]int, n), make([]int, n)
	left[0] = 1
	right[n-1] = 1

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += max(left[i], right[i])
	}
	return ans
}

// 示例1
// 输入: [1,2,2]
// 返回值: 4
func main() {
	nums := CreateSlice[int]()
	fmt.Println(candy(nums))
}
