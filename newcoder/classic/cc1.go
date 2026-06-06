// CC1 二叉树的最小深度
/*
描述
求给定二叉树的最小深度。最小深度是指树的根结点到最近叶子结点的最短路径上结点的数量。
*/
package main

import (
	"fmt"
	"lc/pkg"
	"math"
)

func minDepth(root *pkg.TreeNode) int {
	if root == nil {
		return 0
	}
	return dfsCC1(root, 0)
}
func dfsCC1(node *pkg.TreeNode, curDepth int) int {
	if node.Left == nil && node.Right == nil {
		return curDepth + 1
	}
	leftMin, rightMin := math.MaxInt, math.MaxInt
	if node.Left != nil {
		leftMin = dfsCC1(node.Left, curDepth+1)
	}
	if node.Right != nil {
		rightMin = dfsCC1(node.Right, curDepth+1)
	}
	return min(leftMin, rightMin)
}

// 示例1
// 输入: [1,2,3,4,5]
// 返回值: 2
func main() {
	root := pkg.CreateTree()
	fmt.Println(minDepth(root))
}
