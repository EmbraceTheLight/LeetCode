// 04.02. 最小高度树
/*
给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。
*/
package main

import . "lc/pkg"

func sortedArrayToBST(nums []int) *TreeNode {
	return dfs0402(nums)
}
func dfs0402(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{}
	mid := len(nums) / 2
	root.Val = nums[mid]
	root.Left = dfs0402(nums[:mid])
	root.Right = dfs0402(nums[mid+1:])
	return root
}

// 示例：
// 给定有序数组: [-10,-3,0,5,9],
// 一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
//
//	     0
//	    / \
//	  -3   9
//	  /   /
//	-10  5
func main() {
	nums := CreateSlice[int]()
	PrintTreeByLevel(sortedArrayToBST(nums))
}
