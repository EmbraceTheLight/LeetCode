/*
226. 翻转二叉树
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
*/
package main

import "lc/100/pkg"

//	func dfs(node *pkg.TreeNode)bool{
//		if node.Left==nil || node.Right==nil{
//			return false
//		}
//		var pre = node
//		left:=invertTree(node.Left)
//		right:=invertTree(node.Right)
//	}
func invertTree(root *pkg.TreeNode) *pkg.TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	var pre = root
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	var tmp *pkg.TreeNode
	tmp = left
	pre.Left = right
	pre.Right = tmp
	return root
}
func main() {
	root := pkg.CreateTree()
	pkg.PrintTreeByLevel(invertTree(root))
}
