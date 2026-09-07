// 10.05. 稀疏数组搜索
/*
稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。

提示:
words的长度在[1, 1000000]之间
*/
package main

import (
	"fmt"
	. "lc/pkg"
	"strings"
)

func findString(words []string, s string) int {
	left, right := 0, len(words)-1
	for left <= right {
		mid := (left + right) / 2
		if words[mid] == "" {
			tmpIdx := mid
			for tmpIdx <= right && words[tmpIdx] == "" {
				tmpIdx++
			}
			if tmpIdx > right { // 空字符串超过右边界, 说明右侧没有要找的答案, 向左侧寻找
				right = mid - 1
				continue
			}
			mid = tmpIdx
		}
		res := strings.Compare(words[mid], s)
		if res == -1 {
			left = mid + 1
		} else if res == 1 {
			right = mid - 1
		} else {
			ans := mid
			for tmpIdx := mid; tmpIdx >= left && (words[tmpIdx] == "" || words[tmpIdx] == s); tmpIdx-- {
				if words[tmpIdx] == s {
					ans = tmpIdx
				}
			}
			return ans
		}
	}
	return -1
}

// 示例 1：
// 输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
// 输出：-1
// 说明：不存在返回-1。
// 示例 2：
// 输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
// 输出：4
func main() {
	var s string
	fmt.Println("Input target string:")
	s = CreateString()
	fmt.Println("Input array:")
	words := CreateSlice[string]()
	fmt.Println(findString(words, s))
}
