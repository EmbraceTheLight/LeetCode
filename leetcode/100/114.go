/*
114. 二叉树展开为链表
给你二叉树的根结点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
*/
package main

import "lc/100/pkg"

//func dfs114(node, parent *pkg.TreeNode) {
//	if node == nil {
//		return
//	}
//	dfs114(node.Left, node)
//
//	var tmp *pkg.TreeNode
//	tmp = node
//	if parent != nil && parent.Right != node {
//		for tmp.Right != nil {
//			tmp = tmp.Right
//		}
//
//		tmp.Right = parent.Right
//		parent.Right = node
//		parent.Left = nil
//	}
//
//	dfs114(node.Right, node)
//}

//func flatten(root *pkg.TreeNode) {
//	if root == nil {
//		return
//	}
//	dfs114(root, nil)
//}

// 非递归解法,效果相同，但是更好理解
func flatten(root *pkg.TreeNode) {
	if root == nil {
		return
	}
	node := root
	for node != nil {
		if node.Left == nil {
			node = node.Right
			continue
		}
		tmp := node.Left //找到当前节点的左子树
		tail := tmp      //tail存放左子树的最右侧节点
		for tail.Right != nil {
			tail = tail.Right
		}
		tail.Right = node.Right
		node.Right = tmp
		node.Left = nil
	}
}
func main() {
	root := pkg.CreateTree()
	flatten(root)
	pkg.PrintTreeByLevel(root)
}
