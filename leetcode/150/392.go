// 392. 判断子序列
/*
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
进阶：
如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

致谢：
特别感谢 @pbrother 添加此问题并且创建所有测试用例。

示例 1：
输入：s = "abc", t = "ahbgdc"
输出：true

示例 2：
输入：s = "axc", t = "ahbgdc"
输出：false

提示：
0 <= s.length <= 100
0 <= t.length <= 10^4
两个字符串都只由小写字符组成。
*/
package main

import "fmt"

func isSubsequence(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	idxS, idxT := 0, 0
	for ; idxS < ls && idxT < lt; idxT++ {
		if s[idxS] == t[idxT] {
			idxS++
		}
	}
	return idxS == ls
}

// Test Case 1: s = "abc", t = "ahbgdc"		output: true
// Test Case 2: s = "axc", t = "ahbgdc"		output: false
func main() {
	var s, t string
	fmt.Scan(&s, &t)
	fmt.Println(isSubsequence(s, t))
}
