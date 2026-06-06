/*
129. 求根节点到叶节点数字之和
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：
例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。
叶节点 是指没有子节点的节点。
树中节点的数目在范围 [1, 1000] 内
0 <= Node.val <= 9
树的深度不超过 10
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func sum129(node *pkg.TreeNode, cur int) int {
	if node.Left == nil && node.Right == nil { // 到叶子节点
		return cur*10 + node.Val
	}
	var left, right int
	cur = node.Val + cur*10 //更新当前路径总和
	if node.Left != nil {
		left = sum129(node.Left, cur)
	}
	if node.Right != nil {
		right = sum129(node.Right, cur)
	}
	return left + right
}
func sumNumbers(root *pkg.TreeNode) int {
	var ans int
	ans = sum129(root, 0)
	return ans
}

// BFSsumNumbers BFS解法：维护两个队列，一个存节点，另一个存和
func BFSsumNumbers(root *pkg.TreeNode) int {
	var nodeQue = make([]*pkg.TreeNode, 0)
	var numQue = make([]int, 0)
	var ans int
	nodeQue = append(nodeQue, root)
	numQue = append(numQue, root.Val)
	for len(nodeQue) > 0 {
		var sz = len(nodeQue)
		for i := 0; i < sz; i++ {
			if nodeQue[i].Left != nil {
				nodeQue = append(nodeQue, nodeQue[i].Left)
				numQue = append(numQue, numQue[i]*10+nodeQue[i].Left.Val)
			}
			if nodeQue[i].Right != nil {
				nodeQue = append(nodeQue, nodeQue[i].Right)
				numQue = append(numQue, numQue[i]*10+nodeQue[i].Right.Val)
			}
			if nodeQue[i].Left == nil && nodeQue[i].Right == nil {
				ans += numQue[i]
			}
		}
		nodeQue = nodeQue[sz:]
		numQue = numQue[sz:]
	}
	return ans
}
func main() {
	root := pkg.CreateTree()
	fmt.Println(sumNumbers(root))
	fmt.Println(BFSsumNumbers(root))
}
