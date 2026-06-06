/*
102. 二叉树的层序遍历

给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]

示例 2：
输入：root = [1]
输出：[[1]]

示例 3：
输入：root = []
输出：[]

提示：

树中节点数目在范围 [0, 2000] 内
-1000 <= Node.val <= 1000
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

type queueNode struct {
	node  *pkg.TreeNode
	level int
}
type queue []*queueNode

func (q *queue) push(treeNode *pkg.TreeNode, level int) {
	ins := &queueNode{treeNode, level}
	*q = append(*q, ins)
}
func (q *queue) top() *queueNode {
	return (*q)[0]
}
func (q *queue) pop() *queueNode {
	if len(*q) == 0 {
		return nil
	}
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

func levelOrder(root *pkg.TreeNode) [][]int {
	var ans = make([][]int, 0)
	var q = make(queue, 0)

	if root != nil {
		q.push(root, 1)
	}

	for len(q) > 0 {
		level := q.top().level
		var sli = make([]int, 0)
		for node := q.pop(); ; node = q.pop() {
			if node.node.Left != nil {
				q.push(node.node.Left, level+1)
			}
			if node.node.Right != nil {
				q.push(node.node.Right, level+1)
			}
			sli = append(sli, node.node.Val)
			if len(q) == 0 || q.top().level != level {
				break
			}
		}
		ans = append(ans, sli)
	}
	return ans
}
func main() {
	root := pkg.CreateTree()
	fmt.Println(levelOrder(root))
}
