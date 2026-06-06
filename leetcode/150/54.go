// 54. 螺旋矩阵
/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]


提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/

package main

import (
	"fmt"
	"lc/pkg"
)

func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])
	var ans []int

	count := rows * cols
	left, right, top, bottom := 0, cols-1, 0, rows-1
	for count > 0 {
		for i := left; i <= right && count > 0; i++ {
			ans = append(ans, matrix[top][i])
			count--
		}
		top++

		for i := top; i <= bottom && count > 0; i++ {
			ans = append(ans, matrix[i][right])
			count--
		}
		right--

		for i := right; i >= left && count > 0; i-- {
			ans = append(ans, matrix[bottom][i])
			count--
		}
		bottom--

		for i := bottom; i >= top && count > 0; i-- {
			ans = append(ans, matrix[i][left])
			count--
		}
		left++
	}

	return ans
}

// Test Case1: [[1,2,3],[4,5,6],[7,8,9]]	Output: [1,2,3,6,9,8,7,4,5]
// Test Case2: [[1,2,3,4],[5,6,7,8],[9,10,11,12]]	Output: [1,2,3,4,8,12,11,10,9,5,6,7]
func main() {
	matrix := pkg.CreateSlice2D[int]()
	fmt.Println(spiralOrder(matrix))
}
