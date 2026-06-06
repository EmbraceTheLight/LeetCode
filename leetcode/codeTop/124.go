// 124. 二叉树中的最大路径和
/*
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。
同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。

提示：
树中节点数目范围是 [1, 3 * 10^4]
-1000 <= Node.val <= 1000
*/
package main

import (
	"fmt"
	"lc/pkg"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *pkg.TreeNode) int {
	var ans int = math.MinInt16
	dfs124(root, &ans)
	return ans
}
func dfs124(root *pkg.TreeNode, ans *int) int {
	if root == nil {
		return math.MinInt16
	}
	left := dfs124(root.Left, ans)
	right := dfs124(root.Right, ans)
	*ans = max(*ans, left+right+root.Val, left+root.Val, right+root.Val, root.Val)
	return max(left+root.Val, right+root.Val, root.Val)
}

// 示例 1：
// 输入：root = [1,2,3]
// 输出：6
// 解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
//
// 示例 2：
// 输入：root = [-10,9,20,null,null,15,7]
// 输出：42
// 解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
func main() {
	treeRoot := pkg.CreateTree()
	fmt.Println(maxPathSum(treeRoot))
}
