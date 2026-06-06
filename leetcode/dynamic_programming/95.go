/*
95. 不同的二叉搜索树 II
给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
示例 1：
输入：n = 3
输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

示例 2：
输入：n = 1
输出：[[1]]

提示：
1 <= n <= 8
*/
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 返回从left到end的所有 二叉搜索树 的集合
func dfs95(left, right int) []*TreeNode {
	if left > right {
		return []*TreeNode{nil}
	}
	ret := make([]*TreeNode, 0)
	//遍历[left,right]，以i为根节点
	for i := left; i <= right; i++ {
		//获取在区间[left,i-1]中左子树所有种类集合
		leftSubTrees := dfs95(left, i-1)

		//获取在区间[i+1,right]中右子树所有种类集合
		rightSubTrees := dfs95(i+1, right)

		//遍历左右子树所有可能，并分别拼接到根节点的左儿子、右儿子上
		for _, lst := range leftSubTrees {
			for _, rst := range rightSubTrees {
				curNode := &TreeNode{Val: i}
				curNode.Left = lst
				curNode.Right = rst
				ret = append(ret, curNode) //添加一棵树
			}
		}

	}
	return ret
}
func generateTrees(n int) []*TreeNode {
	return dfs95(1, n)
}

type treeQueue struct {
	queue []*TreeNode
}

func (q *treeQueue) len() int { return len(q.queue) }
func (q *treeQueue) pop() {
	q.queue = q.queue[1:]
}
func (q *treeQueue) top() *TreeNode {
	ret := q.queue[0]
	return ret
}
func (q *treeQueue) Push(node *TreeNode) {
	q.queue = append(q.queue, node)
}

func printTreeByLevel(head *TreeNode) {
	var queue treeQueue
	var nextLevelHead = head //下一层头节点
	var a bool               //表示是否已确认下一层的起始节点
	var level = 0
	queue.Push(head)
	for queue.len() > 0 {
		node := queue.top()
		if node == nextLevelHead {
			level++
			a = false //寻找下一层头节点
			fmt.Println()
			fmt.Printf("level%d: ", level)
		}

		if node.Left != nil {
			queue.Push(node.Left)
		}
		if node.Right != nil {
			queue.Push(node.Right)
		}

		if !a {
			if node.Left != nil {
				nextLevelHead = node.Left
				a = true
			} else if node.Right != nil {
				nextLevelHead = node.Right
				a = true
			}
		}
		fmt.Printf("%d ", node.Val)
		queue.pop()
	}
}

func main() {
	var n int
	fmt.Println("Input n:")
	fmt.Scan(&n)
	ret := generateTrees(n)
	for _, root := range ret {
		printTreeByLevel(root)
	}
}
