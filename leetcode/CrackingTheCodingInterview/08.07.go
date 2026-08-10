// 08.07. 无重复字符串的排列组合
/*
无重复字符串的排列组合。编写一种方法，计算某字符串的所有排列组合，字符串每个字符均不相同。

提示：
字符都是英文字母。
字符串长度在[1, 9]之间。
*/
package main

import "fmt"

func permutation(S string) []string {
	var ans []string
	dfs0807(S, "", [9]bool{}, &ans)
	return ans
}
func dfs0807(s string, tmp string, seen [9]bool, ans *[]string) {
	if len(tmp) == len(s) {
		*ans = append(*ans, tmp)
		return
	}
	for i := 0; i < len(s); i++ {
		if seen[i] == true {
			continue
		}

		seen[i] = true
		tmpS := tmp + string(s[i])
		dfs0807(s, tmpS, seen, ans)
		seen[i] = false
	}
}

// 示例 1：
// 输入：S = "qwe"
// 输出：["qwe", "qew", "wqe", "weq", "ewq", "eqw"]
//
// 示例 2：
// 输入：S = "ab"
// 输出：["ab", "ba"]
func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(permutation(s))
}
