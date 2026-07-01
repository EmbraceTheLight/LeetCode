// 01.08. 零矩阵
/*
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

// 示例 1：
// 输入：
// [[1,1,1],[1,0,1],[1,1,1]]
// 输出：
// [[1,0,1],[0,0,0],[1,0,1]]
//
// 示例 2：
// 输入：
// [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// 输出：
// [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
func setZeroes(matrix [][]int) {
	firstRowIsZero, firstColIsZero := matrix[0][0] == 0, matrix[0][0] == 0

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstColIsZero = true
		}
	}
	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			firstRowIsZero = true
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			for j := 0; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
	if firstRowIsZero {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}

	if firstRowIsZero {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	if firstColIsZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	matrix := CreateSlice2D[int]()
	setZeroes(matrix)
	fmt.Println(matrix)
}
