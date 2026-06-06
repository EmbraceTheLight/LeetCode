// 20. 有效的括号
/* 题目描述见 20.go */
package main

import "fmt"

func isValidR(s string) bool {
	var stack []byte
	mp := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != mp[s[i]] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// Test Case1:	"()"		Output: true
// Test Case2:	"()[]{}"	Output: true
// Test Case3:	"(]"		Output: false
// Test Case4:	"{[]}"		Output: true
// Test Case5:	"([)]"		Output: false
func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(isValidR(s))
}
