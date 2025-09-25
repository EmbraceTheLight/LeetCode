/*
958. 二叉树的完全性检验
给你一棵二叉树的根节点 root ，请你判断这棵树是否是一棵 完全二叉树 。
在一棵 完全二叉树 中，除了最后一层外，所有层都被完全填满，并且最后一层中的所有节点都尽可能靠左。
最后一层（第 h 层）中可以包含 1 到 2h 个节点。
树中节点数目在范围 [1, 100] 内
1 <= Node.val <= 1000
*/
package main

import (
	"math"
	"lc/100/pkg"
)

func isCompleteTree(root *pkg.TreeNode) bool {
	level := 1
	que := make([]*pkg.TreeNode, 0)
	que = append(que, root)
	var lastLevel = false
	for len(que) > 0 {
		c := math.Pow(2, float64(level-1))
		cir := int(c)
		var n int
		for i := 0; i < cir; i++ {
			if que[i] == nil {
				lastLevel = true
				n = i
				break
			}

			que = append(que, que[i].Left)
			que = append(que, que[i].Right)
		}
		if lastLevel == true {
			que = que[n:]
			sz := len(que)
			for i := 0; i < sz; i++ {
				if lastLevel == true && que[i] != nil {
					return false
				}
			}
			return true
		}
		level++
		que = que[cir:]
	}
	return true
}
func main() {
	root := pkg.CreateTree()
	println(isCompleteTree(root))
}
