// 面试题 01.02 判定是否互为字符重排
/*
给定两个由小写字母组成的字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

说明：
0 <= len(s1) <= 100
0 <= len(s2) <= 100
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func CheckPermutation(s1 string, s2 string) bool {
	var mp [26]int
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		mp[s1[i]-'a']++
		mp[s2[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if mp[i] != 0 {
			return false
		}
	}
	return true
}

// 示例 1：
// 输入: s1 = "abc", s2 = "bca"
// 输出: true
//
// 示例 2：
// 输入: s1 = "abc", s2 = "bad"
// 输出: false
func main() {
	var s1, s2 string
	fmt.Println("Input s1:")
	s1 = CreateString()
	fmt.Println("Input s2:")
	s2 = CreateString()
	fmt.Println(CheckPermutation(s1, s2))
}
