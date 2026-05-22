// 94. 二叉树的中序遍历
/*
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

提示：
树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100

进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/
package main

import (
	"fmt"
	"lc/pkg"
)

// 示例 1：
// 输入：root = [1,null,2,3]
// 输出：[1,3,2]
//
// 示例 2：
// 输入：root = []
// 输出：[]
//
// 示例 3：
// 输入：root = [1]
// 输出：[1]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *pkg.TreeNode) []int {
	var ans []int
	dfs94(root, &ans)
	return ans
}
func dfs94(node *pkg.TreeNode, ans *[]int) {
	if node == nil {
		return
	}
	dfs94(node.Left, ans)
	*ans = append(*ans, node.Val)
	dfs94(node.Right, ans)
}

func inorderTraversal2(root *pkg.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var ans []int
	type treeNode struct {
		parent *pkg.TreeNode
		node   *pkg.TreeNode
	}
	stk := make([]*treeNode, 0)
	stk = append(stk, &treeNode{
		parent: nil,
		node:   root,
	})

	for len(stk) > 0 {
		top := stk[len(stk)-1]
		if top.node.Left != nil {
			stk = append(stk, &treeNode{
				parent: top.node,
				node:   top.node.Left,
			})
		} else if top.node.Right != nil {
			ans = append(ans, top.node.Val)
			stk = stk[:len(stk)-1]
			if top.parent != nil {
				top.parent.Left = nil
			}
			stk = append(stk, &treeNode{
				parent: top.node,
				node:   top.node.Right,
			})
		} else {
			ans = append(ans, top.node.Val)
			if top.parent != nil {
				top.parent.Left = nil
			}
			stk = stk[:len(stk)-1]
		}
	}
	return ans
}

func main() {
	tree := pkg.CreateTree()
	fmt.Println(inorderTraversal(tree))
	fmt.Println(inorderTraversal2(tree))
}
