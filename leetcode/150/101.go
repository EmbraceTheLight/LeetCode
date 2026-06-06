// 101. 对称二叉树
/*
给你一个二叉树的根节点 root ， 检查它是否轴对称。

示例 1：
输入：root = [1,2,2,3,4,4,3]
输出：true

示例 2：
输入：root = [1,2,2,null,3,null,3]
输出：false

提示：
树中节点数目在范围 [1, 1000] 内
-100 <= Node.val <= 100

进阶：你可以运用递归和迭代两种方法解决这个问题吗？
*/
package main

import (
	"fmt"
	"lc/pkg"
	"math"
)

func isSymmetric(root *pkg.TreeNode) bool {
	return cmpTwoNodes(root, root)
}
func cmpTwoNodes(left, right *pkg.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return cmpTwoNodes(left.Left, right.Right) && cmpTwoNodes(left.Right, right.Left)
}

func isSymmetricBFS(root *pkg.TreeNode) bool {
	type pair struct {
		level int
		id    int
		node  *pkg.TreeNode
	}

	cmpOneLevel := func(level int, nodes []*pair) bool {
		if len(nodes)%2 != 0 {
			return false
		}
		start, end := 0, len(nodes)-1
		sum := int(math.Pow(2, float64(level-1)) + math.Pow(2, float64(level)) - 1)
		for ; start < end; start, end = start+1, end-1 {
			if nodes[start].id+nodes[end].id != sum || nodes[start].node.Val != nodes[end].node.Val {
				return false
			}
		}
		return true
	}

	curLevel := 1
	queue := []*pair{&pair{1, 1, root}}
	levelNodes := make([]*pair, 0)
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top.level > curLevel {
			if !cmpOneLevel(curLevel, levelNodes) {
				return false
			}
			curLevel = top.level
			levelNodes = make([]*pair, 0)
		}

		levelNodes = append(levelNodes, top)

		if top.node.Left != nil {
			queue = append(queue, &pair{top.level + 1, top.id * 2, top.node.Left})
		}
		if top.node.Right != nil {
			queue = append(queue, &pair{top.level + 1, top.id*2 + 1, top.node.Right})
		}

	}

	if !cmpOneLevel(curLevel, levelNodes) {
		return false
	}
	return true
}

// Test Case1:	[1,2,2,3,4,4,3]	Output: true
// Test Case2:	[1,2,2,null,3,null,3]	Output: false
func main() {
	root := pkg.CreateTree()
	fmt.Println(isSymmetric(root))
	fmt.Println(isSymmetricBFS(root))
}
