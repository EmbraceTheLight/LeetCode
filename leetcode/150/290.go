// 290. 单词规律
/*
给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。
这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。具体来说：
pattern 中的每个字母都 恰好 映射到 s 中的一个唯一单词。
s 中的每个唯一单词都 恰好 映射到 pattern 中的一个字母。
没有两个字母映射到同一个单词，也没有两个单词映射到同一个字母。

示例1:
输入: pattern = "abba", s = "dog cat cat dog"
输出: true

示例 2:
输入:pattern = "abba", s = "dog cat cat fish"
输出: false

示例 3:
输入: pattern = "aaaa", s = "dog cat cat dog"
输出: false

提示:
1 <= pattern.length <= 300
pattern 只包含小写英文字母
1 <= s.length <= 3000
s 只包含小写英文字母和 ' '
s 不包含 任何前导或尾随对空格
s 中每个单词都被 单个空格 分隔
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	sli := strings.Fields(s)

	if len(sli) != len(pattern) {
		return false
	}
	p2s := make(map[byte]string)
	s2p := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		ch := pattern[i]
		str := sli[i]
		if _, ok := p2s[ch]; ok && p2s[ch] != str {
			return false
		}

		if _, ok := s2p[str]; ok && s2p[str] != ch {
			return false
		}
		s2p[str] = ch
		p2s[ch] = str
	}
	return true
}

// Test Case1:	"abba" "dog cat cat dog" 	Output: true
// Test Case2:	"abba" "dog cat cat fish" 	Output: false
// Test Case3:	"aaaa" "dog cat cat dog" 	Output: false
func main() {
	var pattern, s string
	fmt.Scan(&pattern)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s = scanner.Text()
	}
	fmt.Println(wordPattern(pattern, s))
}
