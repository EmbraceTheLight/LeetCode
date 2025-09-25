/*
101. 对称二叉树
给你一个二叉树的根节点 root，检查它是否轴对称。
树中节点数目在范围 [1, 1000] 内
-100 <= Node.val <= 100
*/
package main

import (
	"lc/100/pkg"
)

func isSymmetric(root *pkg.TreeNode) bool {
	var que []*pkg.TreeNode
	que = append(que, root)
	for len(que) > 0 {
		var levelSize = len(que)

		for i := 0; i < levelSize/2; i++ {
			//节点与对称节点
			left := que[i]
			right := que[levelSize-i-1]
			//TODO:查看对称节点的值是否相等
			if left == nil && right == nil {
				continue
			}
			if (left == nil && right != nil) || (left != nil && right == nil) || left.Val != right.Val {
				return false
			}

			//将下一层左半边子节点添加到que中
			que = append(que, left.Left)
			que = append(que, left.Right)
		}
		for i := levelSize / 2; i < levelSize; i++ { //将下一层右半边子节点添加到que中
			left := que[i]
			if left != nil {
				que = append(que, left.Left)
				que = append(que, left.Right)
			}

		}
		que = que[levelSize:]
	}
	return true
}
func main() {
	root := pkg.CreateTree()
	println(isSymmetric(root))
}
