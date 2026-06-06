/*
199. 二叉树的右视图
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func rightSideView(root *pkg.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := make([]int, 0)
	que := make([]*pkg.TreeNode, 0)
	que = append(que, root)
	for len(que) > 0 {
		sz := len(que)
		ans = append(ans, que[sz-1].Val)
		for i := 0; i < sz; i++ {
			if que[i].Left != nil {
				que = append(que, que[i].Left)
			}
			if que[i].Right != nil {
				que = append(que, que[i].Right)
			}
		}
		que = que[sz:]
	}
	return ans
}
func main() {
	root := pkg.CreateTree()
	fmt.Println(rightSideView(root))
}
