// 16.21. 交换和
/*
给定两个整数数组，请交换一对数值（每个数组中取一个数值），使得两个数组所有元素的和相等。
返回一个数组，第一个元素是第一个数组中要交换的元素，第二个元素是第二个数组中要交换的元素。
若有多个答案，返回任意一个均可。若无满足条件的数值，返回空数组。

提示：
1 <= array1.length, array2.length <= 100000
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func findSwapValues(array1 []int, array2 []int) []int {
	sum1, sum2 := 0, 0
	for _, v := range array1 {
		sum1 += v
	}
	for _, v := range array2 {
		sum2 += v
	}
	if sum1 < sum2 {
		sum1, sum2 = sum2, sum1
	}

	diff := sum1 - sum2
	if diff%2 != 0 {
		return []int{}
	}
	diff /= 2
	// sumA + b - a == sumB - b + a
	// diff = sumA - sumB == 2a - 2b
	// a - b == diff / 2 --> a = b + diff / 2
	mp := make(map[int]bool)
	for i := 0; i < len(array1); i++ {
		mp[array1[i]] = true
	}
	for i := 0; i < len(array2); i++ {
		if mp[array2[i]+diff] == true {
			return []int{array2[i] + diff, array2[i]}
		}
	}
	return []int{}
}

// 示例 1：
// 输入：array1 = [4, 1, 2, 1, 1, 2], array2 = [3, 6, 3, 3]
// 输出：[1, 3]
//
// 示例 2：
// 输入：array1 = [1, 2, 3], array2 = [4, 5, 6]
// 输出：[]
func main() {
	array1 := CreateSlice[int]()
	array2 := CreateSlice[int]()
	fmt.Println(findSwapValues(array1, array2))
}
