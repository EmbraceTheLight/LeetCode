/*
124. 二叉树中的最大路径和
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点root，返回其最大路径和。
树中节点数目范围是 [1, 3 * 10^4]
-1000 <= Node.val <= 1000
*/
package main

import (
	"fmt"
	"math"
	"lc/100/pkg"
)

func GetMax(root *pkg.TreeNode, sum *int) int {
	if root == nil {
		return 0
	}
	lmax := max(GetMax(root.Left, sum), 0)
	rmax := max(GetMax(root.Right, sum), 0)

	newMax := root.Val + lmax + rmax
	//m := max(lmax, rmax)
	//*sum = max(m+root.Val, *sum)
	*sum = max(*sum, newMax)

	return root.Val + max(lmax, rmax)

}
func maxPathSum(root *pkg.TreeNode) int {
	var sum = math.MinInt32
	GetMax(root, &sum)
	return sum
}
func main() {
	root := pkg.CreateTree()
	fmt.Println(maxPathSum(root))
}
