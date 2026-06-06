/*
3. 无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

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

import (
	"bufio"
	"fmt"
	"os"
)

//	func lengthOfLongestSubstring(s string) int {
//		if len(s) == 0 {
//			return 0
//		}
//		var left, right int
//		var ans = 1
//		mp := make(map[byte]int)
//
//		l := len(s)
//		for right < l {
//			if _, ok := mp[s[right]]; !ok { //没有对应字符
//				mp[s[right]] = right
//				right++
//			} else {
//				ans = max(right-left, ans)
//				tmp := mp[s[right]]
//				for a := left; a < mp[s[right]]+1; a++ {
//					delete(mp, s[a])
//				}
//				left = tmp + 1
//			}
//		}
//		ans = max(right-left, ans)
//		return ans
//	}
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var left, right int
	var ans = 1
	mp := make([]int, 128)

	l := len(s)
	for right < l {

		mp[s[right]]++
		if mp[s[right]] > 1 {
			ans = max(right-left, ans)
			for ; mp[s[right]] > 1; left++ {
				mp[s[left]]--
			}
		}
		right++
	}
	ans = max(right-left, ans)
	return ans
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Input s:")
	result, _, _ := reader.ReadLine()
	fmt.Println(lengthOfLongestSubstring(string(result)))
}
