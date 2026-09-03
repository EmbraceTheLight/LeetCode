// 10.01. 合并排序的数组
/*
给定两个排序后的数组 A 和 B，其中 A 的末端有足够的缓冲空间容纳 B。 编写一个方法，将 B 合并入 A 并排序。
初始化 A 和 B 的元素数量分别为 m 和 n。

说明：
A.length == n + m
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func merge(A []int, m int, B []int, n int) {
	left, right := m-1, n-1
	i := m + n - 1
	for ; i >= 0 && left >= 0 && right >= 0; i-- {
		if A[left] < B[right] {
			A[i] = B[right]
			right--
		} else {
			A[i] = A[left]
			left--
		}
	}
	for ; right >= 0; right, i = right-1, i-1 {
		A[i] = B[right]
	}
}

// 示例：
// 输入：
// A = [1,2,3,0,0,0], m = 3
// B = [2,5,6],       n = 3
// 输出： [1,2,2,3,5,6]
func main() {
	var m, n int
	fmt.Println("Input m,n:")
	fmt.Scan(&m, &n)
	fmt.Println("Input A:")
	a := CreateSlice[int]()
	fmt.Println("Input B:")
	b := CreateSlice[int]()
	merge(a, m, b, n)
	fmt.Println(a)
}
