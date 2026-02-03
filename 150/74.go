// 74. 搜索二维矩阵
/*
给你一个满足下述两条属性的 m x n 整数矩阵：
每行中的整数从左到右按非严格递增顺序排列。
每行的第一个整数大于前一行的最后一个整数。
给你一个整数 target，如果 target 在矩阵中，返回 true；否则，返回 false 。

提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-10^4 <= matrix[i][j], target <= 10^4
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	targetRow := 0
	for ; targetRow < rows; targetRow++ {
		if matrix[targetRow][0] <= target && target <= matrix[targetRow][cols-1] {
			break
		}
	}
	if targetRow == rows {
		return false
	}
	if searchIndex(matrix[targetRow], target) != -1 {
		return true
	}
	return false
}

// 搜索从第几行开始的逻辑也可以使用二分查找
func searchMatrixBest(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])

	bottom, top := 0, rows-1
	var mid int
	for bottom <= top {
		mid = (bottom + top) / 2
		if matrix[mid][0] <= target && target <= matrix[mid][cols-1] {
			break
		} else if matrix[mid][0] < target {
			bottom = mid + 1
		} else {
			top = mid - 1
		}
	}
	if bottom == rows || top < 0 {
		return false
	}
	if searchIndex(matrix[mid], target) != -1 {
		return true
	}
	return false
}

func searchIndex(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// Test Case1:
// matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// Output: true

// Test Case2：
// matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// Output: false
func main() {
	var target int
	fmt.Println("Input target:")
	fmt.Scan(&target)
	fmt.Println("Input matrix:")
	matrix := pkg.CreateSlice2D[int]()
	fmt.Println(searchMatrix(matrix, target))
}
