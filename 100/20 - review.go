package main

import "fmt"

func isValidR(s string) bool {
	bracketMapping := make(map[byte]byte)
	bracketMapping['{'] = '}'
	bracketMapping['['] = ']'
	bracketMapping['('] = ')'
	stack20 := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack20 = append(stack20, s[i])
		} else if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(stack20) == 0 { // 当前是右括号，但是栈中没有与之匹配的左括号
				return false
			}
			top := stack20[len(stack20)-1]
			stack20 = stack20[:len(stack20)-1]
			if bracketMapping[top] != s[i] {
				return false
			}
		}
	}
	return len(stack20) == 0
}
func main() {
	fmt.Println("input brackets:")
	var s string
	fmt.Scan(&s)
	fmt.Println(isValidR(s))
}
