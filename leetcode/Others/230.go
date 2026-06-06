/*
230. 二叉搜索树中第K小的元素
给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func dfs(node *pkg.TreeNode, k *int, ans *int) {
	if node == nil || *k == 0 {
		return
	}
	dfs(node.Left, k, ans)
	*k--
	if *k == 0 {
		*ans = node.Val
	}
	dfs(node.Right, k, ans)
}
func kthSmallest(root *pkg.TreeNode, k int) int {
	var ans int
	dfs(root, &k, &ans)
	return ans
}
func main() {
	println("Input cnt:")
	var k int
	fmt.Scan(&k)

	root := pkg.CreateTree()
	println(kthSmallest(root, k))
}
