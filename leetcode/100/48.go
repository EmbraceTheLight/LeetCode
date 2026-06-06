/*
48. 旋转图像
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]

示例 2：
输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

提示：
n == matrix.length == matrix[i].length
1 <= n <= 20
-1000 <= matrix[i][j] <= 1000
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

//	func rotate48LineBox(left, top, right, bottom, size int, matrix [][]int) {
//		for curCol := left; curCol < right; curCol++ {
//			tmp1 := matrix[curCol][right]
//			matrix[curCol][right] = matrix[top][curCol]
//
//			tmp2 := matrix[bottom][size-1-curCol]
//			matrix[bottom][size-1-curCol] = tmp1
//
//			tmp1 = matrix[size-1-curCol][left]
//			matrix[size-1-curCol][left] = tmp2
//
//			matrix[top][curCol] = tmp1
//		}
//
// }
//
//	func rotate48(matrix [][]int) {
//		size := len(matrix)
//		//一圈一圈向内旋转，每次旋转一个正方形方框
//		//正方形方框每圈边长 = 外框边长-2
//		var left, top, right, bottom = 0, 0, size - 1, size - 1
//		for i := 0; i < size/2; i++ {
//			rotate48LineBox(left, top, right, bottom, size, matrix)
//			left++
//			top++
//			right--
//			bottom--
//		}
//	}
func rotate48LineBox(left, top, right, bottom, size int, matrix [][]int) {
	for curCol := left; curCol < right; curCol++ {
		//tmp1 := matrix[curCol][right]
		//matrix[curCol][right] = matrix[top][curCol]
		//tmp2 := matrix[bottom][size-1-curCol]
		//matrix[bottom][size-1-curCol] = tmp1
		//tmp1 = matrix[size-1-curCol][left]
		//matrix[size-1-curCol][left] = tmp2
		//matrix[top][curCol] = tmp1
		matrix[top][curCol], matrix[curCol][right], matrix[bottom][size-1-curCol], matrix[size-1-curCol][left] =
			matrix[size-1-curCol][left], matrix[top][curCol], matrix[curCol][right], matrix[bottom][size-1-curCol]
	}
}
func rotate48(matrix [][]int) {
	size := len(matrix)
	//一圈一圈向内旋转，每次旋转一个正方形方框
	//正方形方框每圈边长 = 外框边长-2
	var left, top, right, bottom = 0, 0, size - 1, size - 1
	for i := 0; i < size/2; i++ {
		rotate48LineBox(left, top, right, bottom, size, matrix)
		left++
		top++
		right--
		bottom--
	}
}
func main() {
	matrix := pkg.CreateIntSlice2()
	rotate48(matrix)
	fmt.Println(matrix)
}
