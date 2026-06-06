// 54. 螺旋矩阵
/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

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
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	count := len(matrix) * len(matrix[0])
	ans := make([]int, 0)
outer:
	for {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
			count--
			if count == 0 {
				break outer
			}
		}
		top++

		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
			count--
			if count == 0 {
				break outer
			}
		}
		right--

		for i := right; i >= left; i-- {
			ans = append(ans, matrix[bottom][i])
			count--
			if count == 0 {
				break outer
			}
		}
		bottom--

		for i := bottom; i >= top; i-- {
			ans = append(ans, matrix[i][left])
			count--
			if count == 0 {
				break outer
			}
		}
		left++
	}
	return ans
}

// 示例 1：
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]
//
// 示例 2：
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
func main() {
	matrix := pkg.CreateSlice2D[int]()
	fmt.Println(spiralOrder(matrix))
}
