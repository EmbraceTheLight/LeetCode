// 199. 二叉树的右视图
/*
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例 1：
输入：root = [1,2,3,null,5,null,4]
输出：[1,3,4]

示例 2：
输入：root = [1,2,3,4,null,null,null,5]
输出：[1,3,4,5]

示例 3：
输入：root = [1,null,3]
输出：[1,3]

示例 4：
输入：root = []
输出：[]


提示:
二叉树的节点个数的范围是 [0,100]
-100 <= Node.val <= 100
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func rightSideViewR(root *pkg.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	type node struct {
		node  *pkg.TreeNode
		level int
	}
	ans := make([]int, 0)
	queue := make([]*node, 0)
	queue = append(queue, &node{root, 1})

	curLevel := 1
	tmp := root
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top.level > curLevel {
			ans = append(ans, tmp.Val)
			curLevel++
		}
		tmp = top.node
		if top.node.Left != nil {
			queue = append(queue, &node{top.node.Left, top.level + 1})
		}
		if top.node.Right != nil {
			queue = append(queue, &node{top.node.Right, top.level + 1})
		}
	}
	ans = append(ans, tmp.Val)
	return ans
}

// Test Case1:	[1,2,3,null,5,null,4]	Output:[1,3,4]
// Test Case2:	[1,2,3,4,null,null,null,5]	Output:[1,3,4,5]
// Test Case3:	[1,null,3]	Output:[1,3]
// Test Case4:	[]	Output:[]
func main() {
	root := pkg.CreateTree()
	fmt.Println(rightSideViewR(root))
}
