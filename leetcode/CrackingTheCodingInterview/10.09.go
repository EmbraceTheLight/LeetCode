// 10.09. 排序矩阵查找
/*
给定 M×N 矩阵，每一行、每一列都按升序排列，请编写代码找出某元素。

示例：
现有矩阵 matrix 如下：
[[1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]]
给定 target = 5，返回 true。
给定 target = 20，返回 false。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix)-1, 0
	for row >= 0 && col <= len(matrix[0])-1 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			row--
		} else {
			col++
		}
	}
	return false
}

// 示例 1:
// 输入: [[1, 4, 7, 11, 15], [2, 5, 8, 12, 19], [3, 6, 9, 16, 22], [10, 13, 14, 17, 24], [18, 21, 23, 26, 30]], target = 5
// 输出: true
//
// 示例 2:
// 输入: [[1, 4, 7, 11, 15], [2, 5, 8, 12, 19], [3, 6, 9, 16, 22], [10, 13, 14, 17, 24], [18, 21, 23, 26, 30]], target = 20
// 输出: false
func main() {
	var target int
	fmt.Println("Input target:")
	fmt.Scan(&target)
	matrix := CreateSlice2D[int]()
	fmt.Println(searchMatrix(matrix, target))
}
