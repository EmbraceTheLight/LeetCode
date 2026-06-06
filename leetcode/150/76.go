// 76. 最小覆盖子串
/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。


注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。

示例 2：
输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。

示例 3:
输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。


提示：

m == s.length
n == t.length
1 <= m, n <= 10^5
s 和 t 由英文字母组成

进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？
*/
package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	ls, lt := len(s), len(t)
	var ans string
	var count = 0 // 字符串 t 中的字母种类
	minLen := math.MaxInt

	if lt > ls {
		return ""
	}

	// charMp 为 map[byte]int 类型，则运行速度会稍慢一点
	charMp := [128]int{}
	for i, _ := range t {
		if charMp[t[i]] == 0 {
			count++
		}
		charMp[t[i]]++
	}

	left := 0
	for ; left < ls; left++ {
		if charMp[s[left]] > 0 {
			break
		}
	}
	right := left
	for ; right < ls && left <= ls-lt; right++ {
		ch := s[right]
		charMp[ch]--
		if charMp[ch] == 0 {
			count--
		}

		// 移动左指针，直到 s[left: right+1] 不满足 t 中字符的频率要求
		for count == 0 {
			if right-left+1 < minLen {
				minLen = right - left + 1
				ans = s[left : right+1]
			}
			lastChar := s[left]

			// 这里如果 lastChar 为不在 t 中的字符，不会影响count。
			// 因为在 right 遍历时，charMp[lastChar] 已经小于 0 了。
			// 只有在 lastChar 出现在 t 中时，charMp[lastChar] == 0 才有可能成立。
			if charMp[lastChar] == 0 {
				count++
			}
			charMp[lastChar]++
			left++
		}
	}
	return ans
}

// Test Case 1: "ADOBECODEBANC", "ABC"	Output: "BANC"
// Test Case 2: "a", "a"	Output: "a"
// Test Case 3: "a", "aa"	Output: ""
func main() {
	var s, t string
	fmt.Scan(&s, &t)
	fmt.Println(minWindow(s, t))
}
