/*
110. 平衡二叉树
给定一个二叉树，判断它是否是平衡二叉树
树中的节点数在范围 [0, 5000] 内
-10^4 <= Node.val <= 10^4
*/
package main

import (
	"math"
	"lc/100/pkg"
)

func deep(node *pkg.TreeNode) int {
	if node == nil {
		return 0
	}
	lheight := deep(node.Left)
	rheight := deep(node.Right)
	if math.Abs(float64(lheight-rheight)) > 1 || lheight == -1 || rheight == -1 {
		return -1
	}
	return max(deep(node.Left), deep(node.Right)) + 1
}

//	func isBalanced(root *pkg.TreeNode) bool {
//		if root == nil {
//			return true
//		}
//
//		left := deep(root.Left)
//		right := deep(root.Right)
//		resl := isBalanced(root.Left)
//
//		resr := isBalanced(root.Right)
//		if resl == false || resr == false {
//			return false
//		}
//		if math.Abs(float64(left-right)) > 1 {
//			return false
//		}
//		return true
//	}

func isBalanced(root *pkg.TreeNode) bool {
	if root == nil {
		return true
	}
	return deep(root) > 0
}
func main() {
	root := pkg.CreateTree()
	println(isBalanced(root))
}
