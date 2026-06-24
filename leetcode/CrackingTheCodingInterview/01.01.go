// 面试题 01.01. 判定字符是否唯一
/*
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

限制：
0 <= len(s) <= 100
s[i]仅包含小写字母
如果你不使用额外的数据结构，会很加分。
*/
package main

import "fmt"

func isUnique(astr string) bool {
	var visited [26]bool
	for i := 0; i < len(astr); i++ {
		if visited[astr[i]-'a'] == true {
			return false
		}
		visited[astr[i]-'a'] = true
	}
	return true
}

func isUniqueBitCalc(astr string) bool {
	var visited int32
	for i := 0; i < len(astr); i++ {
		if visited&(1<<int(astr[i]-'a')) > 0 {
			return false
		}
		visited = visited | (1 << int(astr[i]-'a'))
	}
	return true
}

// 示例 1：
// 输入: s = "leetcode"
// 输出: false
//
// 示例 2：
// 输入: s = "abc"
// 输出: true
func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(isUnique(str))
	fmt.Println(isUniqueBitCalc(str))
}
