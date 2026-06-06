/*
139. 单词拆分
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

// s[0,i-1]是否能被拆分可以由s[0,j-1]和s[j,i-1]来决定。前半部分是否合法可以通过dp数组来确定(dp[j-1]),只需判断后半部分是否在字典中即可。
// 最后一步一步推导得出结果
func wordBreak(s string, wordDict []string) bool {
	mp := make(map[string]bool)
	for _, str := range wordDict {
		mp[str] = true
	}
	n := len(s)
	dp := make([]bool, n+1) //dp[i]表示s[0,i-1]是否可以拆分
	dp[0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j-1] && mp[s[j:i]] {
				dp[j+1] = true
				break
			}
		}
	}
	return dp[n]
}
func main() {
	//var s string
	//fmt.Println("Input target string:")
	//fmt.Scan(&s)
	//fmt.Println("Input dict words(space to split):")
	//read := bufio.Reader{}
	//line, _ := read.ReadString('\n')
	//words := strings.Split(line, " ")
	//
	//fmt.Println(wordBreak(s, words))
}
