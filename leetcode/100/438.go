/*
找到字符串中所有字母异位词

给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。

示例 2:
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。

提示:

1 <= s.length, p.length <= 3 * 10^4
s 和 p 仅包含小写字母
*/
package main

import "fmt"

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	pLength := len(p)
	var ans = make([]int, 0)
	//初始化map以便于进行比对
	mp := make(map[uint8]int)
	for i := 0; i < pLength; i++ {
		mp[p[i]]++
	}

	//初始化窗口,长度为要对比的字串p长度
	var count = len(mp) //count统计s窗口中有多少种字符未与p中字符完全比对上，count为0则表示找到一个异位词
	for i := 0; i < pLength; i++ {
		if _, ok := mp[s[i]]; ok {
			mp[s[i]]--
			if mp[s[i]] == 0 {
				count--
			}
		}
	}

	//滑动窗口
	sLength := len(s)
	for start, end := 0, pLength-1; end < sLength-1; {
		if count == 0 {
			ans = append(ans, start)
		}

		//移动左边界
		if _, ok := mp[s[start]]; ok {
			if mp[s[start]] == 0 {
				count++
			}
			mp[s[start]]++
		}
		start++

		//移动右边界
		end++
		if _, ok := mp[s[end]]; ok {
			mp[s[end]]--
			if mp[s[end]] == 0 {
				count--
			}
		}
	}
	if count == 0 {
		ans = append(ans, sLength-pLength)
	}
	return ans
}

func main() {
	var s, p string
	fmt.Println("Input string s:")
	fmt.Scan(&s)
	fmt.Println("Input string p:")
	fmt.Scan(&p)
	fmt.Println(findAnagrams(s, p))
}
