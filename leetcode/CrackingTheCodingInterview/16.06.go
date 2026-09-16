// 16.06. 最小差
/*
给定两个整数数组a和b，计算具有最小差绝对值的一对数值（每个数组中取一个值），并返回该对数值的差

提示：
1 <= a.length, b.length <= 100000
-2147483648 <= a[i], b[i] <= 2147483647
正确结果在区间 [0, 2147483647] 内
*/
package main

import (
	"fmt"
	. "lc/pkg"
	"math"
	"sort"
)

func smallestDifference(a []int, b []int) int {
	if len(a) > len(b) {
		return handler(a, b)
	}
	return handler(b, a)
}

func handler(a, b []int) int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	var ans = math.MaxInt
	for i := 0; i < len(b); i++ {
		left, right := 0, len(a)-1
		for left <= right {
			mid := (left + right) / 2
			ans = min(abs(a[mid], b[i]), ans)
			if b[i] > a[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return ans
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}

func smallestDifference2(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	var ans = math.MaxInt
	for i, j := 0, 0; i < len(a) && j < len(b); {
		ans = min(abs(a[i], b[j]), ans)
		if a[i] < b[j] {
			i++
		} else {
			j++
		}
	}
	return ans
}

// 示例：
// 输入：[1, 3, 15, 11, 2], [23, 127, 235, 19, 8]
// 输出：3，即数值对(11, 8)
func main() {
	a := CreateSlice[int]()
	b := CreateSlice[int]()
	fmt.Println(smallestDifference2(a, b))
}
