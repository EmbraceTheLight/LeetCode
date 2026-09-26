// 16.16. 部分排序
/*
给定一个整数数组，编写一个函数，找出索引m和n，只要将索引区间[m,n]的元素排好序，整个数组就是有序的。
注意：n-m尽量最小，也就是说，找出符合条件的最短序列。
函数返回值为[m,n]，若不存在这样的m和n（例如整个数组是有序的），请返回[-1,-1]。

提示：
0 <= len(array) <= 1000000
*/
package main

import (
	"fmt"
	. "lc/pkg"
	"math"
	"sort"
)

func subSort(array []int) []int {
	if len(array) < 2 {
		return []int{-1, -1}
	}
	left, right := -1, -1
	i := 1
	for ; i < len(array); i++ {
		if array[i] < array[i-1] {
			break
		}
	}
	if i == len(array) {
		return []int{-1, -1}
	}
	i = sort.SearchInts(array[:i], array[i])
	j := len(array) - 2
	for ; j >= i; j-- {
		if array[j] > array[j+1] {
			break
		}
	}

	tmpMin := array[i]
	tmpMax := array[i]
	for idx := i; idx <= j; idx++ {
		tmpMax = max(tmpMax, array[idx])
		tmpMin = min(tmpMin, array[idx])
	}
	left = sort.SearchInts(array[:i], tmpMin)
	right = j + sort.SearchInts(array[j:], tmpMax) - 1
	for ; left <= right && array[left] == tmpMin; left++ {
	}
	for ; right >= left && array[right] == tmpMax; right-- {
	}
	return []int{left, right}
}

// 未排序部分左侧都比未排序中的元素小; 未排序部分右侧都比未排序中的元素大
func subSort2(array []int) []int {
	if len(array) < 2 {
		return []int{-1, -1}
	}
	n := len(array)
	minVal := math.MaxInt
	maxVal := math.MinInt
	left, right := -1, -1
	// 寻找右侧边界: maxVal 对应的数组下标永远小于 i, 因此若 array[i] < maxVal, 说明此处没有排序, 更新 right
	for i := 0; i < n; i++ {
		if array[i] < maxVal {
			right = i
		} else {
			maxVal = array[i]
		}
	}

	// 寻找左侧边界: minVal 对应的数组下标永远大于 i, 因此若 array[i] > minVal, 说明此处没有排序, 更新 right
	for i := n - 1; i >= 0; i-- {
		if array[i] > minVal {
			left = i
		} else {
			minVal = array[i]
		}
	}
	return []int{left, right}
}

// 示例：
// 输入： [1,2,4,7,10,11,7,12,6,7,16,18,19]
// 输出： [3,9]
func main() {
	arr := CreateSlice[int]()
	fmt.Println(subSort(arr))
}
