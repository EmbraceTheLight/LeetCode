// 49. 字母异位词分组
/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

示例 1:
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
解释：
在 strs 中没有字符串可以通过重新排列来形成 "bat"。
字符串 "nat" 和 "tan" 是字母异位词，因为它们可以重新排列以形成彼此。
字符串 "ate" ，"eat" 和 "tea" 是字母异位词，因为它们可以重新排列以形成彼此。

示例 2:
输入: strs = [""]
输出: [[""]]

示例 3:
输入: strs = ["a"]
输出: [["a"]]


提示：
1 <= strs.length <= 10^4
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母
*/
package main

import (
	"fmt"
	"lc/pkg"
	"sort"
)

// 暴力法。超时
func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	var mps [][26]int
	for i := 0; i < n; i++ {
		var mp [26]int
		for j := 0; j < len(strs[i]); j++ {
			mp[strs[i][j]-'a']++
		}
		mps = append(mps, mp)
	}
	var ans [][]string
	visited := make(map[int]bool)
	for i := 0; i < n; i++ {
		var group []string
		if visited[i] {
			continue
		}
		group = append(group, strs[i])
		mp1 := mps[i]
		visited[i] = true
		for j := 0; j < n; j++ {
			mp2 := mps[j]
			if visited[j] {
				continue
			}

			k := 0
			for ; k < 26; k++ {
				if mp1[k] != mp2[k] {
					break
				}
			}
			if k == 26 {
				group = append(group, strs[j])
				visited[j] = true
			}
		}
		ans = append(ans, group)
	}
	return ans
}

func groupAnagramsBetter(strs []string) [][]string {
	n := len(strs)
	type pair struct {
		s        string
		letterMp [26]int
	}

	var mps []*pair
	for i := 0; i < n; i++ {
		mp := &pair{
			s:        strs[i],
			letterMp: [26]int{},
		}
		for j := 0; j < len(strs[i]); j++ {
			mp.letterMp[strs[i][j]-'a']++
		}
		mps = append(mps, mp)
	}

	sort.Slice(mps, func(i, j int) bool {
		for k := 0; k < 26; k++ {
			if mps[i].letterMp[k] == mps[j].letterMp[k] {
				continue
			}
			if mps[i].letterMp[k] > mps[j].letterMp[k] {
				return true
			} else {
				return false
			}
		}
		return false
	})

	var ans [][]string
	left, right := 0, 0
	for left < n {
		var group []string
		mp1 := mps[left].letterMp
		for ; right < n; right++ {
			mp2 := mps[right].letterMp
			k := 0
			for ; k < 26; k++ {
				if mp1[k] != mp2[k] {
					break
				}
			}
			if k == 26 {
				group = append(group, mps[right].s)
			} else {
				left = right
				break
			}
		}
		left = right
		ans = append(ans, group)
	}

	return ans
}

// 优化：使用 map，key 为数组，值为字符串数组 key --> 字符数量完全相等的字符串数组，这样就不用手动排序、一个个比较了。
// 时间复杂度略高
func groupAnagramsBest(strs []string) [][]string {
	n := len(strs)
	mp := make(map[[26]int][]string)
	var ans [][]string
	for i := 0; i < n; i++ {
		var mp1 [26]int
		for j := 0; j < len(strs[i]); j++ {
			mp1[strs[i][j]-'a']++
		}
		mp[mp1] = append(mp[mp1], strs[i])
	}
	for _, slices := range mp {
		ans = append(ans, slices)
	}
	return ans
}

// Test Case1:	["eat", "tea", "tan", "ate", "nat", "bat"]	Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Test Case2:	[""]	Output: [[""]]
// Test Case3:	["a"]	Output: [["a"]]
func main() {
	strs := pkg.CreateSlice[string]()
	fmt.Println(groupAnagramsBest(strs))
}
