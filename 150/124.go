// 124. 二叉树中的最大路径和
/*
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
	"lc/pkg"
)

const minInt = -3e7 - 1

func maxPathSum(root *pkg.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	var ans int = minInt

	dfs124(root, &ans)
	return ans
}

// dfs124 返回包含当前当前节点的最大路径和
func dfs124(root *pkg.TreeNode, ans *int) int {
	if root == nil {
		return minInt
	}

	sumLeft := max(dfs124(root.Left, ans), 0)
	sumRight := max(dfs124(root.Right, ans), 0)

	// 更新最大路径和
	*ans = max(*ans, sumLeft+sumRight+root.Val)
	return root.Val + max(sumLeft, sumRight)
}

// Test Case1:	[1,2,3]										Output: 6
// Test Case2:	[-10, 9, 20, null, null, 15, 7]				Output: 42
// Test Case3:  [2,-1]										Output: 2
// Test Case4:  [-2,1]										Output: 1
// Test Case5:  [-2,-1]										Output: -1
// Test Case6:	[1,-2,3]									Output: 4
// Test Case7:  [1,-2,-3,1,3,-2,null,-1]					Output: 3
// Test Case8:	[5,4,8,11,null,13,4,7,2,null, null,null,1]	Output: 48
func main() {
	root := pkg.CreateTree()
	fmt.Println(maxPathSum(root))
}
