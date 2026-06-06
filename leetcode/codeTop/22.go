// 22. 括号生成
/*
数字 n 代表生成括号的对数，请你设计一个函数，
用于能够生成所有可能的并且 有效的 括号组合。

提示：
1 <= n <= 8
*/
package main

import "fmt"

func generateParenthesis(n int) []string {
	var ans []string
	dfs22(0, n, "", &ans)
	return ans
}
func dfs22(leftBracket int, pairLeft int, cur string, ans *[]string) {
	if pairLeft == 0 && leftBracket == 0 {
		*ans = append(*ans, cur)
		return
	}
	if pairLeft > 0 {
		dfs22(leftBracket+1, pairLeft-1, cur+"(", ans)
	}
	if leftBracket > 0 {
		dfs22(leftBracket-1, pairLeft, cur+")", ans)
	}
}

// 示例 1：
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
//
// 示例 2：
// 输入：n = 1
// 输出：["()"]
func main() {
	fmt.Println("Input n:")
	var n int
	fmt.Scan(&n)
	fmt.Println(generateParenthesis(n))
}
