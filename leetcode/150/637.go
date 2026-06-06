// 637. 二叉树的层平均值
/*
给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：[3.00000,14.50000,11.00000]
解释：第 0 层的平均值为 3,第 1 层的平均值为 14.5,第 2 层的平均值为 11 。
因此返回 [3, 14.5, 11] 。

示例 2:
输入：root = [3,9,20,15,7]
输出：[3.00000,14.50000,11.00000]

提示：
树中节点数量在 [1, 10^4] 范围内
-2^31 <= Node.val <= 2^31 - 1
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func averageOfLevels(root *pkg.TreeNode) []float64 {
	var ans []float64
	var levelSum float64
	type node struct {
		node  *pkg.TreeNode
		level int
	}

	curLevel := 1
	levelCount := 0
	queue := []*node{&node{node: root, level: 1}}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		if top.level > curLevel {
			ans = append(ans, levelSum/float64(levelCount))
			curLevel++
			levelCount = 0
			levelSum = 0
		}
		levelCount++
		levelSum += float64(top.node.Val)
		if top.node.Left != nil {
			queue = append(queue, &node{node: top.node.Left, level: top.level + 1})
		}
		if top.node.Right != nil {
			queue = append(queue, &node{node: top.node.Right, level: top.level + 1})
		}
	}
	ans = append(ans, levelSum/float64(levelCount))
	return ans
}

// Test Case1: [3,9,20,null,null,15,7]	Output: [3.00000,14.50000,11.00000]
// Test Case2: [3,9,20,15,7]	Output: [3.00000,14.50000,11.00000]
func main() {
	root := pkg.CreateTree()
	fmt.Println(averageOfLevels(root))
}
