// 08.14. 布尔运算
/*
给定一个布尔表达式和一个期望的布尔结果 result，布尔表达式由 0 (false)、1 (true)、& (AND)、 | (OR) 和 ^ (XOR) 符号组成。实现一个函数，算出有几种可使该表达式得出 result 值的括号方法。

提示：
运算符的数量不超过 19 个
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func countEval(s string, result int) int {
	numOfBool := len(s)/2 + 1
	dp := make([][][2]int, numOfBool) // dp[intervalLength][j][k] 表示第 intervalLength 个操作数到第 j 个操作数之间，结果为 k 的方案数. k == 0: false, k == 1: true
	for i := 0; i < numOfBool; i++ {
		dp[i] = make([][2]int, numOfBool)
		if getBool(s, i) == true {
			dp[i][i][1] = 1
			dp[i][i][0] = 0
		} else {
			dp[i][i][0] = 1
			dp[i][i][1] = 0
		}
	}

	// 枚举运算符
	for intervalLength := 2; intervalLength <= numOfBool; intervalLength++ { // intervalLength 表示操作数表达式区间长度, 从两个操作数开始
		for startIdx := 0; startIdx+intervalLength-1 < numOfBool; startIdx++ { // startIdx 起始操作数位置
			endIdx := startIdx + intervalLength - 1 // endIdx: 末尾操作数位置
			for k := startIdx; k < endIdx; k++ {
				op := getOperator(s, k)
				if op == '&' {
					dp[startIdx][endIdx][0] += dp[startIdx][k][0] * dp[k+1][endIdx][1]
					dp[startIdx][endIdx][0] += dp[startIdx][k][1] * dp[k+1][endIdx][0]
					dp[startIdx][endIdx][0] += dp[startIdx][k][0] * dp[k+1][endIdx][0]

					dp[startIdx][endIdx][1] += dp[startIdx][k][1] * dp[k+1][endIdx][1]
				} else if op == '|' {
					dp[startIdx][endIdx][1] += dp[startIdx][k][1] * dp[k+1][endIdx][0]
					dp[startIdx][endIdx][1] += dp[startIdx][k][0] * dp[k+1][endIdx][1]
					dp[startIdx][endIdx][1] += dp[startIdx][k][1] * dp[k+1][endIdx][1]

					dp[startIdx][endIdx][0] += dp[startIdx][k][0] * dp[k+1][endIdx][0]
				} else if op == '^' {
					dp[startIdx][endIdx][0] += dp[startIdx][k][0] * dp[k+1][endIdx][0]
					dp[startIdx][endIdx][0] += dp[startIdx][k][1] * dp[k+1][endIdx][1]

					dp[startIdx][endIdx][1] += dp[startIdx][k][0] * dp[k+1][endIdx][1]
					dp[startIdx][endIdx][1] += dp[startIdx][k][1] * dp[k+1][endIdx][0]
				}
			}
		}
	}

	return dp[0][numOfBool-1][result]
}

// getBool 获取第 idx 个操作数, idx 从 0 开始
func getBool(s string, idx int) bool {
	return s[idx*2] == '1'
}

// getOperator 获取第 idx 个操作符, idx 从 0 开始
func getOperator(s string, idx int) byte {
	return s[idx*2+1]
}

// 示例 1：
// 输入：s = "1^0|0|1", result = 0
// 输出：2
// 解释：两种可能的括号方法是
// 1^(0|(0|1))
// 1^((0|0)|1)
//
// 示例 2：
// 输入：s = "0&0&0&1^1|0", result = 1
// 输出：10
func main() {
	var str string
	var result int
	fmt.Println("Input result:")
	fmt.Scan(&result)
	fmt.Println("Input eval str:")
	str = CreateString()
	fmt.Println(countEval(str, result))
}
