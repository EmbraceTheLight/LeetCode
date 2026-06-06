// 383. 赎金信
/*
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
如果可以，返回 true ；否则返回 false 。
magazine 中的每个字符只能在 ransomNote 中使用一次。

示例 1：
输入：ransomNote = "a", magazine = "b"
输出：false

示例 2：
输入：ransomNote = "aa", magazine = "ab"
输出：false

示例 3：
输入：ransomNote = "aa", magazine = "aab"
输出：true

提示：
1 <= ransomNote.length, magazine.length <= 10^5
ransomNote 和 magazine 由小写英文字母组成

*/
package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	var charCount [26]int
	for _, ch := range magazine {
		charCount[ch-'a']++
	}

	for _, ch := range ransomNote {
		charCount[ch-'a']--
		if charCount[ch-'a'] < 0 {
			return false
		}
	}
	return true
}

// Test Case1:	"a", "b"	Output: false
// Test Case2:	"aa", "ab"	Output: false
// Test Case3:	"aa", "aab"	Output: true
func main() {
	var ransomNote, magazine string
	fmt.Scan(&ransomNote, &magazine)
	fmt.Println(canConstruct(ransomNote, magazine))
}
