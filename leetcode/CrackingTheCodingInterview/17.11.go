// 17.11 单词距离
/*
有个内含单词的超大文本文件，给定任意两个不同的单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。
如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同，你能对此优化吗?

提示：
words.length <= 100000
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func findClosest(words []string, word1 string, word2 string) int {
	word1Idx, word2Idx := -1, -1
	var ans int = len(words)
	for i := 0; i < len(words); i++ {
		if words[i] == word1 {
			word1Idx = i
			if word2Idx != -1 {
				ans = min(ans, i-word2Idx)
			}
		} else if words[i] == word2 {
			word2Idx = i
			if word1Idx != -1 {
				ans = min(ans, i-word1Idx)
			}
		}
	}
	return ans
}

// 示例：
// 输入：words = ["I","am","a","student","from","a","university","in","a","city"], word1 = "a", word2 = "student"
// 输出：1
func main() {
	var word1, word2 string
	fmt.Println("Input word1,word2:")
	fmt.Scan(&word1, &word2)
	words := CreateSlice[string]()
	fmt.Println(findClosest(words, word1, word2))
}
