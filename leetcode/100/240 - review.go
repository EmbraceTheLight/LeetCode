/*
240. 搜索二维矩阵 II

编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
输出：true

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 20
输出：false

提示：

m == matrix.length
n == matrix[i].length
1 <= n, m <= 300
-10^9 <= matrix[i][j] <= 10^9
每行的所有元素从左到右升序排列
每列的所有元素从上到下升序排列
-10^9 <= target <= 10^9
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func searchMatrixR(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	for i, j := 0, n-1; i < m && j >= 0; { //从右上角开始遍历，每次可排除一行/一列
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target { //说明第j列的数都比target大，故排除第j列
			j--
		} else { //说明第i行的数都比target小，故排除第i行
			i++
		}
	}
	return false
}
func main() {
	fmt.Println("Input target:")
	var target int
	fmt.Scan(&target)

	fmt.Println("Input matrix:")
	matrix := pkg.CreateIntSlice2()
	fmt.Println(searchMatrixR(matrix, target))
}
