// 04.05. 合法二叉搜索树
/*
实现一个函数，检查一棵二叉树是否为二叉搜索树。
*/
package main

import (
	"fmt"
	. "lc/pkg"
	"math"
)

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, _, res := dfs0405(root)
	return res
}
func dfs0405(root *TreeNode) (int, int, bool) {
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val, true
	}
	var lMin, lMax, rMin, rMax int = math.MaxInt32, math.MinInt32, math.MaxInt32, math.MinInt32
	var lRes, rRes bool = true, true
	if root.Left != nil {
		lMin, lMax, lRes = dfs0405(root.Left)
	}
	if root.Right != nil {
		rMin, rMax, rRes = dfs0405(root.Right)
	}
	if lRes == false || rRes == false {
		return 0, 0, false
	}
	if root.Left != nil && root.Right == nil && root.Val > lMax {
		return min(lMin, rMin, root.Val), max(lMax, rMax, root.Val), true
	} else if root.Left == nil && root.Right != nil && root.Val < rMin {
		return min(lMin, rMin, root.Val), max(lMax, rMax, root.Val), true
	} else if root.Left != nil && root.Right != nil && root.Val > lMax && root.Val < rMin {
		return min(lMin, rMin, root.Val), max(lMax, rMax, root.Val), true
	} else {
		return 0, 0, false
	}
}
func isValidBST2(root *TreeNode) bool {
	return dfs0405_2(root, math.MinInt64, math.MaxInt64)
}

// dfs0405_2 minValue, maxValue 限定了 root.Val 的范围, 即 minValue < root.Val < maxValue
// 该递归函数使用先序遍历的方式遍历二叉树, 遍历到子树时, 更新 minValue, maxValue
func dfs0405_2(root *TreeNode, minValue, maxValue int) bool {
	if root == nil {
		return true
	}
	if root.Val <= minValue || root.Val >= maxValue {
		return false
	}
	return dfs0405_2(root.Left, minValue, root.Val) && dfs0405_2(root.Right, root.Val, maxValue)
}

// 示例 1：
// 输入：[2,1,3]
//
//	  2
//	 / \
//	1   3
//
// 输出：true
//
// 示例 2：
// 输入：[5,1,4,null,null,3,6]
//
//	  5
//	 / \
//	1   4
//	   / \
//	  3   6
//
// 输出：false
// 解释：输入为: [5,1,4,null,null,3,6]。
//
//	根节点的值为 5 ，但是其右子节点值为 4 。
//
// 示例 3：
// 输入：[10,5,15,null,null,6,20]
//
//	  10
//	 / \
//	5   15
//	   / \
//	  6   20
//
// 输出: false
func main() {
	tree := CreateTree()
	fmt.Println(isValidBST(tree))
	fmt.Println(isValidBST2(tree))
}
