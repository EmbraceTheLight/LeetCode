/*
240. 搜索二维矩阵 II
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

示例1：
输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
输出：true

示例2：
输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
输出：false

提示：
m == matrix.length
n == matrix[i].length
1 <= n, m <= 300
-10^9 <= matrix[i][j] <= 10^9
每行的所有元素从左到右升序排列
每列的所有元素从上到下升序排列
-109 <= target <= 10^9
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func searchMatrix(matrix [][]int, target int) bool {
	r := len(matrix)
	c := len(matrix[0])
	for i := 0; i < r && matrix[i][0] <= target; i++ {
		left, right := 0, c-1
		for left <= right {
			mid := int(uint(left+right) >> 1)
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

// 从矩阵左下角开始搜索，一次可以排除一行/一列，二分一次只能排除半行/半列
func searchMatrixBetter(matrix [][]int, target int) bool {
	r := len(matrix)
	c := len(matrix[0])
	for i, j := r-1, 0; i >= 0 && j < c; {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			i--
		} else {
			j++
		}
	}
	return false
}
func main() {
	var target int
	fmt.Println("Input target:")
	fmt.Scan(&target)
	matrix := pkg.CreateIntSlice2()
	fmt.Println(searchMatrix(matrix, target))
	fmt.Println(searchMatrixBetter(matrix, target))
}
