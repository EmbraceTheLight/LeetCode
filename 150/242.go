// 242. 有效的字母异位词
/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的 字母异位词。

示例 1:
输入: s = "anagram", t = "nagaram"
输出: true

示例 2:
输入: s = "rat", t = "car"
输出: false

提示:
1 <= s.length, t.length <= 5 * 10^4
s 和 t 仅包含小写字母


进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
*/
package main

import "fmt"

func isAnagram(s string, t string) bool {
	mp := make(map[rune]int)
	for _, ch := range s {
		mp[ch]++
	}

	for _, ch := range t {
		if _, ok := mp[ch]; ok {
			mp[ch]--
			if mp[ch] == 0 {
				delete(mp, ch)
			}
		} else {
			return false
		}
	}
	return len(mp) == 0
}

// 不使用 map 而是使用数组实现映射关系
func isAnagramBest(s string, t string) bool {
	var mp [26]int
	for i := 0; i < len(s); i++ {
		mp[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		mp[s[i]-'a']--
		if mp[s[i]-'a'] < 0 {
			return false
		}
	}

	for i := 0; i < 26; i++ {
		if mp[i] != 0 {
			return false
		}
	}
	return true
}

// Test Case1:	"anagram", "nagaram"	Output: true
// Test Case2:	"rat", "car"	Output: false
func main() {
	var s, t string
	fmt.Scan(&s, &t)
	fmt.Println(isAnagram(s, t))
}
