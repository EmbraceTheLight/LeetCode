/*
112. 路径总和
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum。
判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum。如果存在，返回 true ；否则，返回 false 。
叶子节点 是指没有子节点的节点。
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func dfs112(node *pkg.TreeNode, sum, target int) bool {
	if node == nil {
		return false
	}
	sum += node.Val

	if node.Left == nil && node.Right == nil && sum == target {
		return true
	}
	return dfs112(node.Left, sum, target) || dfs112(node.Right, sum, target)
}
func hasPathSum(root *pkg.TreeNode, targetSum int) bool {
	return dfs112(root, 0, targetSum)
}
func main() {
	println("Input sum:")
	var tmp int
	fmt.Scan(&tmp)
	root := pkg.CreateTree()

	println(hasPathSum(root, tmp))
}
