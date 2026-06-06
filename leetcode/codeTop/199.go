// 199. 二叉树的右视图
/*
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

提示:
二叉树的节点个数的范围是 [0,100]
-100 <= Node.val <= 100
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func rightSideView(root *pkg.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	type treeNode struct {
		node  *pkg.TreeNode
		level int
	}
	ans := make([]int, 0)
	queue := make([]*treeNode, 0)
	queue = append(queue, &treeNode{
		node:  root,
		level: 1,
	})
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		// 当前是最后一个节点, 或下一个节点在下一层, 说明当前节点是本层最右侧节点. 将它的值添加到答案中
		if len(queue) == 0 || top.level != queue[0].level {
			ans = append(ans, top.node.Val)
		}
		if top.node.Left != nil {
			queue = append(queue, &treeNode{
				node:  top.node.Left,
				level: top.level + 1,
			})
		}
		if top.node.Right != nil {
			queue = append(queue, &treeNode{
				node:  top.node.Right,
				level: top.level + 1,
			})
		}
	}
	return ans
}

// 示例 1：
// 输入：root = [1,2,3,null,5,null,4]
// 输出：[1,3,4]
//
// 示例 2：
// 输入：root = [1,2,3,4,null,null,null,5]
// 输出：[1,3,4,5]
//
// 示例 3：
// 输入：root = [1,null,3]
// 输出：[1,3]
//
// 示例 4：
// 输入：root = []
// 输出：[]
func main() {
	root := pkg.CreateTree()
	fmt.Println(rightSideView(root))
}
