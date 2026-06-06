/*
LCR 143. 子结构判断
给定两棵二叉树 tree1 和 tree2，判断 tree2 是否以 tree1 的某个节点为根的子树具有 相同的结构和节点值 。
注意，空树 不会是以 tree1 的某个节点为根的子树具有 相同的结构和节点值 。
*/

// 第一次完全使用leetcode内置编译器完成，在这里只是搬运
package main

import (
	"fmt"
	"lc/100/pkg"
)

func check(A *pkg.TreeNode, B *pkg.TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil && B != nil {
		return false
	} else if A != nil && B == nil {
		return true
	}
	if A.Val != B.Val {
		return false
	}

	return check(A.Left, B.Left) && check(A.Right, B.Right)
}
func isSubStructure(A *pkg.TreeNode, B *pkg.TreeNode) bool {
	if B == nil {
		return false
	}
	if A == nil {
		return false
	}
	if A.Val == B.Val {
		if check(A, B) == true {
			return true
		}
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
func main() {
	root1 := pkg.CreateTree()
	root2 := pkg.CreateTree()
	fmt.Println(isSubStructure(root1, root2))
}
