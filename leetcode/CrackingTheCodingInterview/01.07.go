// 01.07. 旋转矩阵
/*
给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。
不占用额外内存空间能否做到？

注意：本题与主站 48 题相同：https://leetcode.cn/problems/rotate-image/
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func rotate(matrix [][]int) {
	n := len(matrix)
	left, right, top, bottom := 0, n-1, 0, n-1
	for i := 0; i < n/2; i++ {
		for curLeft := i; curLeft < n-i-1; curLeft++ {
			matrix[top][curLeft], matrix[curLeft][right], matrix[bottom][n-1-curLeft], matrix[n-1-curLeft][left] =
				matrix[n-1-curLeft][left], matrix[top][curLeft], matrix[curLeft][right], matrix[bottom][n-1-curLeft]
		}
		top++
		bottom--
		left++
		right--
	}
}

// 示例 1：
// 给定 matrix =
// [[1,2,3],[4,5,6],[7,8,9]],
//
// 原地旋转输入矩阵，使其变为:
// [[7,4,1],[8,5,2],[9,6,3]]
//
// 示例 2：
// 给定 matrix =
// [[ 5, 1, 9,11], [ 2, 4, 8,10], [13, 3, 6, 7], [15,14,12,16]]
// 原地旋转输入矩阵，使其变为:
// [[15,13, 2, 5], [14, 3, 4, 1], [12, 6, 8, 9], [16, 7,10,11]]
func main() {
	matrix := CreateSlice2D[int]()
	rotate(matrix)
	fmt.Println(matrix)
}
