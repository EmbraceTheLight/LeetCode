/*
131. 分割回文串
给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是
回文串。返回 s 所有可能的分割方案。

示例 1：
输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]

示例 2：
输入：s = "a"
输出：[["a"]]

提示：

1 <= s.length <= 16
s 仅由小写英文字母组成
*/
package main

import "fmt"

func isPalindrome(s string) bool {
	l := len(s)
	itLimit := l / 2
	for i := 0; i < itLimit; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}
func dfs131(s string, ans *[][]string, sub []string, idx int) {
	if idx == len(s) {
		var s = make([]string, len(sub))
		copy(s, sub)
		*ans = append(*ans, s)
		return
	}
	for i := idx; i < len(s); i++ {
		if isPalindrome(s[idx : i+1]) {
			sub = append(sub, s[idx:i+1])
			dfs131(s, ans, sub, i+1)
			sub = sub[:len(sub)-1]
		}
	}
}

func partition(s string) [][]string {
	var ans = make([][]string, 0)
	l := len(s)
	for i := 0; i < l; i++ {
		if isPalindrome(s[0 : i+1]) {
			var sub = make([]string, 0)
			sub = append(sub, s[0:i+1])
			dfs131(s, &ans, sub, i+1)
		}
	}
	return ans
}
func main() {
	fmt.Println("Input a string:")
	var input string
	fmt.Scan(&input)
	for _, v := range partition(input) {
		for _, v := range v {
			fmt.Printf("%s ", v)
		}
		fmt.Println()
	}
}
