/*
108. 将有序数组转换为二叉搜索树
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵
平衡二叉搜索树。
示例 1：
输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

示例 2：
输入：nums = [1,3]
输出：[3,1]
解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。

提示：

1 <= nums.length <= 10^4
-10^4 <= nums[i] <= 10^4
nums 按 严格递增 顺序排列
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func dfs108(nums []int) *pkg.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	l := len(nums)
	mid := l / 2
	node := new(pkg.TreeNode)
	node.Val = nums[mid]
	node.Left = dfs108(nums[:mid])
	node.Right = dfs108(nums[mid+1:])
	return node
}
func sortedArrayToBST(nums []int) *pkg.TreeNode {
	return dfs108(nums)
}
func printLevel(root *pkg.TreeNode) {
	type s struct {
		node  *pkg.TreeNode
		level int
	}
	var q = make([]s, 0)
	q = append(q, s{root, 0})
	for len(q) > 0 {
		level := q[0].level
		fmt.Printf("level%d: ", level)
		for node := q[0].node; node != nil; node = q[0].node {
			fmt.Printf("%d ", node.Val)
			if node.Left != nil {
				q = append(q, s{node.Left, level + 1})
			}
			if node.Right != nil {
				q = append(q, s{node.Right, level + 1})
			}
			q = q[1:]
			if len(q) == 0 {
				break
			}
		}
	}
}
func main() {
	nums := pkg.CreateIntSlice()
	printLevel(sortedArrayToBST(nums))
}
