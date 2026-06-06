/*
76.最小覆盖子串
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
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
*/
package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	mp := make(map[byte]int)
	var cnt int
	var strLen = len(s) + 1
	var next = make([]int, 0) //存放找到的第二个字符的下标
	var idx int               //指向next中下一个该取的元素
	var start int             //记录s的最小覆盖子串开始下标
	
	//初始化mp映射表
	for k := range t {
		mp[t[k]]++
		cnt++ //统计子串字符个数
	}

	l := len(s)
	var left, right int
	for left = 0; left < l; left++ {
		if _, ok := mp[s[left]]; ok {
			break
		}
	}

	if left != l { //找到了
		if len(t) == 1 {
			return string(s[left])
		}
	} else {
		return ""
	}

	cnt--
	mp[s[left]]--
	start = left
	right = left + 1

	for right < l {
		if _, ok := mp[s[right]]; ok { // 找到了一个子串中的字符
			if cnt > 0 && mp[s[right]] > 0 { //还未找完子串中的对应字符，cnt--，否则cnt不变
				cnt--
			}
			mp[s[right]]--
			next = append(next, right)
		}
		for cnt == 0 {
			if right-left < strLen {
				strLen = right - left + 1
				start = left
			}
			mp[s[left]]++
			if mp[s[left]] > 0 { //如果没有有多余的子串中对应的字符，cnt+1，否则不用改变
				cnt++
			}
			left = next[idx]
			idx++
		}
		right++
	}
	if strLen > len(s) { //未找到覆盖子串
		return ""
	} else {
		return s[start : start+strLen]
	}
}
func main() {
	var str1 string
	var str2 string
	fmt.Scanf("%s\n", &str1)
	fmt.Scanf("%s\n", &str2)
	fmt.Println(minWindow(str1, str2))
}
