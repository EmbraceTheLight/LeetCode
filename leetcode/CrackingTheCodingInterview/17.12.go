// 面试题 17.12. BiNode
/*
二叉树数据结构TreeNode可用来表示单向链表（其中left置空，right为下一个链表节点）。
实现一个方法，把二叉搜索树转换为单向链表，要求依然符合二叉搜索树的性质，转换操作应是原址的，也就是在原始的二叉搜索树上直接修改。
返回转换后的单向链表的头节点。

注意：本题相对原题稍作改动

提示：
节点数量不会超过 100000。
*/
package main

import (
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
func convertBiNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	head, _ := dfs1712(root)
	return head
}

func dfs1712(node *TreeNode) (head, tail *TreeNode) {
	if node.Left == nil && node.Right == nil {
		return node, node
	}
	var leftHead, leftTail, rightHead, rightTail *TreeNode
	if node.Left != nil {
		leftHead, leftTail = dfs1712(node.Left)
	}
	if node.Right != nil {
		rightHead, rightTail = dfs1712(node.Right)
	}

	head = node
	if leftHead != nil {
		head = leftHead
		leftTail.Right = node
	}

	tail = node
	if rightHead != nil {
		tail = rightTail
		node.Right = rightHead
	}
	node.Left = nil
	return head, tail
}

// 示例：
// 输入： [4,2,5,1,3,null,6,0]
// 输出： [0,null,1,null,2,null,3,null,4,null,5,null,6]
func main() {
	root := CreateTree()
	PrintTreeByLevel(convertBiNode(root))
}
