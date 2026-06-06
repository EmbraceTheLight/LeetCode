/*
54.螺旋矩阵
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func spiralOrder(matrix [][]int) []int {
	var ans = make([]int, 0)
	var left, right, top, bottom int
	bottom = len(matrix)
	right = len(matrix[0])

	var r, c int
	var cnt int
	var nums = len(matrix) * len(matrix[0])
	top++
Outer:
	for cnt < nums {
		for ; c < right; c++ {
			ans = append(ans, matrix[r][c])
			cnt++
			if cnt == nums {
				break Outer
			}
		}
		c--
		right--
		r++
		for ; r < bottom; r++ {
			ans = append(ans, matrix[r][c])
			cnt++
			if cnt == nums {
				break Outer
			}
		}
		r--
		bottom--
		c--
		for ; c >= left; c-- {
			ans = append(ans, matrix[r][c])
			cnt++
			if cnt == nums {
				break Outer
			}
		}
		c++
		left++
		r--
		for ; r >= top; r-- {
			ans = append(ans, matrix[r][c])
			cnt++
			if cnt == nums {
				break Outer
			}
		}
		r++
		top++
		c++
	}
	return ans
}
func main() {
	matrix := pkg.CreateIntSlice2()
	fmt.Println(spiralOrder(matrix))
}
