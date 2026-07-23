// 04.08. 首个共同祖先
/*
设计并实现一个算法，找出二叉树中某两个节点的第一个共同祖先。不得将其他的节点存储在另外的数据结构中。注意：这不一定是二叉搜索树。
例如，给定如下二叉树: root = [3,5,1,6,2,0,8,null,null,7,4]
    3
   / \
  5   1
 / \ / \
6  2 0  8
  / \
 7   4

提示：
所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	return dfs0408(root, p, q)
}

// dfs0408 递归搜索 root 及其子树中是否存在 p 或 q
func dfs0408(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	lRes := dfs0408(root.Left, p, q)
	rRes := dfs0408(root.Right, p, q)
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	if lRes != nil && rRes != nil {
		return root
	} else if lRes == nil && rRes != nil {
		return rRes
	} else if lRes != nil && rRes == nil {
		return lRes
	} else {
		return nil
	}
}

// 示例 1：
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出：3
// 解释：节点 5 和节点 1 的最近公共祖先是节点 3。
//
// 示例 2：
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// 输出：5
// 解释：节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
func main() {
	fmt.Println("Input tree")
	tree := CreateTree()
	var n int
	fmt.Println("Input target node value1:")
	fmt.Scan(&n)
	targetNode1 := findTarget0408(tree, n)

	fmt.Println("Input target node value2:")
	fmt.Scan(&n)
	targetNode2 := findTarget0408(tree, n)
	fmt.Println(lowestCommonAncestor(tree, targetNode1, targetNode2))
}
func findTarget0408(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == target {
		return root
	}
	lRes, rRes := findTarget0408(root.Left, target), findTarget0408(root.Right, target)
	if lRes != nil {
		return lRes
	}
	return rRes
}
