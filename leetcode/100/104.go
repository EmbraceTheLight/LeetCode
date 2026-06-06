/*
104. 二叉树的最大深度
给定一个二叉树 root ，返回其最大深度。
二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。
*/
package main

import (
	"lc/100/pkg"
)

func maxDepth(root *pkg.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func main() {
	root := pkg.CreateTree()
	println(maxDepth(root))
}
