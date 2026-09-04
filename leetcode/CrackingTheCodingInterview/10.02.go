// 10.02. 变位词组
/*
编写一种方法，对字符串数组进行排序，将所有变位词组合在一起。变位词是指字母相同，但排列不同的字符串。
注意：本题相对原题稍作修改

说明：
所有输入均为小写字母。
不考虑答案输出的顺序。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)
	var ans [][]string
	for _, str := range strs {
		tmp := [26]int{}
		for i := range str {
			tmp[str[i]-'a']++
		}
		mp[tmp] = append(mp[tmp], str)
	}
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

// 示例：
// 输入：["eat", "tea", "tan", "ate", "nat", "bat"],
// 输出：
// [
//
//	["ate","eat","tea"],
//	["nat","tan"],
//	["bat"]
//
// ]
func main() {
	strs := CreateSlice[string]()
	fmt.Println(groupAnagrams(strs))
}
