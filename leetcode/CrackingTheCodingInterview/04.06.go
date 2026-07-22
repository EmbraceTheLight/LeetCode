// 04.06. 后继者
/*
设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。
如果指定节点没有对应的“下一个”节点，则返回null。
*/
package main

import (
	"fmt"
	. "lc/pkg"
)

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	treeNodeList := make([]*TreeNode, 0)
	var node *TreeNode
	// 使用栈对树进行中序遍历
	for len(stack) > 0 {
		node = stack[len(stack)-1]
		for node.Left != nil {
			if node == p {
				break
			}
			node = node.Left
			stack = append(stack, node)
		}

		for len(stack) > 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			treeNodeList = append(treeNodeList, node)
			if node.Right != nil {
				stack = append(stack, node.Right)
				break
			}
		}
	}
	idx := binarySearch0406(treeNodeList, p)
	if idx == len(treeNodeList)-1 {
		return nil
	}
	return treeNodeList[idx+1]
}

// binarySearch0406 对 []*TreeNode 结构进行二分查找
// 根据 TreeNode 的 Val 属性进行二分查找. 返回索引下标
// 要求 []*TreeNode 已经按照 val 进行升序排序
func binarySearch0406(sli []*TreeNode, target *TreeNode) int {
	left, right := 0, len(sli)
	for left <= right {
		mid := (left + right) / 2
		if sli[mid].Val < target.Val {
			left = mid + 1
		} else if sli[mid].Val > target.Val {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// 示例 1：
// 输入：root = [2,1,3], p = 1
//
//	 2
//	/ \
//
// 1   3
// 输出：2
//
// 示例 2：
// 输入：root = [5,3,6,2,4,null,null,1], p = 6
//
//	     5
//	    / \
//	   3   6
//	  / \
//	 2   4
//	/
//
// 1
// 输出：null
func main() {
	fmt.Println("Input tree")
	tree := CreateTree()
	var n int
	fmt.Println("Input target node value:")
	fmt.Scan(&n)
	targetNode := findTarget(tree, n)
	fmt.Println(inorderSuccessor(tree, targetNode))
}

func findTarget(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	node := root
	for node != nil {
		if node.Val < target {
			node = node.Right
		} else if node.Val > target {
			node = node.Left
		} else {
			return node
		}
	}
	return nil
}
