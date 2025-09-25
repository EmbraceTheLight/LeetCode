package main

import "fmt"

func dfs22R(cntLeft, cntRight, n int, bytes []byte, ans *[]string) {
	if cntLeft+cntRight == n*2 {
		*ans = append(*ans, string(bytes))
		return
	}

	if cntLeft < n { //还能添加左括号
		bytes[cntLeft+cntRight] = '('
		dfs22R(cntLeft+1, cntRight, n, bytes, ans)
	}
	if cntLeft > cntRight {
		bytes[cntLeft+cntRight] = ')'
		dfs22R(cntLeft, cntRight+1, n, bytes, ans)
	}
}
func generateParenthesis22R(n int) []string {
	var ans = make([]string, 0)
	bytes := make([]byte, n*2)
	dfs22R(0, 0, n, bytes, &ans)
	return ans
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(generateParenthesis22R(n))
}
