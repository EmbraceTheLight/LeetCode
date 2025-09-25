/*
103. 二叉树的锯齿形层序遍历
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
*/
package main

import (
	"fmt"
	"lc/100/pkg"
)

func zigzagLevelOrder(root *pkg.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	que := make([]*pkg.TreeNode, 0)
	sli := make([]int, 0)
	ans := make([][]int, 0)
	var flag = true
	que = append(que, root)
	for len(que) > 0 {
		sz := len(que)
		if flag == false {

			for i := sz - 1; i >= 0; i-- {
				sli = append(sli, que[i].Val)
				if que[i].Right != nil {
					que = append(que, que[i].Right)
				}
				if que[i].Left != nil {
					que = append(que, que[i].Left)
				}

			}
			que = que[sz:]
			ans = append(ans, sli)
			sli = sli[sz:]

		} else {
			for i := 0; i < sz; i++ {
				sli = append(sli, que[i].Val)
				if que[i].Left != nil {
					que = append(que, que[i].Left)
				}
				if que[i].Right != nil {
					que = append(que, que[i].Right)
				}
			}
			que = que[sz:]
			ans = append(ans, sli)
			sli = sli[sz:]
			flag = false
		}

	}
	return ans
}
func main() {
	root := pkg.CreateTree()
	fmt.Println(zigzagLevelOrder(root))
}
