/*
73. 矩阵置零
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

示例 1：
输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]

示例 2：
输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]

提示：
m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-2^31 <= matrix[i][j] <= 2^31 - 1
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

//func setZeroes(matrix [][]int) {
//	rows := len(matrix)
//	cols := len(matrix[0])
//	row2z := make(map[int]bool)
//	col2z := make(map[int]bool)
//	for i := 0; i < rows; i++ {
//		for j := 0; j < cols; j++ {
//			if matrix[i][j] == 0 {
//				row2z[i] = true
//				col2z[j] = true
//			}
//		}
//	}
//
//	for r := range row2z {
//		for j := 0; j < cols; j++ {
//			matrix[r][j] = 0
//		}
//	}
//	for c := range col2z {
//		for i := 0; i < rows; i++ {
//			matrix[i][c] = 0
//		}
//	}
//}

/*
优化：利用原矩阵第一行与第一列：
若某元素为0，则将第一行、第一列对应位置置0（x轴，y轴映射），表示该行、该列均应置0
先对其他行、列置0
最后再对第一行、第一列置0（如有必要）
对第一行、第一列作特殊处理，增加firstRow2zero firstCol2zero标识第一行、第一列元素是否应该置0
*/
func setZeroes(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])
	var firstRow2zero bool
	var firstCol2zero bool
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					firstRow2zero = true
				}
				if j == 0 {
					firstCol2zero = true
				}
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < rows; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < cols; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 1; j < cols; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < rows; i++ {
				matrix[i][j] = 0
			}
		}
	}
	if firstCol2zero == true {
		for i := 1; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
	if firstRow2zero == true {
		for j := 1; j < cols; j++ {
			matrix[0][j] = 0
		}
	}

}
func main() {
	matrix := pkg.CreateIntSlice2()
	setZeroes(matrix)
	fmt.Println(matrix)
}
