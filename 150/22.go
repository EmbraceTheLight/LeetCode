// 22. 括号生成
/*
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

func generateParenthesis(n int) []string {
	var ans []string

	// totalLeft:  还可以添加的左括号数量	leftCount: 尚未被匹配的左括号数量	cur: 当前括号匹配字符串
	var dfs func(leftCouldBeAdded, notMatchedCount int, cur string)
	dfs = func(leftCouldBeAdded, notMatchedCount int, cur string) {
		if leftCouldBeAdded == 0 {
			// 没有可以插入的左括号了, 但此时还有一些左括号尚未匹配, 需要添加一定数量的右括号, 保证括号匹配
			for notMatchedCount > 0 {
				cur += ")"
				notMatchedCount--
			}
			ans = append(ans, cur)
			return
		}

		// 还有可插入的左括号, 插入左括号
		if leftCouldBeAdded > 0 {
			dfs(leftCouldBeAdded-1, notMatchedCount+1, cur+"(")
		}

		// 有可匹配的左括号, 插入右括号, 匹配
		if notMatchedCount > 0 {
			dfs(leftCouldBeAdded, notMatchedCount-1, cur+")")
		}
	}
	dfs(n, 0, "")
	return ans
}

// Test Case1: n = 3	Output: ["((()))","(()())","(())()","()(())","()()()"]
// Test Case2: n = 1	Output: ["()"]
func main() {
	var n int
	fmt.Println("Input n")
	fmt.Scan(&n)
	fmt.Println(generateParenthesis(n))
}
