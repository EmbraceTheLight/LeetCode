// 125. 验证回文串
/*
如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串。
字母和数字都属于字母数字字符。
给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。

示例 1：
输入: s = "A man, a plan, a canal: Panama"
输出：true
解释："amanaplanacanalpanama" 是回文串。

示例 2：
输入：s = "race a car"
输出：false
解释："raceacar" 不是回文串。

示例 3：
输入：s = " "
输出：true
解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
由于空字符串正着反着读都一样，所以是回文串。

提示：

1 <= s.length <= 2 * 10^5
s 仅由可打印的 ASCII 字符组成
*/
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	var bb bytes.Buffer
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			bb.WriteByte(byte(char + 32))
		} else if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			bb.WriteByte(byte(char))
		}
	}
	str := bb.String()
	n := len(str)
	if n <= 1 {
		return true
	}
	start := n/2 - 1
	for ; start >= 0; start-- {
		if str[start] != str[n-1-start] {
			return false
		}
	}
	return true
}

// 空间复杂度 O(1)，原地做法
func isPalindromeBetter(s string) bool {
	isalnum := func(c byte) bool {
		return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
	}
	n := len(s)
	left, right := 0, n-1
	for left < right {
		if !isalnum(s[left]) {
			left++
		} else if !isalnum(s[right]) {
			right--
		} else {
			if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
				return false
			}
			left++
			right--
		}
	}
	return true
}

// Test Case 1: "A man, a plan, a canal: Panama"	Output: true
// Test Case 2: "race a car"	Output: false
// Test Case 3: " "		Output: true
func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(isPalindrome(s))
}
