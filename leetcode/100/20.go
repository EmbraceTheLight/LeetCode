/*
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
1 <= s.length <= 10^4
s 仅由括号 '()[]{}' 组成
*/
package main

import "fmt"

func isValid(s string) bool {
	var brackets []byte
	l := len(s)
	for i := 0; i < l; i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			brackets = append(brackets, s[i])
		} else {
			if len(brackets) == 0 {
				return false
			}
			var cmp = brackets[len(brackets)-1]
			brackets = brackets[:len(brackets)-1]
			if s[i] == '}' {
				if cmp != '{' {
					return false
				}
			} else if s[i] == ']' {
				if cmp != '[' {
					return false
				}
			} else if s[i] == ')' {
				if cmp != '(' {
					return false
				}
			}
		}
	}
	if len(brackets) > 0 {
		return false
	}
	return true

}
func main() {
	var s string
	fmt.Println("Input a string:")
	fmt.Scan(&s)
	fmt.Println(isValid(s))
}
