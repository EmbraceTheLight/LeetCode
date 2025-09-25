/*
450. 删除二叉搜索树中的节点
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
一般来说，删除节点可分为两个步骤：
首先找到需要删除的节点；
如果找到了，删除它。
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func del(node **pkg.TreeNode, key int) {
	dele := *node
	var tmp *pkg.TreeNode

	if *node == nil {
		return
	}
	if dele.Val == key { //找到了待删除节点
		if dele.Left == nil && dele.Right == nil {
			*node = nil //待删除节点为叶子节点，直接置为nil
		} else if dele.Left == nil && dele.Right != nil {
			*node = dele.Right //待删除节点只有右子树，将待删除节点用其右子树覆盖
		} else if dele.Left != nil && dele.Right == nil {
			*node = dele.Left
		} else { //有左子树与右子树
			pre := dele.Right
			tmp = dele.Right
			if tmp.Left == nil { //右子节点无左子树
				tmp.Left = dele.Left //接管待删除节点的左子树
			} else {
				for tmp.Left != nil { //寻找待删除节点右子树的最小值
					pre = tmp
					tmp = tmp.Left
				}
				//pre接管tmp的左子树
				pre.Left = tmp.Right
				//TODO:用tmp接管待删除节点的左右子树
				tmp.Left = dele.Left
				tmp.Right = dele.Right
			}
			*node = tmp
		}
	} else if dele.Val < key {
		del(&(dele.Right), key)
	} else if dele.Val > key {
		del(&(dele.Left), key)
	}
}
func deleteNode(root *pkg.TreeNode, key int) *pkg.TreeNode {
	del(&root, key)
	return root
}
func main() {
	println("Input cnt:")
	var k int
	fmt.Scan(&k)

	root := pkg.CreateTree()
	pkg.PrintTreeByLevel(root)
	pkg.PrintTreeByLevel(deleteNode(root, k))
}
