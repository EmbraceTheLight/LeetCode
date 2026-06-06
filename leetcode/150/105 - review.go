// 105. 从前序与中序遍历序列构造二叉树
/*
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历，
inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

示例 1:
输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]

示例 2:
输入: preorder = [-1], inorder = [-1]
输出: [-1]

提示:
1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder 和 inorder 均 无重复 元素
inorder 均出现在 preorder
preorder 保证 为二叉树的前序遍历序列
inorder 保证 为二叉树的中序遍历序列

*/
package main

import "lc/pkg"

func buildTreeR(preorder []int, inorder []int) *pkg.TreeNode {
	return makeSubTree(preorder, inorder)
}

func makeSubTree(preorder []int, inorder []int) *pkg.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	treeNode := &pkg.TreeNode{Val: preorder[0]}
	if len(inorder) == 1 {
		return treeNode
	}

	idxInInorder := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			idxInInorder = i
			break
		}
	}

	inorderLeft := inorder[:idxInInorder]
	inorderRight := inorder[idxInInorder+1:]
	preorderLeft := preorder[1 : 1+idxInInorder]
	preorderRight := preorder[1+idxInInorder:]
	treeNode.Left = makeSubTree(preorderLeft, inorderLeft)
	treeNode.Right = makeSubTree(preorderRight, inorderRight)
	return treeNode
}

// Test Case1:	preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]	Our expected output is [3,9,20,null,null,15,7]
// Test Case2:	preorder = [-1], inorder = [-1]	Our expected output is [-1]
func main() {
	preorder := pkg.CreateSlice[int]()
	inorder := pkg.CreateSlice[int]()
	pkg.PrintTreeByLevel(buildTreeR(preorder, inorder))
}
