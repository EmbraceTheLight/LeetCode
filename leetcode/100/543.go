/*
543. 二叉树的直径
给你一棵二叉树的根节点，返回该树的 直径 。
二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。
两节点之间路径的 长度 由它们之间边数表示。
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func dfs543(node *pkg.TreeNode, length *int) int {
	if node == nil {
		return 0
	}
	llen := dfs543(node.Left, length)
	rlen := dfs543(node.Right, length)
	maxlen := llen + rlen + 1
	*length = max(*length, maxlen)
	return max(llen, rlen) + 1

}
func diameterOfBinaryTree(root *pkg.TreeNode) int {
	var l int
	dfs543(root, &l)
	return l - 1
}
func main() {
	root := pkg.CreateTree()
	fmt.Println(diameterOfBinaryTree(root))
}
