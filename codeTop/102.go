// 102. 二叉树的层序遍历
/*
	给你二叉树的根节点 root ，返回其节点值的 层序遍历。（即逐层地，从左到右访问所有节点）。
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func levelOrder(root *pkg.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	tmp := make([]int, 0)
	type treeNodeHelper struct {
		node  *pkg.TreeNode
		level int
	}

	curLevel := 1
	nodeList := make([]treeNodeHelper, 0)
	nodeList = append(nodeList, treeNodeHelper{node: root, level: 1})
	for len(nodeList) > 0 {
		curNode := nodeList[0]
		nodeList = nodeList[1:]
		if curLevel != curNode.level {
			ans = append(ans, tmp)
			curLevel = curNode.level
			tmp = make([]int, 0)
		}
		tmp = append(tmp, curNode.node.Val)
		if curNode.node.Left != nil {
			nodeList = append(nodeList, treeNodeHelper{node: curNode.node.Left, level: curNode.level + 1})
		}
		if curNode.node.Right != nil {
			nodeList = append(nodeList, treeNodeHelper{node: curNode.node.Right, level: curNode.level + 1})
		}
	}
	ans = append(ans, tmp)
	return ans
}

type queueNode struct {
	node *pkg.TreeNode
	idx  float64
}
type queue []*queueNode

func (q *queue) push(treeNode *pkg.TreeNode, idx float64) {
	ins := &queueNode{treeNode, idx}
	*q = append(*q, ins)
}
func (q *queue) pop() *queueNode {
	if len(*q) == 0 {
		return nil
	}
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

// 示例 1：
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
//
// 示例 2：
// 输入：root = [1]
// 输出：[[1]]
//
// 示例 3：
// 输入：root = []
// 输出：[]
func main() {
	root := pkg.CreateTree()
	fmt.Println(levelOrder(root))
}
