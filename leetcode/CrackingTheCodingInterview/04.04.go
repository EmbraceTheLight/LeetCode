// 04.04. 检查平衡性
/*
实现一个函数，检查二叉树是否平衡。在这个问题中，平衡树的定义如下：任意一个节点，其两棵子树的高度差不超过 1。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	mp := make(map[*TreeNode]int)
	calcTreeHeight0404(root, mp)
	return dfs0404(root, mp)
}
func dfs0404(root *TreeNode, heightMap map[*TreeNode]int) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right != nil && heightMap[root.Right] > 1 {
		return false
	}
	if root.Left != nil && root.Right == nil && heightMap[root.Left] > 1 {
		return false
	}
	if heightMap[root.Left]-heightMap[root.Right] > 1 || heightMap[root.Right]-heightMap[root.Left] > 1 {
		return false
	}
	lRes := dfs0404(root.Left, heightMap)
	if lRes == false {
		return false
	}
	return dfs0404(root.Right, heightMap)
}
func calcTreeHeight0404(root *TreeNode, heightMap map[*TreeNode]int) int {
	if root.Left == nil && root.Right == nil {
		heightMap[root] = 1
		return heightMap[root]
	}
	var lHeight, rHeight int
	if root.Left != nil {
		lHeight = calcTreeHeight0404(root.Left, heightMap)
	}
	if root.Right != nil {
		rHeight = calcTreeHeight0404(root.Right, heightMap)
	}
	heightMap[root] = max(lHeight, rHeight) + 1
	return heightMap[root]
}

//	示例 1：
//
// 给定二叉树 [3,9,20,null,null,15,7]
//
//	  3
//	 / \
//	9  20
//	  /  \
//	 15   7
//
// 返回 true 。
// 示例 2：
// 给定二叉树 [1,2,2,3,3,null,null,4,4]
//
//	     1
//	    / \
//	   2   2
//	  / \
//	 3   3
//	/ \
//
// 4   4
// 返回 false 。
func main() {
	fmt.Println(isBalanced(CreateTree()))
}
