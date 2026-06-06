/*
32.最长有效括号
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号
子串
的长度。

示例 1：
输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"

示例 2：
输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"

示例 3：
输入：s = ""
输出：0

提示：
0 <= s.length <= 3 * 10^4
s[i] 为 '(' 或 ')'
*/
package main

import (
	"fmt"
)

/*暴力解法*/
//func findSubstrLen(s string, start int) (int, int) {
//	var cnt = 0 //统计左括号个数
//	var ret = 0
//	var length = 0
//	i := start
//	idx := i
//	for ; i < len(s); i++ {
//		if s[i] == '(' {
//			cnt++
//		} else if s[i] == ')' {
//			cnt--
//			if cnt < 0 {
//				return ret, i
//			}
//			length += 2
//			if cnt == 0 {
//				ret += length
//				idx = i
//				length = 0
//			}
//		}
//	}
//	return ret, idx
//}
//func longestValidParentheses(s string) int {
//	var ans int
//	l := len(s)
//	for i := 0; i < l; {
//		if s[i] == '(' {
//
//		} else if s[i] == ')' {
//			i++
//		}
//	}
//	return ans
//}

func longestValidParentheses(s string) int {
	var ans int
	l := len(s)
	dp := make([]int, l) //以dp[i]结尾的最长有效括号子串
	for i := 1; i < l; i++ {
		if s[i] == ')' { //匹配到 ')'
			if s[i-1] == '(' { //')'前面是'('
				dp[i] = 2
			} else { //')'前面是')',继续向前找'('
				idx := i - dp[i-1] - 1 //寻找与当前')'匹配的'('
				if idx >= 0 && s[idx] == '(' {
					dp[i] = dp[i-1] + 2
				}
				idx--
			}
			if i-dp[i] >= 0 { //到这里,已经找到了匹配的括号子串,如((())),还要判断前面有没有完整的括号子串，如()((())),需要加上前面的()的长度，该子串长度才算完整
				dp[i] += dp[i-dp[i]]
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

// 使用Stack解决，难点在于要先给栈压入未与左括号匹配的右括号的下标，这样遇到左右括号匹配时，遇到栈底元素就不在继续匹配
// 遇到新的没有与之匹配的右括号时，更新栈底元素
// 栈只压入左括号下标，栈底元素除外
func longestValidParenthesesUseStack(s string) int {
	var ans int
	var stk []int
	stk = append(stk, -1) //一开始没有右括号，以-1作为虚拟下标
	for i := 0; i < len(s); i++ {
		if s[i] == ')' { //右括号
			stk = stk[:len(stk)-1] //弹出栈顶元素,注意一开始就弹出栈顶元素
			if len(stk) == 0 {     //栈为空，更新栈底元素
				stk = append(stk, i)
			} else {
				ans = max(ans, i-stk[len(stk)-1]) //栈底元素为-1，若s为(),由于一开始就弹出了栈顶元素，即左括号下标，再减时就应该为1-(-1)=2
			}
		} else { //左括号,直接压栈
			stk = append(stk, i)
		}
	}
	return ans
}
func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(longestValidParentheses(s))
}
