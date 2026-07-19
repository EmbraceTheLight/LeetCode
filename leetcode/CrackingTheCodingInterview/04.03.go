// 04.03. 特定深度节点链表
/*
给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。
返回一个包含所有深度的链表的数组。
*/
package main

import (
	. "lc/pkg"
)

func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return nil
	}
	ans := make([]*ListNode, 0)
	type treeNode struct {
		node  *TreeNode
		level int
	}
	curLevel := 1
	queue := make([]*treeNode, 0)
	queue = append(queue, &treeNode{tree, 1})
	dummy := &ListNode{}
	cur := dummy
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top.level > curLevel {
			ans = append(ans, dummy.Next)
			dummy = &ListNode{}
			cur = dummy
			curLevel = top.level
		}
		cur.Next = &ListNode{Val: top.node.Val}
		cur = cur.Next
		if top.node.Left != nil {
			queue = append(queue, &treeNode{top.node.Left, top.level + 1})
		}
		if top.node.Right != nil {
			queue = append(queue, &treeNode{top.node.Right, top.level + 1})
		}
	}
	ans = append(ans, dummy.Next)
	return ans
}

// 示例：
// 输入：[1,2,3,4,5,null,7,8]
//
//	      1
//	     /  \
//	    2    3
//	   / \    \
//	  4   5    7
//	 /
//	8
//
// 输出：[[1],[2,3],[4,5,7],[8]]
func main() {
	tree := CreateTree()
	for _, list := range listOfDepth(tree) {
		PrintList(list)
	}
}
