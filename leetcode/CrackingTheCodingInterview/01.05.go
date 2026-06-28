// 01.05. 一次编辑
/*
字符串有三种编辑操作:插入一个英文字符、删除一个英文字符或者替换一个英文字符。
给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
*/
package main

import (
	"fmt"
	"math"
)

func oneEditAway(first string, second string) bool {
	n1, n2 := len(first), len(second)
	if int(math.Abs(float64(n1-n2))) >= 2 {
		return false
	}
	idx1, idx2 := 0, 0
	cnt := 0
	for idx1 < n1 && idx2 < n2 {
		if first[idx1] == second[idx2] {
			idx1++
			idx2++
		} else if first[idx1] != second[idx2] && (idx1+1 < n1 && first[idx1+1] == second[idx2]) {
			cnt++
			if cnt > 1 {
				return false
			}
			idx1++
		} else if first[idx1] != second[idx2] && (idx2+1 < n2 && first[idx1] == second[idx2+1]) {
			cnt++
			if cnt > 1 {
				return false
			}
			idx2++
		} else {
			cnt++
			if cnt > 1 {
				return false
			}
			idx1++
			idx2++
		}
	}
	if idx1 != n1 || idx2 != n2 {
		cnt++
	}
	if cnt > 1 {
		return false
	}
	return true
}

// 示例 1：
// 输入：
// first = "pale"
// second = "ple"
// 输出：True
//
// 示例 2：
// 输入：
// first = "pales"
// second = "pal"
// 输出：False
func main() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)
	fmt.Println(oneEditAway(str1, str2))
}
