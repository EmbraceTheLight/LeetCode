// 118. 杨辉三角
/*
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
在「杨辉三角」中，每个数是它左上方和右上方的数的和。

提示:
1 <= numRows <= 30
*/
package main

import "fmt"

func generate(numRows int) [][]int {
	var ans [][]int
	ans = append(ans, []int{}) // 插入一个空元素, 从下标为 1 开始
	for i := 1; i <= numRows; i++ {
		tmp := make([]int, i)
		tmp[0], tmp[i-1] = 1, 1
		for j := 1; j < i-1; j++ {
			tmp[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		ans = append(ans, tmp)
	}
	return ans[1:]
}

// 示例 1:
// 输入: numRows = 5
// 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
//
// 示例 2:
// 输入: numRows = 1
// 输出: [[1]]
func main() {
	var numRows int
	fmt.Scan(&numRows)
	fmt.Println(generate(numRows))
}
