// 103. 二叉树的锯齿形层序遍历
/*
给你二叉树的根节点 root，
返回其节点值的 锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

提示：
树中节点数目在范围 [0, 2000] 内
-100 <= Node.val <= 100
*/
package main

import (
	"fmt"
	"lc/pkg"
)

func zigzagLevelOrder(root *pkg.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	type treeNode struct {
		node  *pkg.TreeNode
		level int
	}
	var ans [][]int
	var tmp []int
	level := 1
	treeNodeQueue := make([]*treeNode, 0)
	treeNodeQueue = append(treeNodeQueue, &treeNode{root, 1})
	for len(treeNodeQueue) > 0 {
		curNode := treeNodeQueue[0]
		if curNode.level != level {
			// 偶数层, 逆序
			if level%2 == 0 {
				n := len(tmp)
				for i := n - 1; i >= n/2; i-- {
					tmp[i], tmp[n-i-1] = tmp[n-i-1], tmp[i]
				}
			}
			ans = append(ans, tmp)
			tmp = make([]int, 0)
			level = curNode.level
		}

		tmp = append(tmp, curNode.node.Val)
		treeNodeQueue = treeNodeQueue[1:]
		if curNode.node.Left != nil {
			treeNodeQueue = append(treeNodeQueue, &treeNode{curNode.node.Left, curNode.level + 1})
		}
		if curNode.node.Right != nil {
			treeNodeQueue = append(treeNodeQueue, &treeNode{curNode.node.Right, curNode.level + 1})
		}
	}
	if len(tmp) > 0 {
		if level%2 == 0 {
			n := len(tmp)
			for i := n - 1; i >= n/2; i-- {
				tmp[i], tmp[n-i-1] = tmp[n-i-1], tmp[i]
			}
		}
		ans = append(ans, tmp)
		tmp = make([]int, 0)
	}
	return ans
}

// 示例 1：
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[20,9],[15,7]]
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
	fmt.Println(zigzagLevelOrder(root))
}
