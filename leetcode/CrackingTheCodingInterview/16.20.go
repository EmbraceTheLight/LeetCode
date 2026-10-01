// 16.20 T9键盘
/*
在老式手机上，用户通过数字键盘输入，手机将提供与这些数字相匹配的单词列表。每个数字映射到0至4个字母。
给定一个数字序列，实现一个算法来返回匹配单词的列表。你会得到一张含有有效单词的列表。

提示：
num.length <= 1000
words.length <= 500
words[i].length == num.length
num中不会出现 0, 1 这两个数字
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func getValidT9Words(num string, words []string) []string {
	n := len(num)
	if n == 0 {
		return []string{}
	}
	var ans []string
	mp := [10][]byte{
		[]byte{},
		[]byte{},
		[]byte{'a', 'b', 'c'},
		[]byte{'d', 'e', 'f'},
		[]byte{'g', 'h', 'i'},
		[]byte{'j', 'k', 'l'},
		[]byte{'m', 'n', 'o'},
		[]byte{'p', 'q', 'r', 's'},
		[]byte{'t', 'u', 'v'},
		[]byte{'w', 'x', 'y', 'z'},
	}
	for i := 0; i < len(words); i++ {
		if judge1620(mp, num, words[i]) == true {
			ans = append(ans, words[i])
		}
	}
	return ans
}

func judge1620(mp [10][]byte, num, word string) bool {
	for i := 0; i < len(num); i++ {
		charList := mp[num[i]-'0']
		matched := false
		for j := range charList {
			if charList[j] == word[i] {
				matched = true
				break
			}
		}
		if matched == false {
			return false
		}
	}
	return true
}

// 示例 1：
// 输入：num = "8733", words = ["tree", "used"]
// 输出：["tree", "used"]
//
// 示例 2：
// 输入：num = "2", words = ["a", "b", "c", "d"]
// 输出：["a", "b", "c"]
func main() {
	var num string
	fmt.Println("Input num:")
	fmt.Scan(&num)
	fmt.Println("Input words:")
	words := CreateSlice[string]()
	fmt.Println(getValidT9Words(num, words))
}
