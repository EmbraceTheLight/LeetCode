// 10.03. 搜索旋转数组
/*
搜索旋转数组。给定一个排序后的数组，包含n个整数，但这个数组已被旋转过很多次了，次数不详。
请编写代码找出数组中的某个元素，假设数组元素原先是按升序排列的。若有多个相同元素，返回索引值最小的一个。

提示:
arr 长度范围在[1, 1000000]之间
*/
package main

import (
	"fmt"
	. "lc/pkg"
	"math"
	"slices"
)

func search(arr []int, target int) int {
	ans := dfs1003(arr, 0, len(arr)-1, target)
	if ans == math.MaxInt {
		return -1
	} else {
		return ans
	}
}

func dfs1003(arr []int, left, right, target int) int {
	var ans = math.MaxInt
	if left > right {
		return ans
	}
	if left == right {
		if arr[left] == target {
			return left
		}
		return math.MaxInt
	}
	if arr[left] < arr[right] {
		idx, found := slices.BinarySearch(arr[left:right+1], target)
		if found == false {
			return math.MaxInt
		}
		return left + idx
	}

	mid := (left + right) / 2

	// 优先搜索索引更小的左半区
	leftAns := dfs1003(arr, left, mid, target)
	if leftAns != math.MaxInt {
		return leftAns
	}

	// 左半区不存在目标，再搜索右半区
	return dfs1003(arr, mid+1, right, target)
}

// 示例 1：
// 输入：arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 5
// 输出：8（元素5在该数组中的索引）
// 示例 2：
// 输入：arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 11
// 输出：-1 （没有找到）
func main() {
	fmt.Println("input target:")
	var target int
	fmt.Scan(&target)
	fmt.Println("input arr:")
	arr := CreateSlice[int]()
	fmt.Println(search(arr, target))
}
