/*
22. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

示例 2：
输入：n = 1
输出：["()"]

提示：
1 <= n <= 8
*/
package main

import "fmt"

// cntLeft记录未匹配的左括号数量,target表示生成的括号的对数,s表示当前生成的括号字符串
func dfs22(cntLeft, target int, s string, ans *[]string) {
	if len(s) == target*2 {
		*ans = append(*ans, s)
		return
	}

	var tmp = s

	s += "("
	if cntLeft+1 <= target-(len(s)-cntLeft)/2 { //此时的左括号数量小于等于最大可以添加的左括号数量，还可以继续添加左括号
		dfs22(cntLeft+1, target, s, ans)
	}
	if cntLeft > 0 {
		tmp += ")"
		dfs22(cntLeft-1, target, tmp, ans)
	}
}
func generateParenthesis(n int) []string {
	var ans = make([]string, 0)
	dfs22(0, n, "", &ans)
	return ans
}
func main() {
	var n int
	fmt.Println("Input int:")
	fmt.Scan(&n)
	fmt.Println(generateParenthesis(n))
}
