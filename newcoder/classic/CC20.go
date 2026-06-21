// CC20 分割回文串
/*
给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

提示：
1 <= s.length <= 16
s 仅由小写英文字母组成
*/
package main

import (
	"fmt"
)

func partition(s string) [][]string {
	n := len(s)
	isPalindrome := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPalindrome[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			isPalindrome[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			isPalindrome[i][j] = isPalindrome[i+1][j-1] && s[i] == s[j]
		}
	}

	ans := make([][]string, 0)
	var dfs func(start int, isPalindrome [][]bool, tmp []string, ans *[][]string, length int)
	dfs = func(start int, isPalindrome [][]bool, tmp []string, ans *[][]string, length int) {
		if start == n && length == n {
			*ans = append(*ans, append([]string{}, tmp...))
			return
		}
		for i := start; i < n; i++ {
			if isPalindrome[start][i] == true {
				tmp = append(tmp, s[start:i+1])
				dfs(i+1, isPalindrome, tmp, ans, length+i+1-start)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	dfs(0, isPalindrome, []string{}, &ans, 0)

	return ans
}

// 示例 1：
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
//
// 示例 2：
// 输入：s = "a"
// 输出：[["a"]]
func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(partition(s))
}
