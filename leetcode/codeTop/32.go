// 32. 最长有效括号
/*
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号 子串 的长度。
左右括号匹配，即每个左括号都有对应的右括号将其闭合的字符串是格式正确的，比如 "(()())"。

提示：
0 <= s.length <= 3 * 10^4
s[i] 为 '(' 或 ')'
*/
package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	var ans int
	left, right := 0, 0 // 当前连续合法区间的左括号数量, 右括号数量
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right { // 左括号数量 == 右括号数量, 此时可能是该连续区间最大合法长度, 更新 ans
			ans = max(ans, 2*right)
		} else if right > left { // 右括号数量大于左括号数量, 当前连续合法区间断开, 应当重置左括号数量, 有括号数量
			left, right = 0, 0
		}
	}
	left, right = 0, 0

	// 重要: 反向遍历
	// 上面的循环会遗漏掉一种情况: 左括号数量始终多余有括号数量, 如"(()".
	// 为了解决这种情况, 再反向遍历一遍, 这样就能覆盖所有情况
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right { // 左括号数量 == 右括号数量, 此时可能是该连续区间最大合法长度, 更新 ans
			ans = max(ans, 2*left)
		} else if left > right { // 左括号数量大于右括号数量, 当前连续合法区间断开, 应当重置左括号数量, 有括号数量
			left, right = 0, 0
		}
	}
	return ans
}

// 示例 1：
// 输入：s = "(()"
// 输出：
// 解释：最长有效括号子串是 "()"
//
// 示例 2：
// 输入：s = ")()())"
// 输出：4
// 解释：最长有效括号子串是 "()()"
//
// 示例 3：
// 输入：s = ""
// 输出：0
func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(longestValidParentheses(s))
}
