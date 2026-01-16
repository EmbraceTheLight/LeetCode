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

type wordState struct {
	word string
	step int
}

// 暴力解法
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
	visit := make(map[string]bool)
	queue := []wordState{{word: beginWord, step: 1}} // 求的是转换序列长度, 从 1 开始

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		visit[top.word] = true
		for i := 0; i < 26; i++ {
			for j := 0; j < n; j++ {
				tmp := top.word[:j] + chars[i] + top.word[j+1:]
				if wordMap[tmp] == false || visit[tmp] == true {
					continue
				}
				if tmp == endWord {
					return top.step + 1
				}
				queue = append(queue, wordState{word: tmp, step: top.step + 1})
				fmt.Println(wordState{word: tmp, step: top.step + 1})
			}
		}
	}
	return 0
}

// 优化版本, 不使用 map[string]bool 作为 visit 的类型, 而是使用 []bool,
// 这样直接根据单词在 wordList 中的下标判断是否遍历过该单词, 时间开销更小
// 可以通过更多的测试用例, 但是仍然超时
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	chars := []string{
		"a", "b", "c", "d", "e", "f",
		"g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u",
		"v", "w", "x", "y", "z",
	}

	n := len(beginWord)
	wordMap := make(map[string]int)
	for i, w := range wordList {
		wordMap[w] = i + 1
	}
	if _, ok := wordMap[endWord]; !ok || len(endWord) != n {
		return 0
	}

	// visit 数组表示是否遍历过该单词. 从 1 开始, visit[0] 无作用. 这样做, 当 beginWord 不在列 wordList 中时, 也可以对 visit[0] 赋值, 而不会影响到 wordList[0]
	visit := make([]bool, len(wordList)+1)

	queue := []wordState{{word: beginWord, step: 1}} // 求的是转换序列长度, 从 1 开始

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		visit[wordMap[top.word]] = true
		for i := 0; i < 26; i++ {
			for j := 0; j < n; j++ {
				tmp := top.word[:j] + chars[i] + top.word[j+1:]
				if _, ok := wordMap[tmp]; !ok {
					continue
				}
				if visit[wordMap[tmp]] == true {
					continue
				}
				if tmp == endWord {
					return top.step + 1
				}
				queue = append(queue, wordState{word: tmp, step: top.step + 1})
				fmt.Println(wordState{word: tmp, step: top.step + 1})
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
	fmt.Println(ladderLength2(beginWord, endWord, wordList))
	fmt.Println(ladderLength3(beginWord, endWord, wordList))
}
