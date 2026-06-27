// 面试题 01.04. 回文排列
/*
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
回文串不一定是字典当中的单词。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func canPermutePalindrome(s string) bool {
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]]++
	}

	isLengthEven := len(s)%2 == 0
	cnt := 0
	for _, v := range mp {
		if v%2 != 0 {
			if isLengthEven == true {
				return false
			}
			cnt++
			if cnt > 1 {
				return false
			}
		}
	}
	return true
}

// 示例1：
// 输入："tactcoa"
// 输出：true（排列有"tacocat"、"atcocta"，等等）
func main() {
	var s string
	fmt.Println("Input s:")
	s = CreateString()
	fmt.Println(canPermutePalindrome(s))
}
