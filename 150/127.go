// 127. 单词接龙
/*

字典 wordList 中从单词 beginWord 到 endWord 的 转换序列 是一个按下述规格形成的序列 beginWord -> s1 -> s2 -> ... -> sk：
每一对相邻的单词只差一个字母。
对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。
sk == endWord
给你两个单词 beginWord 和 endWord 和一个字典 wordList，返回 从 beginWord 到 endWord 的 最短转换序列 中的 单词数目。如果不存在这样的转换序列, 返回 0。

示例 1：
输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
输出：5
解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。

示例 2：
输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
输出：0
解释：endWord "cog" 不在字典中，所以无法进行转换。


提示：

1 <= beginWord.length <= 10
endWord.length == beginWord.length
1 <= wordList.length <= 5000
wordList[i].length == beginWord.length
beginWord、endWord 和 wordList[i] 由小写英文字母组成
beginWord != endWord
wordList 中的所有字符串 互不相同
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	chars := []string{
		"a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u",
		"v", "w", "x", "y", "z",
	}
	n := len(beginWord)
	wordMap := make(map[string]bool)
	for _, w := range wordList {
		wordMap[w] = true
	}
	if wordMap[endWord] == false || len(endWord) != n {
		return 0
	}
	visit1 := make(map[string]bool)
	visit2 := make(map[string]bool)
	queue1 := []string{beginWord}
	queue2 := []string{endWord}
	count1 := 1
	count2 := 1
	for len(queue1) > 0 && len(queue2) > 0 {
		top := queue1[0]
		queue1 = queue1[1:]
		visit1[top] = true
		for i := 0; i < 26; i++ {
			for j := 0; j < n; j++ {
				tmp := top[:j] + chars[i] + top[j+1:]
				if wordMap[tmp] == false || visit1[tmp] == true {
					continue
				}
				if visit1[tmp] && visit2[tmp] {
					return count1 + count2 + 1
				}
				queue1 = append(queue1, tmp)
			}
		}
	}

	for len(queue1) > 0 && len(queue2) > 0 {
		top := queue1[0]
		queue1 = queue1[1:]
		visit1[top] = true
		for i := 0; i < 26; i++ {
			for j := 0; j < n; j++ {
				tmp := top[:j] + chars[i] + top[j+1:]
				if wordMap[tmp] == false || visit1[tmp] == true {
					continue
				}
				if visit1[tmp] && visit2[tmp] {
					return count1 + count2 + 1
				}
				queue1 = append(queue1, tmp)
			}
		}

		top = queue2[0]
		queue2 = queue2[1:]
		visit2[top] = true
		for i := 0; i < 26; i++ {
			for j := 0; j < n; j++ {
				tmp := top[:j] + chars[i] + top[j+1:]
				if wordMap[tmp] == false || visit2[tmp] == true {
					continue
				}
				if visit1[tmp] && visit2[tmp] {
					return count1 + count2
				}
				queue1 = append(queue2, tmp)
			}
		}
	}
	return 0
}

// Test Case1: beginWord: "hit"	 endWord: "cog"	 wordList: ["hot","dot","dog","lot","log","cog"]	Output: 5
// Test Case2: beginWord: "hit"	 endWord: "cog"	 wordList: ["hot","dot","dog","lot","log"]	Output: 0
// Test Case3: beginWord: "hit"	 endWord: "cog"	 wordList: []		Output: 0
func main() {
	fmt.Println("Input beginWord, endWord:")
	var beginWord, endWord string
	fmt.Scan(&beginWord, &endWord)
	fmt.Println("Input wordList:")
	wordList := pkg.CreateSlice[string]()
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
