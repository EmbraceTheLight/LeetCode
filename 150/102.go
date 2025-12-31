// 102. 二叉树的层序遍历
/*
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
	"lc/pkg"
)

func levelOrder(root *pkg.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	type node struct {
		node  *pkg.TreeNode
		level int
	}
	var ans [][]int
	var levelVals []int
	queue := make([]*node, 0)

	curLevel := 1
	queue = append(queue, &node{
		node:  root,
		level: 1,
	})
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if top.level > curLevel {
			curLevel++
			ans = append(ans, levelVals)
			levelVals = make([]int, 0)
		}

		levelVals = append(levelVals, top.node.Val)

		if top.node.Left != nil {
			queue = append(queue, &node{
				node:  top.node.Left,
				level: top.level + 1,
			})
		}

		if top.node.Right != nil {
			queue = append(queue, &node{
				node:  top.node.Right,
				level: top.level + 1,
			})
		}
	}
	ans = append(ans, levelVals)
	return ans
}

// Test Case1:	[3, 9, 20, null, null, 15, 7]		Output:	[[3],[9,20],[15,7]]
// Test Case2:	[1]			Output:	[[1]]
// Test Case3:	[]			Output:	[]
func main() {
	root := pkg.CreateTree()
	fmt.Println(levelOrder(root))
}
