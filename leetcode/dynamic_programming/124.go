/*
124. 二叉树中的最大路径和TODO：[hard]
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。
示例 1：
输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6

示例 2：
输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42

提示：
树中节点数目范围是 [1, 3 * 10^4]
-1000 <= Node.val <= 1000
*/
package main

import (
	"fmt"
	"math"
	"lc/pkg"
)

// 返回以root为根的二叉树最大路径和
func dfs124(root *pkg.TreeNode, sum *int) int {
	if root == nil {
		return 0
	}
	lmax := max(dfs124(root.Left, sum), 0) //防止出现子树为负数的情况（若是出现这种情况，则不选）
	rmax := max(dfs124(root.Right, sum), 0)

	newMax := root.Val + lmax + rmax //将当前节点与子树最大路径和相加，形成备选最大路径和
	*sum = max(*sum, newMax)
	return root.Val + max(lmax, rmax) //选择一条最大路径返回
}

func mydfs124(root *pkg.TreeNode) int {
	if root == nil {
		return 0
	}
	lmax := mydfs124(root.Left)
	rmax := mydfs124(root.Right)

	return root.Val + max(lmax, rmax) //选择一条最大路径返回
}

func maxPathSum(root *pkg.TreeNode) int {
	sum := math.MinInt
	dfs124(root, &sum)
	return sum
}
func main() {
	root := pkg.CreateTree()
	fmt.Println(maxPathSum(root))
}
