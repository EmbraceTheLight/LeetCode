/*
54. 螺旋矩阵
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

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
	"lc/100/pkg"
)

func spiralOrder54r(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	ans := make([]int, 0)

	var left, right, top, bottom int
	left, right, top, bottom = 0, n-1, 0, m-1
	count := m * n
	for {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
			count--
		}
		if count == 0 {
			break
		}
		top++

		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
			count--
		}
		if count == 0 {
			break
		}
		right--

		for i := right; i >= left; i-- {
			ans = append(ans, matrix[bottom][i])
			count--
		}
		if count == 0 {
			break
		}
		bottom--

		for i := bottom; i >= top; i-- {
			ans = append(ans, matrix[i][left])
			count--
		}
		if count == 0 {
			break
		}
		left++
	}
	return ans
}
func main() {
	matrix := pkg.CreateIntSlice2()
	fmt.Println(spiralOrder54r(matrix))
}
