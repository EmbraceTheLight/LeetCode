// 30. 串联所有单词的子串
/*
给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串长度相同。

 s 中的 串联子串 是指一个包含  words 中所有字符串以任意顺序排列连接起来的子串。

例如，如果 words = ["ab","cd","ef"]， 那么 "abcdef"， "abefcd"，"cdabef"， "cdefab"，"efabcd"， 和 "efcdab" 都是串联子串。 "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。
返回所有串联子串在 s 中的开始索引。你可以以 任意顺序 返回答案。



示例 1：
输入：s = "barfoothefoobarman", words = ["foo","bar"]
输出：[0,9]
解释：因为 words.length == 2 同时 words[i].length == 3，连接的子字符串的长度必须为 6。
子串 "barfoo" 开始位置是 0。它是 words 中以 ["bar","foo"] 顺序排列的连接。
子串 "foobar" 开始位置是 9。它是 words 中以 ["foo","bar"] 顺序排列的连接。
输出顺序无关紧要。返回 [9,0] 也是可以的。

示例 2：
输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
输出：[]
解释：因为 words.length == 4 并且 words[i].length == 4，所以串联子串的长度必须为 16。
s 中没有子串长度为 16 并且等于 words 的任何顺序排列的连接。
所以我们返回一个空数组。

示例 3：
输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
输出：[6,9,12]
解释：因为 words.length == 3 并且 words[i].length == 3，所以串联子串的长度必须为 9。
子串 "foobarthe" 开始位置是 6。它是 words 中以 ["foo","bar","the"] 顺序排列的连接。
子串 "barthefoo" 开始位置是 9。它是 words 中以 ["bar","the","foo"] 顺序排列的连接。
子串 "thefoobar" 开始位置是 12。它是 words 中以 ["the","foo","bar"] 顺序排列的连接。

提示：

1 <= s.length <= 10^4
1 <= words.length <= 5000
1 <= words[i].length <= 30
words[i] 和 s 由小写英文字母组成
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 超时。dfs 会消耗大量时间
func findSubstring(s string, words []string) []int {
	lenS := len(s)
	lenWords := len(words)
	lenWord := len(words[0])
	if lenS < lenWords*lenWord {
		return []int{}
	}
	var ans []int

	wordMp := make(map[string]bool)
	for _, word := range words {
		wordMp[word] = true
	}

	seen := make(map[string]bool)
	var dfs func(words []string, step int, str string, visit map[int]bool)
	dfs = func(words []string, step int, str string, visit map[int]bool) {
		if step == len(words) {
			seen[str] = true
			return
		}
		for i := 0; i < len(words); i++ {
			if visit[i] == true {
				continue
			}
			newStr := str + words[i]
			visit[i] = true
			dfs(words, step+1, newStr, visit)
			visit[i] = false
		}
	}
	dfs(words, 0, "", map[int]bool{})

	var curLen int // 当前子串长度
	left, preRight := 0, 0

	// 以 s[right, right+lenWord-1] 为当前子串,以 left 为子串起点，
	// 找到长度 curLen 为 lenWords*lenWord 的 left、right 下标。不一定为串联子串，因为可能找到了几个重复的单词
	for right := preRight + lenWord; right <= lenS; right = preRight + lenWord {

		curStr := s[preRight:right]

		// 找到了串联子串的一部分，即 words 中某个单词
		if wordMp[curStr] == true {
			curLen += lenWord
			if curLen == lenWords*lenWord {
				if seen[s[left:right]] {
					ans = append(ans, left)
				}
				// 滑动窗口整体右移
				left = left + 1
				preRight = left
				curLen = 0
			} else {
				preRight = right
			}

			// 在找到串联子串之前断开了，则重置 curLen，更新 left、right，重新开始寻找。
		} else {
			curLen = 0
			left++
			preRight = left
		}
	}
	return ans
}

// 不使用 dfs 求出所有排列组合，而是使用一个 map 记录 words 中每个单词的出现次数，
// 然后遍历 s，对于 s[i, i+lenWord-1]，判断是否在 words 中
// 使用另一个 map 记录 s 的子串中出现的单词频率，使用另一个变量记录子串出现的 words 中单词的个数，时间表现不佳。因为遍历 s 没有处理到位
func findSubstringBetter(s string, words []string) []int {
	lenS := len(s)
	lenWords := len(words)
	lenWord := len(words[0])
	if lenS < lenWords*lenWord {
		return []int{}
	}
	var ans []int

	countWords := 0
	wordMp := make(map[string]int)
	for _, word := range words {
		if _, ok := wordMp[word]; !ok {
			countWords++
		}
		wordMp[word]++
	}

	var curLen int // 当前子串长度
	var wordCount int
	left, preRight := 0, 0

	wordFreq := make(map[string]int)

	// 以 s[right, right+lenWord-1] 为当前子串,以 left 为子串起点，
	// 找到长度 curLen 为 lenWords*lenWord 的 left、right 下标。不一定为串联子串，因为可能找到了几个重复的单词
	for right := preRight + lenWord; right <= lenS; right = preRight + lenWord {
		curStr := s[preRight:right]

		// 找到了串联子串的一部分，即 words 中某个单词
		if _, ok := wordMp[curStr]; ok {
			curLen += lenWord
			wordFreq[curStr]++
			if wordFreq[curStr] == wordMp[curStr] {
				wordCount++
			}

			// 子串长度 == lenWords*lenWord，且子串中出现的 words 中单词的个数 == countWords，则找到一个串联子串
			// 无论如何，重置状态，以便下一个子串从 left + 1 开始搜索
			if curLen == lenWords*lenWord {
				if wordCount == countWords {
					ans = append(ans, left)
				}

				// 窗口右移，注意重置状态。
				// ! 这里重置的太彻底了，导致时间复杂度较高
				left = left + 1
				wordFreq = make(map[string]int)
				wordCount = 0
				preRight = left
				curLen = 0
			} else {
				// 遍历下一个单词
				preRight = right
			}

			// 在找到串联子串之前断开了，则重置 curLen，更新 left、right，重新开始寻找。
		} else {
			curLen = 0
			left++
			wordFreq = make(map[string]int)
			wordCount = 0
			preRight = left
		}
	}
	return ans
}

func findSubstringBest(s string, words []string) []int {
	lenS := len(s)
	lenWords := len(words)
	lenWord := len(words[0])
	if lenS < lenWords*lenWord {
		return []int{}
	}
	var ans []int

	countWords := 0
	wordMp := make(map[string]int)
	for _, word := range words {
		if _, ok := wordMp[word]; !ok {
			countWords++
		}
		wordMp[word]++
	}

	var curLen int // 当前子串长度
	var wordCount int
	left, preRight := 0, 0

	wordFreq := make(map[string]int)

	// 以 s[right, right+lenWord-1] 为当前子串,以 left 为子串起点，
	// 找到长度 curLen 为 lenWords*lenWord 的 left、right 下标。不一定为串联子串，因为可能找到了几个重复的单词
	for right := preRight + lenWord; right <= lenS; right = preRight + lenWord {
		curStr := s[preRight:right]

		// 找到了串联子串的一部分，即 words 中某个单词
		if _, ok := wordMp[curStr]; ok {
			curLen += lenWord
			wordFreq[curStr]++
			if wordFreq[curStr] == wordMp[curStr] {
				wordCount++
			}

			// 子串长度 == lenWords*lenWord，且子串中出现的 words 中单词的个数 == countWords，则找到一个串联子串
			// 无论如何，重置状态，以便下一个子串从 left + 1 开始搜索
			if curLen == lenWords*lenWord {
				if wordCount == countWords {
					ans = append(ans, left)
				}

				// 滑动窗口右移，注意重置状态
				left = left + 1
				wordFreq = make(map[string]int)
				wordCount = 0
				preRight = left
				curLen = 0
			} else {
				// 遍历下一个单词
				preRight = right
			}

			// 在找到串联子串之前断开了，则重置 curLen，更新 left、right，重新开始寻找。
		} else {
			curLen = 0
			left++
			wordFreq = make(map[string]int)
			wordCount = 0
			preRight = left
		}
	}
	return ans
}

// Test Case 1: "barfoothefoobarman" ["foo","bar"]		Output: [0,9]
// Test Case 2: "wordgoodgoodgoodbestword" ["word","good","best","word"]		Output: []
// Test Case 3: "barfoofoobarthefoobarman" ["bar","foo","the"]		Output: [6,9,12]
func main() {
	var s string
	fmt.Println("Input String:")
	fmt.Scan(&s)
	words := pkg.CreateSlice[string]()
	fmt.Println(findSubstringBetter(s, words))
}
