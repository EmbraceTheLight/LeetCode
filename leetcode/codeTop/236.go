// 236. 二叉树的最近公共祖先
/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

提示：
树中节点数目在范围 [2, 10^5] 内。
-10^9 <= Node.val <= 10^9
所有 Node.val 互不相同 。
p != q
p 和 q 均存在于给定的二叉树中。
*/
package main

import (
	"fmt"
	"lc/pkg"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *pkg.TreeNode) *pkg.TreeNode {
	if root == nil {
		return nil
	}
	if root == p {
		return p
	}
	if root == q {
		return q
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	} else if left != nil && right == nil {
		return left
	} else if left == nil && right != nil {
		return right
	}
	return nil
}

// 示例 1：
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出：3
// 解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
//
// 示例 2：
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// 输出：5
// 解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
//
// 示例 3：
// 输入：root = [1,2], p = 1, q = 2
// 输出：1
func main() {
	var findNode func(root *pkg.TreeNode, targetVal int) *pkg.TreeNode
	findNode = func(root *pkg.TreeNode, targetVal int) *pkg.TreeNode {
		if root == nil {
			return nil
		}
		if root.Val == targetVal {
			return root
		}
		if node := findNode(root.Left, targetVal); node != nil {
			return node
		}
		if node := findNode(root.Right, targetVal); node != nil {
			return node
		}
		return nil
	}
	var left, right int
	fmt.Println("Input left, right value:")
	fmt.Scan(&left, &right)
	root := pkg.CreateTree()
	fmt.Println(lowestCommonAncestor(root, findNode(root, left), findNode(root, right)).Val)
}
