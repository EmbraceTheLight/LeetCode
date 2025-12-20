// 106. 从中序与后序遍历序列构造二叉树
/*
给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树。


示例 1:
输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出：[3,9,20,null,null,15,7]

示例 2:
输入：inorder = [-1], postorder = [-1]
输出：[-1]
提示:

1 <= inorder.length <= 3000
postorder.length == inorder.length
-3000 <= inorder[i], postorder[i] <= 3000
inorder 和 postorder 都由 不同 的值组成
postorder 中每一个值都在 inorder 中
inorder 保证是树的中序遍历
postorder 保证是树的后序遍历
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func buildTree106(inorder []int, postorder []int) *pkg.TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	lastIdx := len(postorder) - 1
	val := postorder[lastIdx]

	node := &pkg.TreeNode{Val: val}
	idxInorder := 0
	for i := 0; i >= 0; i++ {
		if inorder[i] == val {
			idxInorder = i
			break
		}
	}

	node.Left = buildTree106(inorder[:idxInorder], postorder[:idxInorder])
	node.Right = buildTree106(inorder[idxInorder+1:], postorder[idxInorder:lastIdx])
	return node
}

// Test Case1: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]	Output: [3,9,20,null,null,15,7]
// Test Case2: inorder = [-1], postorder = [-1]	Output: [-1]
func main() {
	fmt.Println("Input: inorder:")
	inorder := pkg.CreateSlice[int]()
	fmt.Println("Input: postorder:")
	postorder := pkg.CreateSlice[int]()
	pkg.PrintTreeByLevel(buildTree106(inorder, postorder))
}
