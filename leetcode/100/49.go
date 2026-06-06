/*
49. 字母异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

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
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	strings := make(map[string][]string)
	for _, a := range strs {
		tmpa := []byte(a)

		sort.Slice(tmpa, func(i, j int) bool {
			return tmpa[i] < tmpa[j]
		})
		strings[string(tmpa)] = append(strings[string(tmpa)], a)
	}
	for _, sli := range strings {
		result = append(result, sli)
	}
	fmt.Println(result)
	return result
}

func main() {
	strs := make([]string, 0)
	fmt.Println("Input Value:")
	var tmp string

	for {
		fmt.Scan(&tmp)
		if tmp == "-1" {
			break
		}
		strs = append(strs, tmp)
	}
	fmt.Printf("%+v", strs)
	a := groupAnagrams(strs)
	fmt.Println(a)
}
