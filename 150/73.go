// 73. 矩阵置零
/*
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


进阶：
一个直观的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个仅使用常量空间的解决方案吗？
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 思路：遍历矩阵，遍历到的点如果位 0，则可以将其“投影”到第一行第一列，将第一行、第一列对应位置的数设置为 0
// 若点的行为 0，则不设置第一行对应位置，若列为 0，则不设置第一列对应位置
func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	firstRow, firstCol := false, false //记录第一行、第一列是否有 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				// 第一行、第一列要特殊处理
				if i == 0 {
					firstRow = true
				}
				matrix[i][0] = 0

				if j == 0 {
					firstCol = true
				}
				matrix[0][j] = 0

			}
		}
	}

	for i := 1; i < rows; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < cols; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 1; j < cols; j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < rows; i++ {
				matrix[i][j] = 0
			}
		}
	}

	// 将第一行置 0
	if firstRow {
		for j := 0; j < cols; j++ {
			matrix[0][j] = 0
		}

	}

	// 将第一列置 0
	if firstCol {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
}

// Test Case1:
/*
[[1,1,1]
,[1,0,1]
,[1,1,1]]
*/
// Output:	[[1,0,1],[0,0,0],[1,0,1]]

// Test Case2:
/*
[[0,1,2,0]
,[3,4,5,2]
,[1,3,1,5]]
*/
// Output:	[[0,0,0,0],[0,4,5,0],[0,3,1,0]]

// Test Case3:
/*
[[1,0]]
*/
// Output:	[[0,0]]

// Test Case4:
/*
[[1,2,3,4]
,[5,0,7,8]
,[0,10,11,12]
,[13,14,15,0]]
*/
// Output:	[[0,0,3,0],[0,0,0,0],[0,0,0,0],[0,0,0,0]]
func main() {
	matrix := pkg.CreateSlice2D[int]()
	setZeroes(matrix)
	fmt.Println(matrix)
}
