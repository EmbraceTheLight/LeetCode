// 17. 电话号码的字母组合
/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

示例 2：
输入：digits = "2"
输出：["a","b","c"]

提示：
1 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。
*/
package main

import "fmt"

func letterCombinations(digits string) []string {
	mp := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	var ans []string
	var dfs func(step int, digits string, str string)
	dfs = func(step int, digits string, str string) {
		if step == len(digits) {
			ans = append(ans, str)
			return
		}
		for _, ch := range mp[string(digits[step])] {
			tmp := str + ch
			dfs(step+1, digits, tmp)
		}
	}

	dfs(0, digits, "")
	return ans
}

// Test Case1: "23"	 Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
// Test Case2: "2"	Output: ["a","b","c"]
func main() {
	var digits string
	fmt.Scan(&digits)
	fmt.Println(letterCombinations(digits))
}
