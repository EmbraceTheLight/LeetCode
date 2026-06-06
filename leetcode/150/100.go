// 100. 相同的树
/*
给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1：
输入：p = [1,2,3], q = [1,2,3]
输出：true

示例 2：
输入：p = [1,2], q = [1,null,2]
输出：false

示例 3：
输入：p = [1,2,1], q = [1,1,2]
输出：false


提示：
两棵树上的节点数目都在范围 [0, 100] 内
-10^4 <= Node.val <= 10^4
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func isSameTree(p *pkg.TreeNode, q *pkg.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil || p != nil && q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 层序遍历
func isSameTreeBFS(p *pkg.TreeNode, q *pkg.TreeNode) bool {
	// 两树队列, 保证其中的元素都不为 nil
	queue1 := []*pkg.TreeNode{p}
	queue2 := []*pkg.TreeNode{q}
	if p == nil && q == nil {
		return true
	}
	if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	}

	for len(queue1) > 0 && len(queue2) > 0 {
		node1 := queue1[0]
		node2 := queue2[0]
		queue1 = queue1[1:]
		queue2 = queue2[1:]

		if node1.Val != node2.Val {
			return false
		}

		left1, right1 := node1.Left, node1.Right
		left2, right2 := node2.Left, node2.Right

		// 左子树不相等
		if (left1 == nil && left2 != nil) || (left1 != nil && left2 == nil) {
			return false
		}

		// 右子树不相等
		if (right1 == nil && right2 != nil) || (right1 != nil && right2 == nil) {
			return false
		}

		// 节点不为空则加入队列
		if left1 != nil {
			queue1 = append(queue1, left1)
		}
		if right1 != nil {
			queue1 = append(queue1, right1)
		}
		if left2 != nil {
			queue2 = append(queue2, left2)
		}
		if right2 != nil {
			queue2 = append(queue2, right2)
		}
	}
	if len(queue1) != 0 || len(queue2) != 0 {
		return false
	}
	return true
}

// Test Case1: p = [1,2,3], q = [1,2,3]		Output: true
// Test Case2: p = [1,2], q = [1,null,2]		Output: false
// Test Case3: p = [1,2,1], q = [1,1,2]		Output: false
func main() {
	tree1 := pkg.CreateTree()
	tree2 := pkg.CreateTree()
	fmt.Println(isSameTree(tree1, tree2))
	fmt.Println(isSameTreeBFS(tree1, tree2))
}
