// 105. 从前序与中序遍历序列构造二叉树
/*
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

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

func buildTree(preorder []int, inorder []int) *pkg.TreeNode {
	var dfs105 func(preorder []int, inorder []int) *pkg.TreeNode
	dfs105 = func(preorder []int, inorder []int) *pkg.TreeNode {
		if len(inorder) == 0 || len(preorder) == 0 {
			return nil
		}
		rootVal := preorder[0]
		root := &pkg.TreeNode{Val: rootVal}
		cnt := 0
		for i := 0; i < len(inorder); i++ {
			if inorder[i] == rootVal {
				break
			}
			cnt++
		}
		root.Left = dfs105(preorder[1:cnt+1], inorder[:cnt])
		root.Right = dfs105(preorder[cnt+1:], inorder[cnt+1:])
		return root
	}
	ans := dfs105(preorder, inorder)
	return ans
}

// 示例 1:
// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
//
// 示例 2:
// 输入: preorder = [-1], inorder = [-1]
// 输出: [-1]
func main() {
	preorder := pkg.CreateSlice[int]()
	inorder := pkg.CreateSlice[int]()
	pkg.PrintTreeByLevel(buildTree(preorder, inorder))
}
