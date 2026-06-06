package main

import (
	"fmt"
	"strings"
)

/*
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

示例 2：
输入：digits = ""
输出：[]

示例 3：
输入：digits = "2"
输出：["a","b","c"]

提示：
0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。
*/

func dfs17(mp map[byte][]string, digits string, now int, res *[]string, tmp string) {
	if now == len(digits) {
		t := strings.Clone(tmp)
		*res = append(*res, t)
		return
	}
	for _, v := range mp[digits[now]] {
		tmp += v
		dfs17(mp, digits, now+1, res, tmp)
		tmp = tmp[:len(tmp)-1]
	}

}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mp := make(map[byte][]string)
	mp['2'] = []string{"a", "b", "c"}
	mp['3'] = []string{"d", "e", "decode"}
	mp['4'] = []string{"g", "h", "i"}
	mp['5'] = []string{"j", "k", "l"}
	mp['6'] = []string{"m", "n", "o"}
	mp['7'] = []string{"p", "q", "r", "s"}
	mp['8'] = []string{"t", "u", "v"}
	mp['9'] = []string{"w", "x", "y", "z"}
	var res = make([]string, 0)
	dfs17(mp, digits, 0, &res, "")
	return res
}
func main() {
	fmt.Println("input a string:")
	var s string
	fmt.Scan(&s)
	fmt.Println(letterCombinations(s))
}
