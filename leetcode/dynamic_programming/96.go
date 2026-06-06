/*
96. 不同的二叉搜索树
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。

示例 1：
输入：n = 3
输出：5

示例 2：
输入：n = 1
输出：1

提示：
1 <= n <= 19
*/
package main

import "fmt"

func numTrees(n int) int {
	dp := make([]int, n+1) //dp[i]：从 1 到 i 互不相同的 二叉搜索树 的种数
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum += dp[j] * dp[i-1-j] //左子树有j个节点，一共有i个节点，除去根节点(i-1),右子树就只剩 (i-1)-j 个节点。将左右子树种数相乘，再累加，就得到了答案
		}
		dp[i] = sum
	}
	return dp[n]
}
func main() {
	var n int
	fmt.Println("Input n:")
	fmt.Scan(&n)
	fmt.Println(numTrees(n))
}
