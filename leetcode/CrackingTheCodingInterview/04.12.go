// 04.12. 求和路径
/*
给定一棵二叉树，其中每个节点都含有一个整数数值(该值或正或负)。
设计一个算法，打印节点数值总和等于某个给定值的所有路径的数量。
注意，路径不一定非得从二叉树的根节点或叶节点开始或结束，但是其方向必须向下(只能从父节点指向子节点方向)。

提示：
节点总数 <= 10000
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func pathSum(root *TreeNode, sum int) int {
	var ans int
	mp := make(map[int]int)
	mp[0] = 1
	dfs0412(root, 0, sum, mp, &ans)
	return ans
}
func dfs0412(root *TreeNode, sum, targetSum int, mp map[int]int, ans *int) {
	if root == nil {
		return
	}
	sum += root.Val
	if mp[sum-targetSum] != 0 {
		*ans += mp[sum-targetSum]
	}
	mp[sum]++
	dfs0412(root.Left, sum, targetSum, mp, ans)
	dfs0412(root.Right, sum, targetSum, mp, ans)
	mp[sum]--
}

// 示例：
// 给定如下二叉树 [5,4,8,11,null,13,4,7,2,null,null,5,1]，以及目标和 sum = 22，
//
//	      5
//	     / \
//	    4   8
//	   /   / \
//	  11  13  4
//	 /  \    / \
//	7    2  5   1
//
// 输出：
// 3
// 解释：和为 22 的路径有：[5,4,11,2], [5,8,4,5], [4,11,7]
func main() {
	fmt.Println("Input sum:")
	var sum int
	fmt.Scan(&sum)
	fmt.Println("Input tree:")
	tree := CreateTree()
	fmt.Println("Output:", pathSum(tree, sum))
}
