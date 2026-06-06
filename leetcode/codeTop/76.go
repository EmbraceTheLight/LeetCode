// 76. 最小覆盖子串
/*
给定两个字符串 s 和 t，长度分别是 m 和 n，返回 s 中的 最短窗口 子串，
使得该子串包含 t 中的每一个字符（包括重复字符）。如果没有这样的子串，返回空字符串 ""。
测试用例保证答案唯一。

提示：
m == s.length
n == t.length
1 <= m, n <= 105
s 和 t 由英文字母组成


进阶：你能设计一个在 O(m + n) 时间内解决此问题的算法吗？
*/
package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	var ans string
	var mp [128]int
	for i := 0; i < len(t); i++ {
		mp[t[i]]++
	}
	charCount := 0
	for i := 0; i < len(mp); i++ {
		if mp[i] > 0 {
			charCount++
		}
	}
	left, right := 0, 0
	minLength := math.MaxInt
	for right < len(s) {
		for ; right < len(s) && charCount > 0; right++ {
			mp[s[right]]--
			if mp[s[right]] == 0 {
				charCount--
			}
		}
		if charCount == 0 && right-left < minLength {
			ans = s[left:right]
			minLength = len(ans)
		}
		for ; left <= right && left < len(s); left++ {
			mp[s[left]]++
			if mp[s[left]]-1 == 0 {
				if charCount == 0 && right-left < minLength {
					ans = s[left:right]
					minLength = len(ans)
				}
				charCount++
				left++
				break
			}
		}
	}

	return ans
}

// 示例 1：
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//
// 示例 2：
// 输入：s = "a", t = "a"
// 输出："a"
// 解释：整个字符串 s 是最小覆盖子串。
//
// 示例 3:
// 输入: s = "a", t = "aa"
// 输出: ""
// 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
// 因此没有符合条件的子字符串，返回空字符串。
func main() {
	var s, t string
	fmt.Scan(&s, &t)
	fmt.Println(minWindow(s, t))
}
