/*
662. 二叉树最大宽度
给你一棵二叉树的根节点 root ，返回树的 最大宽度 。
树的 最大宽度 是所有层中最大的 宽度 。
每一层的 宽度 被定义为该层最左和最右的非空节点（即，两个端点）之间的长度。将这个二叉树视作与满二叉树结构相同，
两端点间会出现一些延伸到这一层的 null 节点，这些 null 节点也计入长度。
题目数据保证答案将会在  32 位 带符号整数范围内。
*/
package main

import (
	"lc/pkg"
)

type treeQueue struct {
	node  *pkg.TreeNode
	index int //节点编号
}

// #################################################################
func widthOfBinaryTree(root *pkg.TreeNode) int {
	if root == nil {
		return 0
	}
	var ret int
	q := treeQueue{
		node:  root,
		index: 1,
	}
	var que = make([]treeQueue, 0)
	que = append(que, q)
	for len(que) > 0 {
		var levelSize = len(que) //标识该层节点数
		first := que[0]          //记录该层第一个节点数
		last := first
		for i := 0; i < levelSize; i++ {
			if que[i].node.Left != nil {
				que = append(que, treeQueue{
					node:  que[i].node.Left,
					index: que[i].index * 2,
				})
			}
			if que[i].node.Right != nil {
				que = append(que, treeQueue{
					node:  que[i].node.Right,
					index: que[i].index*2 + 1,
				})
			}
			last = que[i] //更新该层最后一个节点的位置
		}
		que = que[levelSize:]

		ret = max(ret, last.index-first.index+1)

	}
	return ret
}
func main() {
	root := pkg.CreateTree()
	println(widthOfBinaryTree(root))
}
