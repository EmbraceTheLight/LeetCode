// 14. 最长公共前缀
/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。


示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"

示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

提示：
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 如果非空，则仅由小写英文字母组成
*/
package main

import (
	"fmt"
	"lc/pkg"
	"strings"
)

func longestCommonPrefixReview(strs []string) string {
	n := len(strs)
	minLen := len(strs[0])
	var sb strings.Builder
	for i := 1; i < n; i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}

Outer:
	for i := 0; i < minLen; i++ {
		char := strs[0][i]
		for j := 1; j < n; j++ {
			if strs[j][i] != char {
				break Outer
			}
		}
		sb.WriteByte(char)
	}
	return sb.String()
}

// Test Case1: ["flower","flow","flight"]  Output: "fl"
// Test Case2: ["dog","racecar","car"]  Output: ""
func main() {
	strs := pkg.CreateSlice[string]()
	fmt.Println(strs)
	fmt.Println(longestCommonPrefixReview(strs))
}
