// 103. 二叉树的锯齿形层序遍历
/*
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[20,9],[15,7]]

示例 2：
输入：root = [1]
输出：[[1]]

示例 3：
输入：root = []
输出：[]

提示：
树中节点数目在范围 [0, 2000] 内
-100 <= Node.val <= 100
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func zigzagLevelOrderR(root *pkg.TreeNode) [][]int {
	type node struct {
		node  *pkg.TreeNode
		level int
	}
	if root == nil {
		return [][]int{}
	}
	curLevel := 1
	queue := []*node{&node{node: root, level: curLevel}}
	var result [][]int
	var tmp []int
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top.level > curLevel {
			result = append(result, tmp)
			tmp = make([]int, 0)
			curLevel = top.level
		}
		tmp = append(tmp, top.node.Val)
		if top.node.Left != nil {
			queue = append(queue, &node{node: top.node.Left, level: top.level + 1})
		}
		if top.node.Right != nil {
			queue = append(queue, &node{node: top.node.Right, level: top.level + 1})
		}
	}
	result = append(result, tmp)
	for i := 1; i < curLevel; i += 2 {
		for j := 0; j < len(result[i])/2; j++ {
			result[i][j], result[i][len(result[i])-1-j] = result[i][len(result[i])-1-j], result[i][j]
		}
	}
	return result
}

// Test Case1:	[3,9,20,null,null,15,7]		Output:	[[3],[20,9],[15,7]]
// Test Case2:	[1]			Output:	[[1]]
// Test Case3:	[]			Output:	[]
func main() {
	root := pkg.CreateTree()
	fmt.Println(zigzagLevelOrderR(root))
}
