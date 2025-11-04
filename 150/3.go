// 3. 无重复字符的最长子串
/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

提示：
0 <= s.length <= 5 * 10^4
s 由英文字母、数字、符号和空格组成
*/
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var ans int
	n := len(s)
	mp := make([]byte, 128)
	left, right := 0, 0
	for ; right < n; right++ {
		mp[s[right]]++
		for mp[s[right]] > 1 {
			mp[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

// Test Case 1: "abcabcbb"	Output: 3
// Test Case 2: "bbbbb"		Output: 1
// Test Case 3: "pwwkew"	Output: 3
func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(lengthOfLongestSubstring(s))
}
