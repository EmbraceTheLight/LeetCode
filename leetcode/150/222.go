// 222. 完全二叉树的节点个数
/*
给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。
若最底层为第 h 层（从第 0 层开始），则该层包含 1  ~ 2^h 个节点。

示例 1：
输入：root = [1,2,3,4,5,6]
输出：6

示例 2：
输入：root = []
输出：0

示例 3：
输入：root = [1]
输出：1

提示：
树中节点的数目范围是[0, 5 * 10^4]
0 <= Node.val <= 5 * 10^4
题目数据保证输入的树是 完全二叉树
进阶：遍历树来统计节点是一种时间复杂度为 O(n) 的简单解决方案。你可以设计一个更快的算法吗？
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func countNodes(root *pkg.TreeNode) int {
	var ans int
	dfs222(root, &ans)
	return ans
}
func dfs222(node *pkg.TreeNode, ans *int) {
	if node == nil {
		return
	}
	*ans += 1
	dfs222(node.Left, ans)
	dfs222(node.Right, ans)
}

// 利用完全二叉树的性质, 遍历子树的左右节点，如果左右节点高度相等，说明当前子树为满二叉树
// 否则，不为满二叉树，此时继续遍历该子树的左右节点，直至遍历完所有节点\
// 时间复杂度 O(logn*logn)
func countNodesBest(root *pkg.TreeNode) int {
	if root == nil {
		return 0
	}

	// 遍历当前节点的左右子树
	left, right := root, root
	leftHeight, rightHeight := 0, 0
	for left != nil {
		left = left.Left
		leftHeight++
	}

	for right != nil {
		right = right.Right
		rightHeight++
	}

	// 当前节点的左右子树高度相等，说明是满二叉树，返回 2^height - 1，停止遍历当前节点
	if leftHeight == rightHeight {
		return 1<<leftHeight - 1
	}

	// 否则，继续遍历当前节点的左右子树
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// Test Case1: [1,2,3,4,5,6]	Output: 6
// Test Case2: []	Output: 0
// Test Case3: [1]	Output: 1
func main() {
	root := pkg.CreateTree()
	fmt.Println(countNodes(root))
}
