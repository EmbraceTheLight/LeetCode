/*
139.单词拆分
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

示例 1：
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。

示例 2：
输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
注意，你可以重复使用字典中的单词。

示例 3：
输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false

提示：
1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s 和 wordDict[i] 仅由小写英文字母组成
wordDict 中的所有字符串 互不相同
*/
package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	n := len(s)

	//dp[i]表示s中以s[i]结尾的子串是否可以由wordDict中的单词拼接
	dp := make([]bool, n+1)
	dp[0] = true
	mp := make(map[string]bool)
	for _, v := range wordDict {
		mp[v] = true
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && mp[s[j:i]] == true {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
func main() {
	var s string
	fmt.Println("Input the string s:")
	fmt.Scanf("%s\n", &s)
	var wordDict []string
	fmt.Println("Input -1 to quit.")
	var tmp string
	for {
		fmt.Scanf("%s\n", &tmp)
		if tmp == "-1" {
			break
		}
		wordDict = append(wordDict, tmp)
	}
	fmt.Println(wordBreak(s, wordDict))
}
