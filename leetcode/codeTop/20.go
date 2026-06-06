// 20. 有效的括号
/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

提示：
1 <= s.length <= 10^4
s 仅由括号 '()[]{}' 组成
*/
package main

import "fmt"

func isValid(s string) bool {
	stk := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stk = append(stk, s[i])
		} else {
			if len(stk) == 0 {
				return false
			}
			top := stk[len(stk)-1]
			if s[i] == ')' {
				if top != '(' {
					return false
				}
			} else if s[i] == ']' {
				if top != '[' {
					return false
				}
			} else if s[i] == '}' {
				if top != '{' {
					return false
				}
			}
			stk = stk[:len(stk)-1]
		}
	}
	return len(stk) == 0
}

// 示例 1：
// 输入：s = "()"
// 输出：true
//
// 示例 2：
// 输入：s = "()[]{}"
// 输出：true
//
// 示例 3：
// 输入：s = "(]"
// 输出：false
//
// 示例 4：
// 输入：s = "([])"
// 输出：true
//
// 示例 5：
// 输入：s = "([)]"
// 输出：false
func main() {
	var s string
	fmt.Println("Input s:")
	fmt.Scan(&s)
	fmt.Println(isValid(s))
}
