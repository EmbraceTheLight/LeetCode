// 112. 路径总和
/*
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。
叶子节点 是指没有子节点的节点。

示例 1：
输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出：true
解释：等于目标和的根节点到叶节点路径如上图所示。

示例 2：
输入：root = [1,2,3], targetSum = 5
输出：false
解释：树中存在两条根节点到叶子节点的路径：
(1 --> 2): 和为 3
(1 --> 3): 和为 4
不存在 sum = 5 的根节点到叶子节点的路径。

示例 3：
输入：root = [], targetSum = 0
输出：false
解释：由于树是空的，所以不存在根节点到叶子节点的路径。

提示：
树中节点的数目在范围 [0, 5000] 内
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func hasPathSum(root *pkg.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return dfs112(root, 0, targetSum)
}
func dfs112(node *pkg.TreeNode, curSum, targetSum int) bool {
	if node.Left == nil && node.Right == nil {
		return node.Val+curSum == targetSum
	}

	var res bool
	if node.Left != nil {
		res = dfs112(node.Left, curSum+node.Val, targetSum)
		if res {
			return true
		}
	}
	if node.Right != nil {
		res = dfs112(node.Right, curSum+node.Val, targetSum)
		if res {
			return true
		}
	}
	return false
}

// Test Case1: targetSum = 22	[5,4,8,11,null,13,4,7,2,null,1]	 Output: true
// Test Case2: targetSum = 5	[1,2,3]	 Output: false
func main() {
	var targetSum int
	fmt.Println("Input targetSum:")
	fmt.Scan(&targetSum)
	tree := pkg.CreateTree()
	fmt.Println(hasPathSum(tree, targetSum))
}
