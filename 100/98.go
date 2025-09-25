/*
98. 验证二叉搜索树
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效 二叉搜索树定义如下：
节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
树中节点数目范围在[1, 10^4] 内
-2^31 <= Node.val <= 2^31 - 1
*/
package main

import "lc/100/pkg"

func dfs(cur *pkg.TreeNode, min, max *pkg.TreeNode) bool {
	if cur == nil {
		return true
	}
	if min != nil && cur.Val <= min.Val {
		return false
	}
	if max != nil && cur.Val >= max.Val {
		return false
	}
	return dfs(cur.Left, min, cur) && dfs(cur.Right, cur, max)
}
func isValidBST(root *pkg.TreeNode) bool {
	return dfs(root, nil, nil)
}
func main() {
	root := pkg.CreateTree()
	println(isValidBST(root))
}
