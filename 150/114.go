// 114. 二叉树展开为链表
/*
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

示例 1：
输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [0]
输出：[0]

提示：
树中结点数在范围 [0, 2000] 内
-100 <= Node.val <= 100

进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？
*/
package main

import "lc/pkg"

func flatten(root *pkg.TreeNode) {
	if root == nil {
		return
	}
	dfs114(root)
}

// dfs114 返回处理好的单叉树的头结点和尾结点
func dfs114(root *pkg.TreeNode) (head *pkg.TreeNode, tail *pkg.TreeNode) {
	if root.Left == nil && root.Right == nil {
		return root, root
	}

	var newTail *pkg.TreeNode

	var leftHead, leftTail *pkg.TreeNode
	var rightHead, rightTail *pkg.TreeNode

	originRight := root.Right
	if root.Left != nil {
		leftHead, leftTail = dfs114(root.Left)
		newTail = leftTail

		root.Right = leftHead
		root.Left = nil
	}
	if originRight != nil {
		rightHead, rightTail = dfs114(originRight)
		if newTail != nil {
			newTail.Right = rightHead
		} else {
			root.Right = rightHead
		}
		newTail = rightTail
	}
	return root, newTail
}

// Test Case1:	[1,2,5,3,4,null,6]	Output: [1,null,2,null,3,null,4,null,5,null,6]
// Test Case2:	[]	Output: []
// Test Case3:	[0]	Output: [0]
// Test Case4: 	[1,2,3]	Output: [1,null,2,null,3]
// Test Case5:  [1,null,2,3]	Output: [1,null,2,null,3]
func main() {
	root := pkg.CreateTree()
	flatten(root)
	pkg.PrintTreeByLevel(root)
}
