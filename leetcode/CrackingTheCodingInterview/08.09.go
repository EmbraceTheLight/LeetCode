// 08.09. 括号
/*
括号。设计一种算法，打印n对括号的所有合法的（例如，开闭一一对应）组合。
说明：解集不能包含重复的子集。
*/
package main

import "fmt"

func generateParenthesis(n int) []string {
	var ans []string
	dfs0809(n, n, 0, "", &ans)
	return ans
}

// n: 总括号对数
// leftPar 剩余的未添加的左括号数量
// curPar 当前未匹配的左括号数量
func dfs0809(n int, leftPar, curLeftPar int, cur string, ans *[]string) {
	if len(cur) == n*2 {
		*ans = append(*ans, cur)
		return
	}
	if leftPar == 0 { // cur 中的左括号已全部补齐, 需要添加右括号
		dfs0809(n, leftPar, curLeftPar-1, cur+")", ans)
	} else if curLeftPar == 0 { // cur 中的括号已全部配对, 需要添加左括号
		dfs0809(n, leftPar-1, curLeftPar+1, cur+"(", ans)
	} else { // 普通情况: 既可以添加左括号, 又可以添加右括号
		dfs0809(n, leftPar, curLeftPar-1, cur+")", ans)
		dfs0809(n, leftPar-1, curLeftPar+1, cur+"(", ans)
	}
}

// 示例：
// 输入：n = 3
// 输出：
// [
//
//	"((()))",
//	"(()())",
//	"(())()",
//	"()(())",
//	"()()()"
//
// ]
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(generateParenthesis(n))
}
