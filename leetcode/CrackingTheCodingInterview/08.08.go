// 08.08. 有重复字符串的排列组合
/*
有重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合。

提示:
字符都是英文字母。
字符串长度在[1, 9]之间。
*/
package main

import (
	"fmt"
)

func permutation0808(S string) []string {
	var ans []string
	resMap := make(map[string]bool)
	dfs0808(S, "", [9]bool{}, resMap, &ans)
	return ans
}

func dfs0808(s string, tmp string, seen [9]bool, resMap map[string]bool, ans *[]string) {
	if len(tmp) == len(s) {
		if resMap[tmp] == true {
			return
		}
		resMap[tmp] = true
		*ans = append(*ans, tmp)
		return
	}
	for i := 0; i < len(s); i++ {
		if seen[i] == true {
			continue
		}

		seen[i] = true
		tmpS := tmp + string(s[i])
		dfs0808(s, tmpS, seen, resMap, ans)
		seen[i] = false
	}
}

func permutation0808_2(S string) []string {
	var ans []string
	dfs0808_2(S, "", [9]bool{}, &ans)
	return ans
}
func dfs0808_2(s string, tmp string, seen [9]bool, ans *[]string) {
	if len(tmp) == len(s) {
		*ans = append(*ans, tmp)
		return
	}
	visit := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if seen[i] == true || visit[s[i]] == true { // 本轮已经遍历过相同的字符了
			continue
		}

		seen[i] = true
		visit[s[i]] = true
		tmpS := tmp + string(s[i])
		dfs0808_2(s, tmpS, seen, ans)
		seen[i] = false
	}
}

// 示例 1：
// 输入：S = "qqe"
// 输出：["eqq","qeq","qqe"]
//
// 示例 2：
// 输入：S = "ab"
// 输出：["ab", "ba"]
func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(permutation0808_2(s))
}
