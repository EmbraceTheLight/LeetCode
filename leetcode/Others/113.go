/*
113. 路径总和 II
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
叶子节点 是指没有子节点的节点。
树中节点总数在范围 [0, 5000] 内
-1000 <= Node.val <= 1000
-1000 <= targetSum <= 1000
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func find(node *pkg.TreeNode, target, sum int, ret []int, ans *[][]int) {
	if node.Left == nil && node.Right == nil {
		sum += node.Val
		if sum == target {
			ret = append(ret, node.Val)
			ins := make([]int, len(ret)) //要深复制切片，否则会有问题
			copy(ins, ret)
			*ans = append(*ans, ins)
		}
		return
	}
	sum += node.Val
	ret = append(ret, node.Val)
	if node.Left != nil {
		find(node.Left, target, sum, ret, ans)
	}
	if node.Right != nil {
		find(node.Right, target, sum, ret, ans)
	}
}

func pathSum(root *pkg.TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := make([][]int, 0)
	ret := make([]int, 0)
	find(root, targetSum, 0, ret, &ans)
	return ans
}
func main() {
	fmt.Println("Input Target:")
	var target int
	fmt.Scan(&target)
	root := pkg.CreateTree()
	fmt.Println(pathSum(root, target))
}
